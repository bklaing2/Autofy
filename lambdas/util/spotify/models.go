package spotify

import (
	"context"
	"time"

	zmb3spotify "github.com/zmb3/spotify/v2"
)

type Spotify struct {
	ctx    context.Context
	client *zmb3spotify.Client
}

type Track struct {
	ID          string
	ReleaseDate time.Time
}

type FetchArtistTracks = func(string) []Track
