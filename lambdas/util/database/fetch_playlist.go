package database

import (
	"errors"

	"github.com/bklaing2/autofy/lambdas/util/models"
	"github.com/jackc/pgx/v5"
)

func (database *Database) FetchPlaylist(playlistID string) (models.Playlist, error) {
	// Setup
	db := database.client
	ctx := database.ctx

	playlist := models.Playlist{}

	// Query for the specific playlist
	row := db.QueryRow(ctx, selectPlaylistByID, playlistID)
	err := row.Scan(
		&playlist.ID,
		&playlist.UserID,
		&playlist.Artists,
		&playlist.FollowedArtists,
		&playlist.IncludeFollowedArtists,
		&playlist.UpdateWhenArtistPosts,
		&playlist.UpdateWhenUserFollowsArtist,
		&playlist.UpdateWhenUserUnfollowsArtist,
		&playlist.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return playlist, errors.New("playlist not found")
		}
		return playlist, err
	}

	return playlist, nil
}
