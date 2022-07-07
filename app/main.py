import os
import uuid
from urllib.parse import urlparse
from flask import Flask, session, request, redirect, url_for, render_template, jsonify
from flask_session import Session
import redis
from rq import Queue
from pymongo import MongoClient
from bson.objectid import ObjectId
import spotipy

from app.src.playlist import Playlist
from app.src.cache_handlers import MemoryCacheHandler, RedisCacheHandler
from app.src.worker import conn


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
url = urlparse(os.environ.get('REDIS_TLS_URL'))
r = redis.Redis(host=url.hostname, port=url.port, username=url.username, password=url.password, ssl=True, ssl_cert_reqs=None)

q = Queue(connection=conn)

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
    user = spotify.current_user()
    profile_picture = user['images'][0] if 'images' in user and len(user['images']) > 0 else None

    return render_template('index.html', name=user['display_name'], profile_picture=profile_picture)



@app.route('/sign_out')
def sign_out():
    r.delete(session.get('uuid'))
    session.clear()
    return redirect(url_for('index'))






@app.route('/create-playlist', methods=['POST'])
def create_playlist():
    cache_handler = RedisCacheHandler(r, session.get('uuid'))
    auth_manager = spotipy.oauth2.SpotifyOAuth(cache_handler=cache_handler)
    if not auth_manager.validate_token(cache_handler.get_cached_token()):
        return redirect(url_for('index'))

    spotify = spotipy.Spotify(auth_manager=auth_manager)

    # Add default document
    new_id = playlists_coll.insert_one(Playlist(spotify).get_json(new=True)).inserted_id

    q.enqueue(create_playlist_helper, new_id, job_timeout=7200) # 2 hrs
    return { 'playlist': { 'id': str(new_id), 'spotifyPlaylists': 'generating' } }

def create_playlist_helper(playlist_id):
    playlist_obj = playlists_coll.find_one({'_id': ObjectId(playlist_id)})

    cache_handler = MemoryCacheHandler(token_info=playlist_obj['token'])
    auth_manager = spotipy.oauth2.SpotifyOAuth(cache_handler=cache_handler)
    if not auth_manager.validate_token(cache_handler.get_cached_token()):
        return jsonify({'success': False})

    spotify = spotipy.Spotify(auth_manager=auth_manager)

    # Create playlist
    new_playlist = Playlist(spotify)
    new_playlist.generate()

    # Update in database
    playlists_coll.update_one({'_id': ObjectId(playlist_obj['_id'])}, {'$set': new_playlist.get_json()})



@app.route('/get-playlists', methods=['GET'])
def get_playlists():
    cache_handler = RedisCacheHandler(r, session.get('uuid'))
    auth_manager = spotipy.oauth2.SpotifyOAuth(cache_handler=cache_handler)
    if not auth_manager.validate_token(cache_handler.get_cached_token()):
        return redirect(url_for('index'))

    spotify = spotipy.Spotify(auth_manager=auth_manager)

    # Get all playlists with user id
    playlists = []
    for playlist in playlists_coll.find({'userId': spotify.current_user()['id']}):
        playlists.append({
            'id': str(playlist['_id']),
            'spotifyPlaylists': playlist['playlistIds']
        })


    return jsonify({ 'playlists': playlists })


@app.route('/update-playlists', methods=['GET'])
def update_playlists():
    print('Updating all playlists...')
    for playlist_obj in playlists_coll.find({ 'token': { '$exists': True }}):
        print(playlist_obj['_id'])

        cache_handler = MemoryCacheHandler(token_info=playlist_obj['token'])
        auth_manager = spotipy.oauth2.SpotifyOAuth(cache_handler=cache_handler)
        if not auth_manager.validate_token(cache_handler.get_cached_token()):
            print('Error with token')
            continue

        spotify = spotipy.Spotify(auth_manager=auth_manager)


        playlist = Playlist(spotify, playlist_obj)
        if playlist.has_deleted():
            playlist.delete_all()
            playlists_coll.delete_one({'_id': ObjectId(playlist_obj['_id'])})
        else:
            playlist.update()
            playlists_coll.update_one({'_id': ObjectId(playlist_obj['_id'])}, {'$set': playlist.get_json()})

    return jsonify({'success': True})

@app.route('/update-playlist/<playlist_id>', methods=['GET'])
def update_playlist(playlist_id):
    playlist_obj = playlists_coll.find_one({'_id': ObjectId(playlist_id)})

    cache_handler = MemoryCacheHandler(token_info=playlist_obj['token'])
    auth_manager = spotipy.oauth2.SpotifyOAuth(cache_handler=cache_handler)
    if not auth_manager.validate_token(cache_handler.get_cached_token()):
        return jsonify({'success': False})

    spotify = spotipy.Spotify(auth_manager=auth_manager)

    playlist = Playlist(spotify, playlist_obj)

    if playlist.has_deleted():
        playlist.delete_all()
        playlists_coll.delete_one({'_id': ObjectId(playlist_id)})
    else:
        playlist.update()
        playlists_coll.update_one({'_id': ObjectId(playlist_id)}, { '$set': playlist.get_json() })

    return jsonify({'success': True})






if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8888)