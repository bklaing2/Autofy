package models

import "time"

type Playlist struct {
	ID                            string    `json:"id"`
	UserID                        string    `json:"user_id"`
	Artists                       []string  `json:"artists"`
	FollowedArtists               []string  `json:"followed_artists"`
	IncludeFollowedArtists        bool      `json:"include_followed_artists"`
	UpdateWhenArtistPosts         bool      `json:"update_when_artist_posts"`
	UpdateWhenUserFollowsArtist   bool      `json:"update_when_user_follows_artist"`
	UpdateWhenUserUnfollowsArtist bool      `json:"update_when_user_unfollows_artist"`
	UpdatedAt                     time.Time `json:"updated_at"`
}

type User struct {
	ID           string `json:"id"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
