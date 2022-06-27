import os
from urllib.parse import urlparse
from flask import Flask, session, request, redirect, url_for, render_template, jsonify, abort
from flask_session import Session
from pymongo import MongoClient
from bson.objectid import ObjectId
import spotipy
import uuid
import redis

from app.playlist import Playlist # , generate_playlist, update_playlist
from app.redis_cache_handler import RedisCacheHandler


# Set up server
app = Flask(__name__)
app.config['SECRET_KEY'] = os.urandom(64)
app.config['SESSION_TYPE'] = 'filesystem'
app.config['SESSION_FILE_DIR'] = './.flask_session-test/'
Session(app)


# users: {
#   session_id
# }

# playlists: {
#   userId
#   playlistIds: []
#   updatedAt
#   settings: {}
# }

# Mega Playlist Settings {
#   type
#   sortingOrder
#   updateWhen: []?
# }


# Set up Redis
url = urlparse(os.environ.get("REDIS_TLS_URL"))
r = redis.Redis(host=url.hostname, port=url.port, username=url.username, password=url.password, ssl=True, ssl_cert_reqs=None)

# Set up Mongo
mongo = MongoClient(os.environ.get('MONGO_URL'))
db = mongo.autofy
playlists_coll = db.playlists


@app.route('/')
def index():
    # Setup
    # If visitor is unknown, give random ID
    if not session.get('uuid'):
        session['uuid'] = str(uuid.uuid4())

    # Spotify auth
    cache_handler = RedisCacheHandler(r, session.get('uuid'))
    auth_manager = spotipy.oauth2.SpotifyOAuth(scope='user-follow-read playlist-read-private playlist-modify-private', cache_handler=cache_handler, show_dialog=True)

    # If redirected from Spotify auth, add access token
    if request.args.get('code'):
        auth_manager.get_access_token(request.args.get('code'))
        return redirect(url_for('index'))


    # If not logged in, show sign in page
    if not auth_manager.validate_token(cache_handler.get_cached_token()):
        auth_url = auth_manager.get_authorize_url()
        return render_template('sign-in.html', auth_url=auth_url)






    # If logged in, show home page
    spotify = spotipy.Spotify(auth_manager=auth_manager)
    user_id = spotify.current_user()['id']
    # check_if_playlists_deleted(spotify)


    if (profile_picture := spotify.current_user()['images'][0]) is not None:
        return render_template('index.html',
                               name=spotify.current_user()['display_name'],
                               profile_picture=profile_picture,
                               playlist_ids=get_playlist_ids(user_id))

    else:
        return render_template('index.html', name=spotify.current_user()['display_name'], playlist_ids=get_playlist_ids(user_id))



@ app.route('/sign_out')
def sign_out():
    r.delete(session.get('uuid'))
    session.clear()
    return redirect(url_for('index'))






@app.route('/create-playlist', methods=['POST'])
def create_playlist():
    # Create playlist
    form_data = request.form
    for key, value in form_data.items():
        print(key, ':', value)

    print(form_data.getlist('update'))

    cache_handler = RedisCacheHandler(r, session.get('uuid'))
    auth_manager = spotipy.oauth2.SpotifyOAuth(cache_handler=cache_handler)
    if not auth_manager.validate_token(cache_handler.get_cached_token()):
        return redirect(url_for('index'))

    spotify = spotipy.Spotify(auth_manager=auth_manager)


    # Create playlist
    new_playlist = Playlist(spotify)
    new_playlist.generate()

    # Add to database
    obj = new_playlist.get_json()
    playlists_coll.insert_one(obj)

    return { 'playlistIds': obj['playlistIds'] }




def get_playlist_ids(user_id):
    # Get all playlists with user id
    playlist_ids = []
    for playlist in playlists_coll.find({'userId': user_id}):
        playlist_ids.extend(playlist['playlistIds'])

    return playlist_ids


@app.route('/update-playlists', methods=['PUT'])
def update_playlists():
    # Setup
    # If visitor is unknown, give random ID
    if not session.get('uuid'):
        session['uuid'] = str(uuid.uuid4())

    # Spotify auth
    cache_handler = RedisCacheHandler(r, session.get('uuid'))
    auth_manager = spotipy.oauth2.SpotifyOAuth(scope='user-follow-read playlist-read-private playlist-modify-private',
                                               cache_handler=cache_handler, show_dialog=True)

    # If redirected from Spotify auth, add access token
    if request.args.get('code'):
        auth_manager.get_access_token(request.args.get('code'))
        return redirect(url_for('index'))

    # If not logged in, show sign in page
    if not auth_manager.validate_token(cache_handler.get_cached_token()):
        abort(403)

    # If logged in, show home page
    spotify = spotipy.Spotify(auth_manager=auth_manager)
    user_id = spotify.current_user()['id']

    for playlist in playlists_coll.find({'userId': user_id}):
        new_playlist = Playlist(spotify, playlist)
        new_playlist.update()

    return jsonify({'success': True})

@app.route('/update-playlist/<playlist_id>', methods=['GET'])
def update_playlist(playlist_id):
    cache_handler = RedisCacheHandler(r, session.get('uuid'))
    auth_manager = spotipy.oauth2.SpotifyOAuth(cache_handler=cache_handler)
    if not auth_manager.validate_token(cache_handler.get_cached_token()):
        return redirect(url_for('index'))

    spotify = spotipy.Spotify(auth_manager=auth_manager)

    playlist_obj = playlists_coll.find_one({'_id': ObjectId(playlist_id)})
    playlist = Playlist(spotify, playlist_obj)
    playlist.update()

    # Add to database
    playlists_coll.update_one({'_id': ObjectId(playlist_id)}, { '$set': playlist.get_json() })

    return playlist_id



def check_if_playlists_deleted(spotify):
    playlist_ids = get_playlist_ids(spotify.current_user()['id'])
    playlists = []

    offset = 0
    while True:
        results = list(map(lambda res: res['id'], spotify.current_user_playlists(offset=offset)['items']))

        if len(results) == 0: break

        playlists.extend(results)
        offset = offset + 50




    for id in playlist_ids:
        # If playlist not on spotify, remove from database
        if id not in playlists:
            playlists_coll.delete_many({'playlist_id': id})






if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8888)