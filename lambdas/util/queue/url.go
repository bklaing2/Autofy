package main

import (
	"fmt"

	"github.com/sst/sst/v3/sdk/golang/resource"
)

func getQueueUrl(queueName string) (string, error) {
	queueUrl, err := resource.Get(queueName, "url")
	if err != nil {
		return "", fmt.Errorf("Failed to get PlaylistsToUpdate queue url: %v", err)
	}

	return queueUrl.(string), nil
}
