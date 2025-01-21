package main

import "testing"

func TestUpdateSettings(t *testing.T) {
	t.Run("artists updated", func(t *testing.T) {
		playlist := Playlist{
			Artists: []string{"artist 1", "artist 2"},
		}
		updates := Playlist{
			Artists: []string{"artist 2", "artist 3", "artist 4"},
		}
		wanted := Playlist{
			Artists: []string{"artist 2", "artist 3", "artist 4"},
		}

		updatedPlaylist, _, _, err := updatePlaylist(playlist, updates, fetchArtistSongs)

		if err != nil {
			t.Fatalf("Error updating playlist: %v", err)
		}

		if !compare(updatedPlaylist.Artists, wanted.Artists) {
			t.Fatalf("Expected %v, got %v", wanted.Artists, updatedPlaylist.Artists)
		}
	})

	t.Run("followed artists updated", func(t *testing.T) {
		playlist := Playlist{
			FollowedArtists: []string{"artist 1", "artist 2"},
		}
		updates := Playlist{
			FollowedArtists: []string{"artist 2", "artist 3", "artist 4"},
		}
		wanted := Playlist{
			FollowedArtists: []string{"artist 2", "artist 3", "artist 4"},
		}

		updatedPlaylist, _, _, err := updatePlaylist(playlist, updates, fetchArtistSongs)

		if err != nil {
			t.Fatalf("Error updating playlist: %v", err)
		}

		if !compare(updatedPlaylist.FollowedArtists, wanted.FollowedArtists) {
			t.Fatalf("Expected %v, got %v", wanted.Artists, updatedPlaylist.Artists)
		}
	})

	t.Run("update when artist posts flag updated", func(t *testing.T) {
		playlist := Playlist{
			UpdateWhenArtistPosts: false,
		}
		updates := Playlist{
			UpdateWhenArtistPosts: true,
		}
		wanted := Playlist{
			UpdateWhenArtistPosts: true,
		}

		updatedPlaylist, _, _, err := updatePlaylist(playlist, updates, fetchArtistSongs)

		if err != nil {
			t.Fatalf("Error updating playlist: %v", err)
		}

		if updatedPlaylist.UpdateWhenArtistPosts != wanted.UpdateWhenArtistPosts {
			t.Fatalf("Expected %v, got %v", wanted.UpdateWhenArtistPosts, updatedPlaylist.UpdateWhenArtistPosts)
		}
	})

	t.Run("update when user follows artist flag updated", func(t *testing.T) {
		playlist := Playlist{
			UpdateWhenUserFollowsArtist: false,
		}
		updates := Playlist{
			UpdateWhenUserFollowsArtist: true,
		}
		wanted := Playlist{
			UpdateWhenUserFollowsArtist: true,
		}

		updatedPlaylist, _, _, err := updatePlaylist(playlist, updates, fetchArtistSongs)

		if err != nil {
			t.Fatalf("Error updating playlist: %v", err)
		}

		if updatedPlaylist.UpdateWhenUserFollowsArtist != wanted.UpdateWhenUserFollowsArtist {
			t.Fatalf("Expected %v, got %v", wanted.UpdateWhenUserFollowsArtist, updatedPlaylist.UpdateWhenUserFollowsArtist)
		}
	})

	t.Run("update when user unfollows artist flag updated", func(t *testing.T) {
		playlist := Playlist{
			UpdateWhenUserUnfollowsArtist: false,
		}
		updates := Playlist{
			UpdateWhenUserUnfollowsArtist: true,
		}
		wanted := Playlist{
			UpdateWhenUserUnfollowsArtist: true,
		}

		updatedPlaylist, _, _, err := updatePlaylist(playlist, updates, fetchArtistSongs)

		if err != nil {
			t.Fatalf("Error updating playlist: %v", err)
		}

		if updatedPlaylist.UpdateWhenUserUnfollowsArtist != wanted.UpdateWhenUserUnfollowsArtist {
			t.Fatalf("Expected %v, got %v", wanted.UpdateWhenUserUnfollowsArtist, updatedPlaylist.UpdateWhenUserUnfollowsArtist)
		}
	})
}

func TestUpdatePlaylist(t *testing.T) {
	playlist := Playlist{
		Artists:                       []string{"artist 1"},
		FollowedArtists:               []string{"artist 2"},
		UpdateWhenArtistPosts:         true,
		UpdateWhenUserFollowsArtist:   true,
		UpdateWhenUserUnfollowsArtist: true,
		UpdatedAt:                     3,
	}

	playlistUpdates := Playlist{
		Artists:                       []string{},
		FollowedArtists:               []string{"artist 2", "artist 3"},
		UpdateWhenArtistPosts:         true,
		UpdateWhenUserFollowsArtist:   true,
		UpdateWhenUserUnfollowsArtist: true,
	}

	wantedSongsToAdd := []string{"song 6", "song 7", "song 8", "song 9"}
	wantedSongsToRemove := []string{"song 1", "song 2", "song 4", "song 5"}
	wantedUpdatedPlaylist := Playlist{
		Artists:                       []string{},
		FollowedArtists:               []string{"artist 2", "artist 3"},
		UpdateWhenArtistPosts:         true,
		UpdateWhenUserFollowsArtist:   true,
		UpdateWhenUserUnfollowsArtist: true,
	}

	updatedPlaylist, songsToAdd, songsToRemove, err := updatePlaylist(playlist, playlistUpdates, fetchArtistSongs)

	if err != nil {
		t.Fatalf("Error updating playlist: %v", err)
	}

	songsToAddIds := []string{}
	for _, s := range songsToAdd {
		songsToAddIds = append(songsToAddIds, s.ID)
	}
	if !compare(songsToAddIds, wantedSongsToAdd) {
		t.Fatalf("Expected %v, got %v", wantedSongsToAdd, songsToAddIds)
	}

	songsToRemoveIds := []string{}
	for _, s := range songsToRemove {
		songsToRemoveIds = append(songsToRemoveIds, s.ID)
	}
	if !compare(songsToRemoveIds, wantedSongsToRemove) {
		t.Fatalf("Expected %v, got %v", wantedSongsToRemove, songsToRemoveIds)
	}

	if !compare(updatedPlaylist.Artists, wantedUpdatedPlaylist.Artists) {
		t.Fatalf("Expected %v, got %v", wantedUpdatedPlaylist, updatedPlaylist)
	}

	if !compare(updatedPlaylist.FollowedArtists, wantedUpdatedPlaylist.FollowedArtists) {
		t.Fatalf("Expected %v, got %v", wantedUpdatedPlaylist, updatedPlaylist)
	}
}
