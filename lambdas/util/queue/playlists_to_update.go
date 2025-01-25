package queue

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/bklaing2/autofy/lambdas/util/models"
)

func PlaylistsToUpdate(ctx context.Context, playlist models.Playlist) error {
	// Setup
	queueUrl, err := getQueueUrl("PlaylistsToUpdate")
	if err != nil {
		return err
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		return err
	}

	sqsClient := sqs.NewFromConfig(cfg)

	// Create the JSON payload
	payload := models.PlaylistsToUpdatePayload{
		Updates: playlist,
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Add to queue
	_, err = sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueUrl),
		MessageBody: aws.String(string(payloadJson)),
	})

	return err
}
