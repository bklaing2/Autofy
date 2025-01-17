package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

const QUERY = `
	SELECT row_to_json(joined)
	FROM (
		SELECT row_to_json(playlists) AS playlist, row_to_json(users) AS user
		FROM playlists
		JOIN users ON playlists.user_id = users.id
	) joined
`

func main() {
	// Setup
	dbUrl, queueUrl := os.Getenv("DATABASE_URL"), os.Getenv("AWS_SQS_URL")

	if dbUrl == "" {
		log.Fatalf("DATABASE_URL environment variable is not set")
	}

	if queueUrl == "" {
		log.Fatalf("AWS_SQS_URL environment variable is not set")
	}

	ctx := context.Background()

	db, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("Failed to load AWS configuration: %v", err)
	}

	// Fetch playlists from database
	playlists, err := db.Query(ctx, QUERY)
	if err != nil {
		db.Close()
		log.Fatalf("Failed to fetch playlist IDs: %v", err)
	}
	defer playlists.Close()

	// Add playlists to SQS
	sqsClient := sqs.NewFromConfig(cfg)

	for playlists.Next() {
		var playlist string
		if err := playlists.Scan(&playlist); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		log.Printf("Adding playlist to queue: %s", playlist)

		_, err = sqsClient.SendMessage(context.TODO(), &sqs.SendMessageInput{
			QueueUrl:    aws.String(queueUrl),
			MessageBody: aws.String(playlist),
		})
		if err != nil {
			log.Printf("Error adding playlist to queue: %v", err)
		}
	}

	// Check for any iteration errors
	if err := playlists.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
	}
}
