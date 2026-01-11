package entity

type SpotifySearchResponse struct {
	Playlists Items `json:"playlists"`
}

type Items struct {
	Items []*ID `json:"items"`
}

type ID struct {
	ID string `json:"id"`
}