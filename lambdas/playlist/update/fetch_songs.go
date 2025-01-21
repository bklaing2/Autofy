package main

type FetchArtistSongs func(Artist) []Song

func fetchSongsToAdd(playlist, playlistUpdates Playlist, fetchArtistSongs FetchArtistSongs) []Song {
	songsToAdd := []Song{}

	playlistArtists := ToSet(append(playlist.Artists, playlist.FollowedArtists...))

	// Fetch songs released since the playlist was last updated
	if playlist.UpdateWhenArtistPosts {
		playlistUpdatesArtists := ToSet(append(playlistUpdates.Artists, playlistUpdates.FollowedArtists...))
		artistsToUpdate := playlistArtists.Intersection(playlistUpdatesArtists)

		for _, artist := range artistsToUpdate.List() {
			artistSongs := fetchArtistSongs(artist)

			for _, song := range artistSongs {
				if song.Released > playlist.UpdatedAt {
					songsToAdd = append(songsToAdd, song)
				}
			}
		}
	}

	// Fetch songs by artists added to the playlist and followed since the playlist was last updated
	artistsToAdd := ToSet(playlistUpdates.Artists)
	if playlist.UpdateWhenUserFollowsArtist {
		artistsToAdd.Append(playlistUpdates.FollowedArtists)
	}

	artistsToAdd = artistsToAdd.Difference(playlistArtists)

	for _, artist := range artistsToAdd.List() {
		artistSongs := fetchArtistSongs(artist)
		songsToAdd = append(songsToAdd, artistSongs...)
	}

	return songsToAdd
}

func fetchSongsToRemove(playlist, playlistUpdates Playlist, fetchArtistSongs FetchArtistSongs) []Song {
	songsToRemove := []Song{}

	playlistUpdatesArtists := ToSet(append(playlistUpdates.Artists, playlistUpdates.FollowedArtists...))

	// Fetch songs by artists removed from the playlist and unfollowed since the playlist was last updated
	artistsToRemove := ToSet(playlist.Artists).Difference(playlistUpdatesArtists)

	if playlist.UpdateWhenUserUnfollowsArtist {
		artistsUnfollowed := ToSet(playlist.FollowedArtists).Difference(playlistUpdatesArtists)
		artistsToRemove = artistsToRemove.Union(artistsUnfollowed)
	}

	for _, artist := range artistsToRemove.List() {
		artistSongs := fetchArtistSongs(artist)
		songsToRemove = append(songsToRemove, artistSongs...)
	}

	return songsToRemove
}
