package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bklaing2/autofy/lambdas/util/database"
	"github.com/bklaing2/autofy/lambdas/util/queue"
)

func queuePlaylists(ctx context.Context) error {
	// Setup
	db := database.CreatePool(ctx)
	defer db.Close()

	// Fetch playlists from database
	playlists, err := db.FetchPlaylists()
	if err != nil {
		return err
	}
	defer playlists.Close()

	// Add playlists to queue
	var playlist db.Playlist
	var user db.User

	for playlists.Next(&playlist, &user) {
		log.Printf("Adding playlist to queue: %s", playlist.ID)

		err := queue.UpdatePlaylist(ctx, playlist, user)
		if err != nil {
			log.Printf("Error adding playlist to queue: %v", err)
		}
	}

	// Check for any iteration errors
	if err := playlists.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(queuePlaylists)
}
