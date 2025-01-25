package database

import (
	"errors"

	"github.com/bklaing2/autofy/lambdas/util/models"
	"github.com/jackc/pgx/v5"
)

func (database *Database) FetchUser(userID string) (models.User, error) {
	// Setup
	db := database.client
	ctx := database.ctx

	user := models.User{
		ID: userID,
	}

	// Query for the user tokens
	row := db.QueryRow(ctx, selectUserTokensByID, userID)
	err := row.Scan(
		&user.AccessToken,
		&user.RefreshToken,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}
