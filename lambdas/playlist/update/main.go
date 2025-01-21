package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func processMessage() {
}

func processMessages(event events.SQSEvent) error {
	for _, record := range event.Records {
		_ = record.Body
		// updatePlaylist(record.Body)
	}

	return nil
}

func main() {
	lambda.Start(processMessages)
}
