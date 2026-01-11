package entity

type SpotifyPlaylistResponse struct {
	Name   string         `json:"name"`
	Tracks PlaylistTracks `json:"tracks"`
}

type PlaylistTracks struct {
	Items []TrackItem `json:"items"`
}

type TrackItem struct {
	Track TrackDetails `json:"track"`
}

type TrackDetails struct {
	Name         string       `json:"name"`
	Artists      []Artist     `json:"artists"`
	ExternalURLs ExternalURLs `json:"external_urls"`
}

type Artist struct {
	Name string `json:"name"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}