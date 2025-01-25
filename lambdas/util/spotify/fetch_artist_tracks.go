package spotify

import (
	"log"

	"github.com/zmb3/spotify/v2"
)

const LIMIT = 50
const MARKET = "US"

var OPTIONS = []spotify.RequestOption{
	spotify.Limit(LIMIT),
	spotify.Market(MARKET),
}

func (sp *Spotify) FetchArtistTracks(artistID string) []Track {
	// Setup
	client := sp.client
	ctx := sp.ctx

	artistSpotifyID := spotify.ID(artistID)
	var tracks []Track

	// Fetch albums for the artist
	page, err := client.GetArtistAlbums(ctx, artistSpotifyID, nil, OPTIONS...)
	if err != nil {
		log.Printf("Error fetching albums for artist %s: %v", artistSpotifyID, err)
		return nil
	}

	// Fetch tracks for each album
	for _, album := range page.Albums {
		albumTracks, err := client.GetAlbumTracks(ctx, album.ID, OPTIONS...)
		if err != nil {
			log.Printf("Error fetching tracks for album %s: %v", album.Name, err)
			continue
		}

		for _, track := range albumTracks.Tracks {
			tracks = append(tracks, Track{
				ID:          track.ID.String(),
				ReleaseDate: track.Album.ReleaseDateTime(),
			})
		}
	}

	return tracks
}
