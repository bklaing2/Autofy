package models

type TracksToUpdatePayload struct {
	PlaylistID     string   `json:"playlist_id"`
	TracksToAdd    []string `json:"tracks_to_add"`
	TracksToRemove []string `json:"tracks_to_remove"`
}

type PlaylistsToUpdatePayload struct {
	Updates Playlist `json:"updates"`
}
