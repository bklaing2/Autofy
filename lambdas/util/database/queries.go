package database

// Playlist queries
const selectPlaylists = `
	SELECT 
		p.id, 
		p.user_id,
		p.artists, 
		p.followed_artists, 
		p.include_followed_artists,
		p.update_when_artist_posts, 
		p.update_when_user_follows_artist, 
		p.update_when_user_unfollows_artist, 
		p.updated_at, 
	FROM playlists p
`

const selectPlaylistByID = selectPlaylists + `
	WHERE p.id = $1
`

const updatePlaylist = `
	UPDATE playlists
	SET
		artists = $2,
		followed_artists = $3,
		include_followed_artists = $4,
		update_when_artist_posts = $5,
		update_when_user_follows_artist = $6,
		update_when_user_unfollows_artist = $7,
		updated_at = $8
	WHERE id = $1
`

// User queries
const selectUserTokensByID = `
	SELECT access_token, refresh_token
	FROM users
	WHERE id = $1
`

const updateUserTokens = `
	UPDATE users
	SET access_token = $2, refresh_token = $3
	WHERE id = $1
`
