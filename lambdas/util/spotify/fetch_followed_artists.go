package spotify

import (
	"log"

	"github.com/zmb3/spotify/v2"
)

func (sp *Spotify) FetchFollowedArtists() []string {
	// Setup
	client := sp.client
	ctx := sp.ctx

	var followedArtistsIDs []string

	// Fetch followed artists
	after := ""
	for {
		page, err := client.CurrentUsersFollowedArtists(ctx, spotify.After(after))
		if err != nil {
			log.Printf("Error fetching followed artists: %v", err)
			return followedArtistsIDs
		}

		for _, artist := range page.Artists {
			followedArtistsIDs = append(followedArtistsIDs, artist.ID.String())
		}

		after = page.Cursor.After
		if after == "" {
			break
		}
	}

	return followedArtistsIDs
}
