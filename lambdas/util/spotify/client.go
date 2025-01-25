package spotify

import (
	"context"

	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

func CreateClient(ctx context.Context, refreshToken string, accessToken string) *Spotify {
	token := &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	client := spotifyauth.New().Client(ctx, token)

	return &Spotify{
		ctx:    ctx,
		client: spotify.New(client),
	}
}
