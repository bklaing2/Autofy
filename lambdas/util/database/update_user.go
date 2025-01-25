package database

import (
	"github.com/bklaing2/autofy/lambdas/util/models"
)

func (database *Database) UpdateUser(user models.User) error {
	// Setup
	db := database.client
	ctx := database.ctx

	// Update the user's tokens
	_, err := db.Exec(ctx, updateUserTokens,
		user.ID,
		user.AccessToken,
		user.RefreshToken,
	)
	if err != nil {
		return err
	}

	return nil
}
