package main

import (
	"fmt"

	"github.com/sst/sst/v3/sdk/golang/resource"
)

func DbUrl() (string, error) {
	username, err := resource.Get("Db", "username")
	if err != nil {
		return "", fmt.Errorf("Failed to get database username: %v", err)
	}

	password, err := resource.Get("Db", "password")
	if err != nil {
		return "", fmt.Errorf("Failed to get database password: %v", err)
	}

	host, err := resource.Get("Db", "host")
	if err != nil {
		return "", fmt.Errorf("Failed to get database host: %v", err)
	}

	port, err := resource.Get("Db", "port")
	if err != nil {
		return "", fmt.Errorf("Failed to get database port: %v", err)
	}

	name, err := resource.Get("Db", "name")
	if err != nil {
		return "", fmt.Errorf("Failed to get database name: %v", err)
	}

	username, password, host, port, name = username.(string), password.(string), host.(string), port.(string), name.(string)
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, name), nil
}

func QueueUrl() (string, error) {
	queueUrl, err := resource.Get("UpdatePlaylist", "url")
	if err != nil {
		return "", fmt.Errorf("Failed to get UpdatePlaylist queue url: %v", err)
	}

	return queueUrl.(string), nil
}
