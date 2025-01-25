package main

import (
	"github.com/bklaing2/autofy/lambdas/util/models"
	"github.com/bklaing2/autofy/lambdas/util/spotify"
)

func updatePlaylist(playlist, playlistUpdates models.Playlist, fetchArtistTracks spotify.FetchArtistTracks) (models.Playlist, []string, []string, error) {
	tracksToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistTracks)
	tracksToRemove := fetchTracksToRemove(playlist, playlistUpdates, fetchArtistTracks)

	return playlistUpdates, tracksToAdd, tracksToRemove, nil
}
