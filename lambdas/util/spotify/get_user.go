package spotify

import (
	"log"

	"github.com/bklaing2/autofy/lambdas/util/models"
)

func (sp *Spotify) User() models.User {
	client := sp.client
	ctx := sp.ctx

	user, err := client.CurrentUser(ctx)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return models.User{}
	}

	tokens, err := client.Token()
	if err != nil {
		log.Printf("Error fetching tokens: %v", err)
		return models.User{}
	}

	return models.User{
		ID:           user.ID,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}
