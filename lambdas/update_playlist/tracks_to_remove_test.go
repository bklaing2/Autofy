package main

import (
	"testing"

	"github.com/bklaing2/autofy/lambdas/util/models"
)

func TestSongsToRemoveWhenArtistsRemoved(t *testing.T) {
	t.Run("artist is in playlist", func(t *testing.T) {
		playlist := models.Playlist{
			Artists: []string{"artist 1"},
		}

		playlistUpdates := models.Playlist{
			Artists: []string{},
		}

		wantedSongsToRemove := []string{"song 1", "song 2", "song 4", "song 5"}

		songsToRemove := fetchTracksToRemove(playlist, playlistUpdates, fetchArtistSongs)

		songsToRemoveIds := []string{}
		for _, s := range songsToRemove {
			songsToRemoveIds = append(songsToRemoveIds, s.ID)
		}

		if !compare(songsToRemoveIds, wantedSongsToRemove) {
			t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
		}
	})

	t.Run("artist is not in playlist", func(t *testing.T) {
		playlist := models.Playlist{
			Artists: []string{},
		}

		playlistUpdates := models.Playlist{
			Artists: []string{},
		}

		wantedSongsToRemove := []string{}

		songsToRemove := fetchTracksToRemove(playlist, playlistUpdates, fetchArtistSongs)

		songsToRemoveIds := []string{}
		for _, s := range songsToRemove {
			songsToRemoveIds = append(songsToRemoveIds, s.ID)
		}

		if !compare(songsToRemoveIds, wantedSongsToRemove) {
			t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
		}
	})

	t.Run("artist is in playlist followed", func(t *testing.T) {
		playlist := models.Playlist{
			Artists: []string{"artist 1"},
		}

		playlistUpdates := models.Playlist{
			FollowedArtists: []string{"artist 1"},
		}

		wantedSongsToRemove := []string{}

		songsToRemove := fetchTracksToRemove(playlist, playlistUpdates, fetchArtistSongs)

		songsToRemoveIds := []string{}
		for _, s := range songsToRemove {
			songsToRemoveIds = append(songsToRemoveIds, s.ID)
		}

		if !compare(songsToRemoveIds, wantedSongsToRemove) {
			t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
		}
	})
}

func TestSongsToRemoveWhenUserUnfollowsArtist(t *testing.T) {
	t.Run("flag unset", func(t *testing.T) {
		playlist := models.Playlist{
			FollowedArtists:               []string{"artist 1"},
			UpdateWhenUserUnfollowsArtist: false,
		}

		playlistUpdates := models.Playlist{
			FollowedArtists:               []string{},
			UpdateWhenUserUnfollowsArtist: false,
		}

		wantedSongsToRemove := []string{}

		songsToRemove := fetchTracksToRemove(playlist, playlistUpdates, fetchArtistSongs)

		songsToRemoveIds := []string{}
		for _, s := range songsToRemove {
			songsToRemoveIds = append(songsToRemoveIds, s.ID)
		}

		if !compare(songsToRemoveIds, wantedSongsToRemove) {
			t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
		}
	})

	t.Run("flag set", func(t *testing.T) {
		playlist := models.Playlist{
			FollowedArtists:               []string{"artist 1"},
			UpdateWhenUserUnfollowsArtist: true,
		}

		playlistUpdates := models.Playlist{
			FollowedArtists:               []string{},
			UpdateWhenUserUnfollowsArtist: true,
		}

		wantedSongsToRemove := []string{"song 1", "song 2", "song 4", "song 5"}

		songsToRemove := fetchTracksToRemove(playlist, playlistUpdates, fetchArtistSongs)

		songsToRemoveIds := []string{}
		for _, s := range songsToRemove {
			songsToRemoveIds = append(songsToRemoveIds, s.ID)
		}

		if !compare(songsToRemoveIds, wantedSongsToRemove) {
			t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
		}
	})

	t.Run("artist not followed", func(t *testing.T) {
		playlist := models.Playlist{
			FollowedArtists:               []string{},
			UpdateWhenUserUnfollowsArtist: true,
		}

		playlistUpdates := models.Playlist{
			FollowedArtists:               []string{},
			UpdateWhenUserUnfollowsArtist: true,
		}

		wantedSongsToRemove := []string{}

		songsToRemove := fetchTracksToRemove(playlist, playlistUpdates, fetchArtistSongs)

		songsToRemoveIds := []string{}
		for _, s := range songsToRemove {
			songsToRemoveIds = append(songsToRemoveIds, s.ID)
		}

		if !compare(songsToRemoveIds, wantedSongsToRemove) {
			t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
		}
	})
}

func TestSongsToRemove(t *testing.T) {
	playlist := models.Playlist{
		Artists:                       []string{"artist 1", "artist 2"},
		FollowedArtists:               []string{"artist 3"},
		UpdateWhenArtistPosts:         true,
		UpdateWhenUserUnfollowsArtist: true,
		UpdatedAt:                     3,
	}

	playlistUpdates := models.Playlist{
		Artists:                       []string{"artist 1"},
		FollowedArtists:               []string{},
		UpdateWhenArtistPosts:         true,
		UpdateWhenUserUnfollowsArtist: true,
	}

	wantedSongsToRemove := []string{"song 3", "song 6", "song 7", "song 8", "song 9"}

	songsToAdd := fetchTracksToRemove(playlist, playlistUpdates, fetchArtistSongs)

	songsToAddIds := []string{}
	for _, s := range songsToAdd {
		songsToAddIds = append(songsToAddIds, s.ID)
	}

	if !compare(songsToAddIds, wantedSongsToRemove) {
		t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToAddIds)
	}
}
