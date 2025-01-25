package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sst/sst/v3/sdk/golang/resource"
)

type Database struct {
	ctx    context.Context
	client *pgxpool.Pool
}

func CreatePool(ctx context.Context) (*Database, error) {
	dbUrl, err := getDbUrl()
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return nil, err
	}

	return &Database{
		ctx:    ctx,
		client: db,
	}, nil
}

func (db *Database) Close() {
	db.client.Close()
}

func getDbUrl() (string, error) {
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
