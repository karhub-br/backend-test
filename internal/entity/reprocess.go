package entity

type Reprocess struct {
	QueryType string    `json:"query_type"`
	BeerStyle BeerStyle `json:"beer_style"`
}
