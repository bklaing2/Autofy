package database

import (
	"github.com/bklaing2/autofy/lambdas/util/models"
	"github.com/jackc/pgx/v5"
)

type PlaylistUserIterator struct {
	rows pgx.Rows
}

// Next moves the iterator to the next row and populates the given Playlist and User structs.
func (iter *PlaylistUserIterator) Next(playlist *models.Playlist) bool {
	if !iter.rows.Next() {
		return false
	}

	// Scan the current row into Playlist and User structs
	err := iter.rows.Scan(
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
		return false
	}

	return true
}

// Err returns any error encountered during iteration.
func (iter *PlaylistUserIterator) Err() error {
	return iter.rows.Err()
}

// Close cleans up the underlying rows.
func (iter *PlaylistUserIterator) Close() {
	iter.rows.Close()
}

func (database *Database) FetchPlaylists() (*PlaylistUserIterator, error) {
	// Setup
	db := database.client
	ctx := database.ctx

	// Query for all playlists
	rows, err := db.Query(ctx, selectPlaylists)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &PlaylistUserIterator{rows: rows}, nil
}
