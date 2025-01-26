package spotify

import (
	"log"

	zmb3spotify "github.com/zmb3/spotify/v2"
)

func (sp *Spotify) RemoveTracksFromPlaylist(playlistID string, trackIDs []string) {
	// Setup
	client := sp.client
	ctx := sp.ctx

	playlistSpotifyID := zmb3spotify.ID(playlistID)
	trackSpotifyIDs := []zmb3spotify.ID{}

	for _, trackId := range trackIDs {
		trackSpotifyIDs = append(trackSpotifyIDs, zmb3spotify.ID(trackId))
	}

	// Remove tracks in batches of 100
	for i := 0; i < len(trackSpotifyIDs); i += batchSize {
		end := i + batchSize
		if end > len(trackSpotifyIDs) {
			end = len(trackSpotifyIDs)
		}

		_, err := client.RemoveTracksFromPlaylist(ctx, playlistSpotifyID, trackSpotifyIDs[i:end]...)
		if err != nil {
			log.Printf("Error removing tracks from playlist: %v", err)
		}
	}

}
