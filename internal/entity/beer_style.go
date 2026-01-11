package entity

type BeerStyle  struct {
	Style string `json:"beerStyle"`
	MinTemperature int `json:"minTemperature"`
	MaxTemperature int `json:"maxTemperature"`
}

type BeerTemperature struct {
	Temperature int
}

type BeerPlaylistResponse struct {	
	BeerStyle string   `json:"beerStyle"`
	Playlist  Playlist `json:"playlist"`
}

type Playlist struct {
	Name   string  `json:"name"`
	Tracks []Track `json:"tracks"`
}

type Track struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Link   string `json:"link"`
}