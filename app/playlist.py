from datetime import datetime, date


# Maximum playlist size allowed by Spotify
MAX_LENGTH = 10000


class Playlist:

    def __init__(self, spotify, playlist={}):
        self.spotify = spotify

        self.user_id = spotify.current_user()['id']
        self.playlist_ids = playlist['playlistIds'] if 'playlistIds' in playlist else []
        self.artist_ids = playlist['artists'] if 'artists' in playlist else None

        self.updated_at = playlist['updatedAt'] if 'updatedAt' in playlist else None


    def generate(self):
        print('Generating playlist...\n')


        # Get all tracks from followed artists and add to new playlist
        print('Getting followed artists...')
        self.artist_ids = self.get_followed_artist_ids()
        print(f'{len(self.artist_ids)} artists\n')

        print('Getting track ids...')
        track_ids = self.get_track_ids_by_artist_ids(self.artist_ids)
        print(f'{len(track_ids)} tracks\n')

        print('Adding tracks...\n')
        self.add_tracks(track_ids)


        self.updated_at = datetime.now()
        print('Playlists generated!')
        print(self.playlist_ids)


    def update(self):
        followed_artist_ids = self.get_followed_artist_ids()

        # Get unfollowed and followed artists since last update
        unfollowed_since = list(set(self.artist_ids) - set(followed_artist_ids))
        followed_since = list(set(followed_artist_ids) - set(self.artist_ids))


        # Remove tracks by unfollowed artists
        tracks = self.get_track_ids_by_artist_ids(unfollowed_since)
        self.remove_tracks(tracks)

        # Add tracks by followed artists
        tracks = self.get_track_ids_by_artist_ids(followed_since)
        self.add_tracks(tracks)

        # Add tracks released since playlist was last updated
        tracks = self.get_track_ids_uploaded_since_last_updated(self.artist_ids)
        self.add_tracks(tracks)


        self.artist_ids = followed_artist_ids
        self.updated_at = datetime.now()




    def add_tracks(self, track_ids):
        i = 0

        while len(track_ids) > 0:

            # Get next available playlist's capacity
            if i < len(self.playlist_ids):
                playlist_capacity = MAX_LENGTH - self.spotify.playlist(playlist_id=self.playlist_ids[i])['tracks']['total']

            # Or create a new playlist
            else:
                playlist = self.spotify.user_playlist_create(
                    self.user_id,
                    f"Everything - {datetime.now().strftime('%m/%d/%Y')}{f' ({i+1})' if i > 0 else ''}",
                    public=False,
                    collaborative=False,
                    description='Created on ' + datetime.now().strftime('%m/%d/%Y, %H:%M:%S') + ' by autofy')

                self.playlist_ids.append(playlist['id'])
                playlist_capacity = MAX_LENGTH


            # Get track ids to add to playlist, and remove added tracks from list
            tracks_to_add = track_ids[:playlist_capacity]
            track_ids = track_ids[playlist_capacity:]

            # Add tracks to playlist
            while tracks_to_add:
                self.spotify.playlist_add_items(playlist_id=self.playlist_ids[i], items=tracks_to_add[:50])
                tracks_to_add = tracks_to_add[50:]

            i += 1


    def remove_tracks(self, track_ids):
        # Iterate through each playlist and remove tracks
        for playlist_id in self.playlist_ids:
            i = 0
            while i < len(track_ids):
                self.spotify.playlist_remove_all_occurrences_of_items(playlist_id=playlist_id, items=track_ids[i:i+100])

                if self.spotify.playlist(playlist_id=playlist_id)['tracks']['total'] == 0:
                    self.playlist_ids.remove(playlist_id)
                    self.spotify.current_user_unfollow_playlist(playlist_id)
                    break

                i += 100




    # Return values as a dictionary
    def get_json(self):
        return {
            'userId': self.user_id,
            'playlistIds': self.playlist_ids,
            'artists': self.artist_ids,
            'updatedAt': self.updated_at
        }






    # Helper functions


    def get_followed_artist_ids(self):
        artists = []
        prev_artist = None

        while True:
            results = self.spotify.current_user_followed_artists(after=prev_artist)['artists']['items']
            if len(results) == 0: break

            artist_ids = get_ids(results)
            artists.extend(get_ids(results))
            prev_artist = artist_ids[-1]

        return artists



    def get_albums_by_artist_id(self, artist_id):
        albums = []
        offset = 0

        while True:
            results = self.spotify.artist_albums(
                artist_id,
                album_type='album',
                country='US',
                offset=offset
            )['items']

            if len(results) == 0: break

            albums.extend(results)
            offset = offset + 20

        return albums


    def get_track_ids_by_album_id(self, album_id):
        tracks = []
        offset = 0
        while True:
            results = self.spotify.album_tracks(album_id, offset=offset)['items']

            if len(results) == 0: break

            tracks.extend(get_ids(results))
            offset = offset + 50

        return tracks


    def get_track_ids_by_artist_ids(self, artist_ids):
        album_ids = []
        for artist_id in artist_ids:
            albums = self.get_albums_by_artist_id(artist_id)
            album_ids.extend(get_ids(albums))

        track_ids = []
        for album_id in album_ids:
            track_ids.extend(self.get_track_ids_by_album_id(album_id))

        return track_ids


    def get_track_ids_uploaded_since_last_updated(self, artist_ids):
        if not self.updated_at: return []

        album_ids = []
        for artist_id in artist_ids:
            albums = self.get_albums_by_artist_id(artist_id)
            albums = list(filter(self.released_since_last_updated, albums))
            album_ids.extend(get_ids(albums))

        track_ids = []
        for album_id in album_ids:
            track_ids.extend(self.get_track_ids_by_album_id(album_id))

        return track_ids




    def released_since_last_updated(self, item):
        try: release_date = datetime.strptime(item['release_date'], '%Y-%m-%d')
        except ValueError: release_date = datetime.strptime(item['release_date'], '%Y')
        return release_date > self.updated_at









def get_ids(arr):
    return list(map(lambda i: i['id'], arr))






# if __name__ == '__main__':
#     # Set up Spotify
#     scope = "user-follow-read playlist-modify-private"
#     sp = spotipy.Spotify(auth_manager=SpotifyOAuth(scope=scope))
#     generate_playlist(sp)