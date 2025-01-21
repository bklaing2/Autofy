package main

import (
	"context"
	"log"

	"github.com/bklaing2/autofy/lambdas"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

const QUERY = "SELECT id FROM playlists"

func queuePlaylists(ctx context.Context) error {
	// Setup
	dbUrl, err := lambdas.DbUrl()
	if err != nil {
		return err
	}

	queueUrl, err := lambdas.QueueUrl()
	if err != nil {
		return err
	}

	db, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		return err
	}

	// Fetch playlists from database
	playlists, err := db.Query(ctx, QUERY)
	if err != nil {
		db.Close()
		return err
	}
	defer playlists.Close()

	// Add playlists to queue
	sqsClient := sqs.NewFromConfig(cfg)

	for playlists.Next() {
		var playlistId string
		if err := playlists.Scan(&playlistId); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		log.Printf("Adding playlist to queue: %s", playlistId)

		_, err = sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
			QueueUrl:    aws.String(queueUrl),
			MessageBody: aws.String(playlistId),
		})
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
