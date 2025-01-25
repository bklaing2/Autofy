package database

import (
	"github.com/bklaing2/autofy/lambdas/util/models"
)

func (database *Database) UpdatePlaylist(playlist models.Playlist) error {
	// Setup
	db := database.client
	ctx := database.ctx

	// Update the playlist
	_, err := db.Exec(ctx, updatePlaylist,
		playlist.ID,
		playlist.Artists,
		playlist.FollowedArtists,
		playlist.IncludeFollowedArtists,
		playlist.UpdateWhenArtistPosts,
		playlist.UpdateWhenUserFollowsArtist,
		playlist.UpdateWhenUserUnfollowsArtist,
		playlist.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
