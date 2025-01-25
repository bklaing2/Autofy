package main

import (
	"github.com/bklaing2/autofy/lambdas/util/models"
	"github.com/bklaing2/autofy/lambdas/util/set"
	"github.com/bklaing2/autofy/lambdas/util/spotify"
)

func fetchTracksToAdd(playlist, updates models.Playlist, fetchArtistTracks spotify.FetchArtistTracks) []string {
	tracksToAdd := []string{}

	playlistArtists := set.ToSet(playlist.Artists)
	if playlist.IncludeFollowedArtists {
		playlistArtists.Append(updates.FollowedArtists)
	}

	// Fetch songs released since the playlist was last updated
	if playlist.UpdateWhenArtistPosts {
		playlistUpdatesArtists := set.ToSet(append(updates.Artists, updates.FollowedArtists...))
		artistsToUpdate := playlistArtists.Intersection(playlistUpdatesArtists)

		for _, artistID := range artistsToUpdate.List() {
			artistTracks := fetchArtistTracks(artistID)

			for _, track := range artistTracks {
				if track.ReleaseDate.Compare(playlist.UpdatedAt) <= 0 {
					tracksToAdd = append(tracksToAdd, track.ID)
				}
			}
		}
	}

	// Fetch songs by artists added to the playlist and followed since the playlist was last updated
	artistsToAdd := set.ToSet(updates.Artists)
	if playlist.IncludeFollowedArtists && playlist.UpdateWhenUserFollowsArtist {
		artistsToAdd.Append(updates.FollowedArtists)
	}

	artistsToAdd = artistsToAdd.Difference(playlistArtists)

	for _, artistID := range artistsToAdd.List() {
		artistTracks := fetchArtistTracks(artistID)

		for _, track := range artistTracks {
			tracksToAdd = append(tracksToAdd, track.ID)
		}
	}

	return tracksToAdd
}

func fetchTracksToRemove(playlist, updates models.Playlist, fetchArtistTracks spotify.FetchArtistTracks) []string {
	tracksToRemove := []string{}

	playlistUpdatesArtists := set.ToSet(updates.Artists)
	if playlist.IncludeFollowedArtists {
		playlistUpdatesArtists.Append(updates.FollowedArtists)
	}

	// Fetch songs by artists removed from the playlist and unfollowed since the playlist was last updated
	artistsToRemove := set.ToSet(playlist.Artists).Difference(playlistUpdatesArtists)

	if playlist.IncludeFollowedArtists && playlist.UpdateWhenUserUnfollowsArtist {
		artistsUnfollowed := set.ToSet(playlist.FollowedArtists).Difference(playlistUpdatesArtists)
		artistsToRemove = artistsToRemove.Union(artistsUnfollowed)
	}

	for _, artistID := range artistsToRemove.List() {
		artistTracks := fetchArtistTracks(artistID)

		for _, track := range artistTracks {
			tracksToRemove = append(tracksToRemove, track.ID)
		}
	}

	return tracksToRemove
}
