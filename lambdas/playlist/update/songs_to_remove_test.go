package main

import "testing"

func TestSongsToRemoveWhenArtistsRemoved(t *testing.T) {
	t.Run("artist is in playlist", func(t *testing.T) {
		playlist := Playlist{
			Artists: []string{"artist 1"},
		}

		playlistUpdates := Playlist{
			Artists: []string{},
		}

		wantedSongsToRemove := []string{"song 1", "song 2", "song 4", "song 5"}

		songsToRemove := fetchSongsToRemove(playlist, playlistUpdates, fetchArtistSongs)

		songsToRemoveIds := []string{}
		for _, s := range songsToRemove {
			songsToRemoveIds = append(songsToRemoveIds, s.ID)
		}

		if !compare(songsToRemoveIds, wantedSongsToRemove) {
			t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
		}
	})

	t.Run("artist is not in playlist", func(t *testing.T) {
		playlist := Playlist{
			Artists: []string{},
		}

		playlistUpdates := Playlist{
			Artists: []string{},
		}

		wantedSongsToRemove := []string{}

		songsToRemove := fetchSongsToRemove(playlist, playlistUpdates, fetchArtistSongs)

		songsToRemoveIds := []string{}
		for _, s := range songsToRemove {
			songsToRemoveIds = append(songsToRemoveIds, s.ID)
		}

		if !compare(songsToRemoveIds, wantedSongsToRemove) {
			t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
		}
	})

	t.Run("artist is in playlist followed", func(t *testing.T) {
		playlist := Playlist{
			Artists: []string{"artist 1"},
		}

		playlistUpdates := Playlist{
			FollowedArtists: []string{"artist 1"},
		}

		wantedSongsToRemove := []string{}

		songsToRemove := fetchSongsToRemove(playlist, playlistUpdates, fetchArtistSongs)

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
		playlist := Playlist{
			FollowedArtists:               []string{"artist 1"},
			UpdateWhenUserUnfollowsArtist: false,
		}

		playlistUpdates := Playlist{
			FollowedArtists:               []string{},
			UpdateWhenUserUnfollowsArtist: false,
		}

		wantedSongsToRemove := []string{}

		songsToRemove := fetchSongsToRemove(playlist, playlistUpdates, fetchArtistSongs)

		songsToRemoveIds := []string{}
		for _, s := range songsToRemove {
			songsToRemoveIds = append(songsToRemoveIds, s.ID)
		}

		if !compare(songsToRemoveIds, wantedSongsToRemove) {
			t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
		}
	})

	t.Run("flag set", func(t *testing.T) {
		playlist := Playlist{
			FollowedArtists:               []string{"artist 1"},
			UpdateWhenUserUnfollowsArtist: true,
		}

		playlistUpdates := Playlist{
			FollowedArtists:               []string{},
			UpdateWhenUserUnfollowsArtist: true,
		}

		wantedSongsToRemove := []string{"song 1", "song 2", "song 4", "song 5"}

		songsToRemove := fetchSongsToRemove(playlist, playlistUpdates, fetchArtistSongs)

		songsToRemoveIds := []string{}
		for _, s := range songsToRemove {
			songsToRemoveIds = append(songsToRemoveIds, s.ID)
		}

		if !compare(songsToRemoveIds, wantedSongsToRemove) {
			t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
		}
	})

	t.Run("artist not followed", func(t *testing.T) {
		playlist := Playlist{
			FollowedArtists:               []string{},
			UpdateWhenUserUnfollowsArtist: true,
		}

		playlistUpdates := Playlist{
			FollowedArtists:               []string{},
			UpdateWhenUserUnfollowsArtist: true,
		}

		wantedSongsToRemove := []string{}

		songsToRemove := fetchSongsToRemove(playlist, playlistUpdates, fetchArtistSongs)

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
	playlist := Playlist{
		Artists:                       []string{"artist 1", "artist 2"},
		FollowedArtists:               []string{"artist 3"},
		UpdateWhenArtistPosts:         true,
		UpdateWhenUserUnfollowsArtist: true,
		UpdatedAt:                     3,
	}

	playlistUpdates := Playlist{
		Artists:                       []string{"artist 1"},
		FollowedArtists:               []string{},
		UpdateWhenArtistPosts:         true,
		UpdateWhenUserUnfollowsArtist: true,
	}

	wantedSongsToRemove := []string{"song 3", "song 6", "song 7", "song 8", "song 9"}

	songsToAdd := fetchSongsToRemove(playlist, playlistUpdates, fetchArtistSongs)

	songsToAddIds := []string{}
	for _, s := range songsToAdd {
		songsToAddIds = append(songsToAddIds, s.ID)
	}

	if !compare(songsToAddIds, wantedSongsToRemove) {
		t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToAddIds)
	}
}
