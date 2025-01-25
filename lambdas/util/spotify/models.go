package spotify

import (
	"context"
	"time"

	"github.com/zmb3/spotify/v2"
)

type Spotify struct {
	ctx    context.Context
	client *spotify.Client
}

type Track struct {
	ID          string
	ReleaseDate time.Time
}

type FetchArtistTracks = func(string) []Track
