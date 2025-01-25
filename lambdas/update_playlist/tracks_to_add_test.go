package main

import (
	"testing"

	"github.com/bklaing2/autofy/lambdas/util/models"
	"github.com/bklaing2/autofy/lambdas/util/spotify"
)

func TestSongsToAddUpdateWhenArtistPosts(t *testing.T) {
	t.Run("flag unset", func(t *testing.T) {
		playlist := models.Playlist{
			Artists:               []string{"artist 1", "artist 2"},
			UpdateWhenArtistPosts: false,
			UpdatedAt:             3,
		}

		playlistUpdates := models.Playlist{
			Artists:               []string{"artist 1", "artist 2"},
			UpdateWhenArtistPosts: false,
		}

		wantedSongsToAdd := []string{}

		songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("flag set", func(t *testing.T) {
		playlist := models.Playlist{
			Artists:               []string{"artist 1", "artist 2"},
			UpdateWhenArtistPosts: true,
			UpdatedAt:             3,
		}

		playlistUpdates := models.Playlist{
			Artists:               []string{"artist 1", "artist 2"},
			UpdateWhenArtistPosts: true,
		}

		wantedSongsToAdd := []string{"song 4", "song 5", "song 6", "song 7"}

		songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})
}

func TestSongsToAddWhenArtistsAdded(t *testing.T) {
	playlist := models.Playlist{}
	playlistUpdates := models.Playlist{
		Artists: []string{"artist 1", "artist 2"},
	}

	wantedSongsToAdd := []string{"song 1", "song 2", "song 4", "song 5", "song 3", "song 6", "song 7"}

	songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

	songsToAddIds := []string{}
	for _, s := range songsToAdd {
		songsToAddIds = append(songsToAddIds, s.ID)
	}

	if !compare(songsToAddIds, wantedSongsToAdd) {
		t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
	}
}

func TestSongsToAddWhenUserFollowsArtist(t *testing.T) {
	t.Run("flag unset", func(t *testing.T) {
		playlist := models.Playlist{
			UpdateWhenUserFollowsArtist: false,
		}

		playlistUpdates := models.Playlist{
			FollowedArtists:             []string{"artist 1", "artist 2"},
			UpdateWhenUserFollowsArtist: false,
		}

		wantedSongsToAdd := []string{}

		songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("flag set", func(t *testing.T) {
		playlist := models.Playlist{
			UpdateWhenUserFollowsArtist: true,
		}

		playlistUpdates := models.Playlist{
			FollowedArtists:             []string{"artist 1", "artist 2"},
			UpdateWhenUserFollowsArtist: true,
		}

		wantedSongsToAdd := []string{"song 1", "song 2", "song 4", "song 5", "song 3", "song 6", "song 7"}

		songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})
}

func TestSongsToAddWithCollisions(t *testing.T) {
	t.Run("artist followed is already in playlist - update when artist posts flag unset", func(t *testing.T) {
		playlist := models.Playlist{
			Artists:                     []string{"artist 1"},
			UpdateWhenArtistPosts:       false,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		playlistUpdates := models.Playlist{
			Artists:                     []string{"artist 1"},
			FollowedArtists:             []string{"artist 1"},
			UpdateWhenArtistPosts:       false,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		wantedSongsToAdd := []string{}

		songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("artist followed is already in playlist - update when artist posts flag set", func(t *testing.T) {
		playlist := models.Playlist{
			Artists:                     []string{"artist 1"},
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		playlistUpdates := models.Playlist{
			Artists:                     []string{"artist 1"},
			FollowedArtists:             []string{"artist 1"},
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		wantedSongsToAdd := []string{"song 4", "song 5"}

		songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("artist added is already in playlist", func(t *testing.T) {
		playlist := models.Playlist{
			Artists:                     []string{"artist 1"},
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		playlistUpdates := models.Playlist{
			Artists:                     []string{"artist 1"},
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		wantedSongsToAdd := []string{"song 4", "song 5"}

		songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("artist added is already followed - update when artist posts flag unset", func(t *testing.T) {
		playlist := models.Playlist{
			FollowedArtists:       []string{"artist 1"},
			UpdateWhenArtistPosts: false,
			UpdatedAt:             3,
		}

		playlistUpdates := models.Playlist{
			Artists:               []string{"artist 1"},
			UpdateWhenArtistPosts: false,
			UpdatedAt:             3,
		}

		wantedSongsToAdd := []string{}

		songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("artist added is already followed - update when artist posts flag set", func(t *testing.T) {
		playlist := models.Playlist{
			FollowedArtists:       []string{"artist 1"},
			UpdateWhenArtistPosts: true,
			UpdatedAt:             3,
		}

		playlistUpdates := models.Playlist{
			Artists:               []string{"artist 1"},
			UpdateWhenArtistPosts: true,
			UpdatedAt:             3,
		}

		wantedSongsToAdd := []string{"song 4", "song 5"}

		songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("artist added is same as artist followed since last update", func(t *testing.T) {
		playlist := models.Playlist{
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		playlistUpdates := models.Playlist{
			Artists:                     []string{"artist 1"},
			FollowedArtists:             []string{"artist 1"},
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		wantedSongsToAdd := []string{"song 1", "song 2", "song 4", "song 5"}

		songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})
}

func TestSongsToAdd(t *testing.T) {
	playlist := models.Playlist{
		FollowedArtists:             []string{"artist 1"},
		UpdateWhenArtistPosts:       true,
		UpdateWhenUserFollowsArtist: true,
		UpdatedAt:                   3,
	}

	playlistUpdates := models.Playlist{
		Artists:                     []string{"artist 1", "artist 2"},
		FollowedArtists:             []string{"artist 3"},
		UpdateWhenArtistPosts:       true,
		UpdateWhenUserFollowsArtist: true,
		UpdatedAt:                   3,
	}

	wantedSongsToAdd := []string{"song 4", "song 5", "song 3", "song 6", "song 7", "song 8", "song 9"}

	songsToAdd := fetchTracksToAdd(playlist, playlistUpdates, fetchArtistSongs)

	songsToAddIds := []string{}
	for _, s := range songsToAdd {
		songsToAddIds = append(songsToAddIds, s.ID)
	}

	if !compare(songsToAddIds, wantedSongsToAdd) {
		t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
	}
}

func artistSongs() map[string][]spotify.Track {
	artistSongs := make(map[string][]spotify.Track)
	artistSongs["artist 1"] = []spotify.Track{
		{ID: "song 1", ReleaseDate: 0},
		{ID: "song 2", ReleaseDate: 1},
		{ID: "song 4", ReleaseDate: 4},
		{ID: "song 5", ReleaseDate: 5},
	}
	artistSongs["artist 2"] = []spotify.Track{
		{ID: "song 3", ReleaseDate: 2},
		{ID: "song 6", ReleaseDate: 4},
		{ID: "song 7", ReleaseDate: 5},
	}
	artistSongs["artist 3"] = []spotify.Track{
		{ID: "song 8", ReleaseDate: 2},
		{ID: "song 9", ReleaseDate: 4},
	}

	return artistSongs
}

func fetchArtistSongs(artist string) []spotify.Track {
	return artistSongs()[artist]
}

func compare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[string]int)
	for _, s := range a {
		counts[s]++
	}

	for _, s := range b {
		if counts[s] == 0 {
			return false
		}
		counts[s]--
	}

	return true
}
