package main

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bklaing2/autofy/lambdas/util/database"
	"github.com/bklaing2/autofy/lambdas/util/models"
	"github.com/bklaing2/autofy/lambdas/util/spotify"
)

func processMessage(ctx context.Context, playlistID string, tracksToAdd []string, tracksToRemove []string) error {
	// Setup
	db := database.CreatePool(ctx)
	defer db.Close()

	user, err := db.FetchUser(playlistID)
	if err != nil {
		return err
	}

	client := spotify.CreateClient(ctx, user.AccessToken, user.RefreshToken)

	// And and remove tracks in parallel
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		client.AddTracksToPlaylist(playlistID, tracksToAdd)
	}()

	go func() {
		defer wg.Done()
		client.RemoveTracksFromPlaylist(playlistID, tracksToRemove)
	}()

	wg.Wait()

	err = db.UpdateUser(client.User())
	if err != nil {
		return err
	}

	return nil
}

func processMessages(ctx context.Context, event events.SQSEvent) error {
	for _, record := range event.Records {
		var payload models.TracksToUpdatePayload
		err := json.Unmarshal([]byte(record.Body), &payload)
		if err != nil {
			log.Println(err)
			continue
		}

		processMessage(ctx, payload.PlaylistID, payload.TracksToAdd, payload.TracksToRemove)
	}

	return nil
}

func main() {
	lambda.Start(processMessages)
}
