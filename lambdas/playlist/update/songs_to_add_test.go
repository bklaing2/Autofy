package main

import "testing"

func TestSongsToAddUpdateWhenArtistPosts(t *testing.T) {
	t.Run("flag unset", func(t *testing.T) {
		playlist := Playlist{
			Artists:               []string{"artist 1", "artist 2"},
			UpdateWhenArtistPosts: false,
			UpdatedAt:             3,
		}

		playlistUpdates := Playlist{
			Artists:               []string{"artist 1", "artist 2"},
			UpdateWhenArtistPosts: false,
		}

		wantedSongsToAdd := []string{}

		songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("flag set", func(t *testing.T) {
		playlist := Playlist{
			Artists:               []string{"artist 1", "artist 2"},
			UpdateWhenArtistPosts: true,
			UpdatedAt:             3,
		}

		playlistUpdates := Playlist{
			Artists:               []string{"artist 1", "artist 2"},
			UpdateWhenArtistPosts: true,
		}

		wantedSongsToAdd := []string{"song 4", "song 5", "song 6", "song 7"}

		songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

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
	playlist := Playlist{}
	playlistUpdates := Playlist{
		Artists: []string{"artist 1", "artist 2"},
	}

	wantedSongsToAdd := []string{"song 1", "song 2", "song 4", "song 5", "song 3", "song 6", "song 7"}

	songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

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
		playlist := Playlist{
			UpdateWhenUserFollowsArtist: false,
		}

		playlistUpdates := Playlist{
			FollowedArtists:             []string{"artist 1", "artist 2"},
			UpdateWhenUserFollowsArtist: false,
		}

		wantedSongsToAdd := []string{}

		songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("flag set", func(t *testing.T) {
		playlist := Playlist{
			UpdateWhenUserFollowsArtist: true,
		}

		playlistUpdates := Playlist{
			FollowedArtists:             []string{"artist 1", "artist 2"},
			UpdateWhenUserFollowsArtist: true,
		}

		wantedSongsToAdd := []string{"song 1", "song 2", "song 4", "song 5", "song 3", "song 6", "song 7"}

		songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

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
		playlist := Playlist{
			Artists:                     []string{"artist 1"},
			UpdateWhenArtistPosts:       false,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		playlistUpdates := Playlist{
			Artists:                     []string{"artist 1"},
			FollowedArtists:             []string{"artist 1"},
			UpdateWhenArtistPosts:       false,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		wantedSongsToAdd := []string{}

		songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("artist followed is already in playlist - update when artist posts flag set", func(t *testing.T) {
		playlist := Playlist{
			Artists:                     []string{"artist 1"},
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		playlistUpdates := Playlist{
			Artists:                     []string{"artist 1"},
			FollowedArtists:             []string{"artist 1"},
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		wantedSongsToAdd := []string{"song 4", "song 5"}

		songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("artist added is already in playlist", func(t *testing.T) {
		playlist := Playlist{
			Artists:                     []string{"artist 1"},
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		playlistUpdates := Playlist{
			Artists:                     []string{"artist 1"},
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		wantedSongsToAdd := []string{"song 4", "song 5"}

		songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("artist added is already followed - update when artist posts flag unset", func(t *testing.T) {
		playlist := Playlist{
			FollowedArtists:       []string{"artist 1"},
			UpdateWhenArtistPosts: false,
			UpdatedAt:             3,
		}

		playlistUpdates := Playlist{
			Artists:               []string{"artist 1"},
			UpdateWhenArtistPosts: false,
			UpdatedAt:             3,
		}

		wantedSongsToAdd := []string{}

		songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("artist added is already followed - update when artist posts flag set", func(t *testing.T) {
		playlist := Playlist{
			FollowedArtists:       []string{"artist 1"},
			UpdateWhenArtistPosts: true,
			UpdatedAt:             3,
		}

		playlistUpdates := Playlist{
			Artists:               []string{"artist 1"},
			UpdateWhenArtistPosts: true,
			UpdatedAt:             3,
		}

		wantedSongsToAdd := []string{"song 4", "song 5"}

		songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

		songsToAddIds := []string{}
		for _, s := range songsToAdd {
			songsToAddIds = append(songsToAddIds, s.ID)
		}

		if !compare(songsToAddIds, wantedSongsToAdd) {
			t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
		}
	})

	t.Run("artist added is same as artist followed since last update", func(t *testing.T) {
		playlist := Playlist{
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		playlistUpdates := Playlist{
			Artists:                     []string{"artist 1"},
			FollowedArtists:             []string{"artist 1"},
			UpdateWhenArtistPosts:       true,
			UpdateWhenUserFollowsArtist: true,
			UpdatedAt:                   3,
		}

		wantedSongsToAdd := []string{"song 1", "song 2", "song 4", "song 5"}

		songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

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
	playlist := Playlist{
		FollowedArtists:             []string{"artist 1"},
		UpdateWhenArtistPosts:       true,
		UpdateWhenUserFollowsArtist: true,
		UpdatedAt:                   3,
	}

	playlistUpdates := Playlist{
		Artists:                     []string{"artist 1", "artist 2"},
		FollowedArtists:             []string{"artist 3"},
		UpdateWhenArtistPosts:       true,
		UpdateWhenUserFollowsArtist: true,
		UpdatedAt:                   3,
	}

	wantedSongsToAdd := []string{"song 4", "song 5", "song 3", "song 6", "song 7", "song 8", "song 9"}

	songsToAdd := fetchSongsToAdd(playlist, playlistUpdates, fetchArtistSongs)

	songsToAddIds := []string{}
	for _, s := range songsToAdd {
		songsToAddIds = append(songsToAddIds, s.ID)
	}

	if !compare(songsToAddIds, wantedSongsToAdd) {
		t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
	}
}

func artistSongs() map[string][]Song {
	artistSongs := make(map[string][]Song)
	artistSongs["artist 1"] = []Song{
		{ID: "song 1", Released: 0},
		{ID: "song 2", Released: 1},
		{ID: "song 4", Released: 4},
		{ID: "song 5", Released: 5},
	}
	artistSongs["artist 2"] = []Song{
		{ID: "song 3", Released: 2},
		{ID: "song 6", Released: 4},
		{ID: "song 7", Released: 5},
	}
	artistSongs["artist 3"] = []Song{
		{ID: "song 8", Released: 2},
		{ID: "song 9", Released: 4},
	}

	return artistSongs
}

func fetchArtistSongs(artist string) []Song {
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
