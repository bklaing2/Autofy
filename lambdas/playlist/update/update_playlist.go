package main

type Playlist struct {
	ID                            string   `json:"id"`
	UserID                        string   `json:"user_id"`
	Artists                       []string `json:"artists"`
	FollowedArtists               []string `json:"followed_artists"`
	UpdateWhenArtistPosts         bool     `json:"update_when_artist_posts"`
	UpdateWhenUserFollowsArtist   bool     `json:"update_when_user_follows_artist"`
	UpdateWhenUserUnfollowsArtist bool     `json:"update_when_user_unfollows_artist"`
	UpdatedAt                     int      `json:"updated_at"`
}

type User struct {
	ID           string `json:"id"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type Artist = string
type Song struct {
	ID       string
	Released int
}

func updatePlaylist(playlist, playlistUpdates Playlist, fetchArtistSongs func(Artist) []Song) (Playlist, []Song, []Song, error) {
	songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)
	songsToRemove := fetchSongsToRemove(playlist, playlistUpdates, fetchArtistSongs)

	return playlistUpdates, songsToAdd, songsToRemove, nil
}
