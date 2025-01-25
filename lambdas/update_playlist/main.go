package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bklaing2/autofy/lambdas/util/database"
	"github.com/bklaing2/autofy/lambdas/util/models"
	"github.com/bklaing2/autofy/lambdas/util/queue"
	"github.com/bklaing2/autofy/lambdas/util/spotify"
)

func processMessage(ctx context.Context, updates models.Playlist) error {
	// Setup
	db := database.CreatePool(ctx)
	defer db.Close()

	var playlist models.Playlist
	var user models.User

	// Fetch playlist and user from database
	playlist, err := db.FetchPlaylist(updates.ID)
	if err != nil {
		return err
	}

	user, err = db.FetchUser(playlist.UserID)
	if err != nil {
		return err
	}

	// Create spotify client and fetch user's followed artists
	client := spotify.CreateClient(ctx, user.AccessToken, user.RefreshToken)
	if updates.IncludeFollowedArtists {
		updates.FollowedArtists = client.FetchFollowedArtists()
	}

	// Get updates, tracks to add, and tracks to remove
	updates, tracksToAdd, tracksToRemove, err := updatePlaylist(playlist, updates, client.FetchArtistTracks)
	log.Printf("Tracks to add: %v", tracksToAdd)
	log.Printf("Tracks to remove: %v", tracksToRemove)

	// Write playlist updates and user tokens to DB and queue tracks to update
	err = db.UpdatePlaylist(playlist)
	if err != nil {
		return err
	}

	err = db.UpdateUser(client.User())
	if err != nil {
		return err
	}

	err = queue.TracksToUpdate(ctx, playlist.ID, tracksToAdd, tracksToRemove)
	if err != nil {
		return err
	}

	return err
}

func processMessages(ctx context.Context, event events.SQSEvent) error {
	for _, record := range event.Records {
		var payload models.PlaylistsToUpdatePayload
		err := json.Unmarshal([]byte(record.Body), &payload)
		if err != nil {
			log.Println(err)
			continue
		}

		processMessage(ctx, payload.Updates)
	}

	return nil
}

func main() {
	lambda.Start(processMessages)
}
