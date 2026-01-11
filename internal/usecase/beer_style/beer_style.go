package beerstyle

import (
	"context"
	"encoding/json"
	"fmt"
	"karhub/internal/entity"

	"math"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

type beerStyle struct {
	beerRepository  Repository
	spotifySearch   Search
	spotifyPlaylist Playlist
	rabbit          Rabbit
}

func (b *beerStyle) Create(ctx context.Context, beerStyle entity.BeerStyle) (entity.BeerStyle, error) {

	beer, err := b.beerRepository.Insert(ctx, beerStyle)
	if err != nil {
		b.publish("insert", beerStyle)
		log.Error(fmt.Sprintf("error to create beerStyle: %s", err))
		return entity.BeerStyle{}, err
	}

	return beer, nil
}

func (b *beerStyle) Read(ctx context.Context, temperature entity.BeerTemperature) (entity.BeerPlaylistResponse, error) {

	beers, err := b.beerRepository.Get(ctx, temperature.Temperature)
	if err != nil {
		log.Error(fmt.Sprintf("error to get beerStyle by temperature: %s", err))
		return entity.BeerPlaylistResponse{}, err
	}

	beerStyle := b.selectBeerStyle(beers, temperature)

	playlistID, err := b.spotifySearch.Playlist(ctx, beerStyle.Style)
	if err != nil {
		log.Error(fmt.Sprintf("error to search in spotify: %s", err))
		return entity.BeerPlaylistResponse{}, err
	}

	if len(playlistID.Playlists.Items) == 0 {
		log.Error(fmt.Sprintf("error playlist id is empty: %s", err))
		return entity.BeerPlaylistResponse{}, fmt.Errorf("no playlist found for beer style: %s", beerStyle.Style)
	}

	playlist, err := b.spotifyPlaylist.Get(ctx, playlistID.Playlists.Items[0].ID)
	if err != nil {
		log.Error(fmt.Sprintf("error to get playlist in spotify: %s", err))
		return entity.BeerPlaylistResponse{}, err
	}

	if playlist.Tracks.Items == nil {
		log.Error(fmt.Sprintf("error playlist is empty: %s", err))
		return entity.BeerPlaylistResponse{}, fmt.Errorf("no playlist found for beer style: %s", beerStyle.Style)
	}

	return b.beerStyleToBeerPlaylistResponse(beerStyle, playlist), nil
}

func (b *beerStyle) selectBeerStyle(beers []entity.BeerStyle, temperature entity.BeerTemperature) entity.BeerStyle {

	var (
		currentBeer       entity.BeerStyle
		currentDifference float64
	)

	if len(beers) == 1 {
		return beers[0]
	}

	for _, beer := range beers {
		avg := (float64(beer.MinTemperature) + float64(beer.MaxTemperature)) / 2

		difference := math.Abs(float64(temperature.Temperature) - float64(avg))

		if currentBeer == (entity.BeerStyle{}) {
			currentBeer = beer
			currentDifference = difference
			continue
		}

		if difference < currentDifference {
			currentBeer = beer
			currentDifference = difference
		}

		if difference == currentDifference {
			if strings.ToLower(beer.Style) < strings.ToLower(currentBeer.Style) {
				currentBeer = beer
				currentDifference = difference
			}
		}

	}

	return currentBeer

}

func (b *beerStyle) beerStyleToBeerPlaylistResponse(beerStyle entity.BeerStyle, spotifyPlaylistResponse entity.SpotifyPlaylistResponse) entity.BeerPlaylistResponse {
	beerPlaylist := entity.BeerPlaylistResponse{
		BeerStyle: beerStyle.Style,
		Playlist: entity.Playlist{
			Name: spotifyPlaylistResponse.Name,
		},
	}

	for _, track := range spotifyPlaylistResponse.Tracks.Items {
		currentTrack := entity.Track{
			Name:   track.Track.Name,
			Artist: track.Track.Artists[0].Name,
			Link:   track.Track.ExternalURLs.Spotify,
		}
		beerPlaylist.Playlist.Tracks = append(beerPlaylist.Playlist.Tracks, currentTrack)
	}

	return beerPlaylist
}

func (b *beerStyle) Update(ctx context.Context, beerStyle entity.BeerStyle) (entity.BeerStyle, error) {
	resp, err := b.beerRepository.Update(ctx, beerStyle)
	if err != nil {
		b.publish("update", beerStyle)
		log.Error(fmt.Sprintf("error to update beerStyle: %s", err))
		return entity.BeerStyle{}, err
	}

	return resp, nil
}

func (b *beerStyle) Delete(ctx context.Context, style string) error {

	beerStyle := entity.BeerStyle{Style: style}

	if err := b.beerRepository.Delete(ctx, style); err != nil {
		b.publish("delete", beerStyle)
		log.Error(fmt.Sprintf("error to delete style: %s", err))
		return err
	}

	return nil
}

func (b *beerStyle) publish(querier string, beerStyle entity.BeerStyle) {
	var reprocessEntity entity.Reprocess

	reprocessEntity.QueryType = querier
	reprocessEntity.BeerStyle = beerStyle

	body, err := json.Marshal(reprocessEntity)
	if err != nil {
		log.Error(fmt.Sprintf("error to publish to reprocess queue: %s", err))
		return
	}

	if err := b.rabbit.Publish("reprocess-queue", body); err != nil {
		log.Errorf(fmt.Sprintf("Error to publish in rabbit queue: %s", err))
		return
	}
}

func NewBeerStyle(beerRepo Repository, spotifySearch Search, spotifyPlaylist Playlist, rabbit Rabbit) *beerStyle {
	return &beerStyle{beerRepository: beerRepo, spotifySearch: spotifySearch, spotifyPlaylist: spotifyPlaylist, rabbit: rabbit}
}
