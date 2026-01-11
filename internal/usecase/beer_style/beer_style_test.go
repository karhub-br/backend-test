package beerstyle

import (
	"context"
	"errors"
	"karhub/internal/entity"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
)

type mockRepo struct {
	insertFn func(beer entity.BeerStyle) (entity.BeerStyle, error)
	getFn    func(temp int) ([]entity.BeerStyle, error)
}

func (m *mockRepo) Insert(ctx context.Context, b entity.BeerStyle) (entity.BeerStyle, error) {
	return m.insertFn(b)
}
func (m *mockRepo) Get(ctx context.Context, t int) ([]entity.BeerStyle, error) { return m.getFn(t) }
func (m *mockRepo) Update(ctx context.Context, b entity.BeerStyle) (entity.BeerStyle, error) {
	return b, nil
}
func (m *mockRepo) Delete(ctx context.Context, s string) error { return nil }

type mockRabbit struct {
	published bool
	consumed  bool
}

func (m *mockRabbit) Publish(q string, b []byte) error {
	m.published = true
	return nil
}

func (m *mockRabbit) Consume(queue string) (<-chan amqp.Delivery, error) {
	m.consumed = true
	return nil, nil
}

type mockSpotify struct {
	playlistID string
}

func (m *mockSpotify) Playlist(ctx context.Context, s string) (entity.SpotifySearchResponse, error) {
	return entity.SpotifySearchResponse{
		Playlists: entity.Items{
			Items: []*entity.ID{{ID: m.playlistID}},
		},
	}, nil
}

func (m *mockSpotify) Get(ctx context.Context, id string) (entity.SpotifyPlaylistResponse, error) {
	return entity.SpotifyPlaylistResponse{
		Name: "Beer Party",
		Tracks: entity.PlaylistTracks{
			Items: []entity.TrackItem{
				{
					Track: entity.TrackDetails{
						Name:    "test",
						Artists: []entity.Artist{{Name: "test1"}},
						ExternalURLs: entity.ExternalURLs{
							Spotify: "test.com",
						},
					},
				},
			},
		},
	}, nil
}

func TestBeerStyle_AllFunctions(t *testing.T) {
	ctx := context.Background()

	t.Run("Create: Error should publish in Rabbit", func(t *testing.T) {
		rabbit := &mockRabbit{}
		repo := &mockRepo{
			insertFn: func(b entity.BeerStyle) (entity.BeerStyle, error) {
				return entity.BeerStyle{}, errors.New("db error")
			},
		}

		uc := NewBeerStyle(repo, nil, nil, rabbit)
		_, err := uc.Create(ctx, entity.BeerStyle{Style: "Pilsen"})

		if err == nil {
			t.Error("should return error")
		}
		if !rabbit.published {
			t.Error("should have published to rabbit")
		}
	})

	t.Run("Read: Correct playlist and style selection", func(t *testing.T) {
		repo := &mockRepo{
			getFn: func(temp int) ([]entity.BeerStyle, error) {
				return []entity.BeerStyle{
					{Style: "IPA", MinTemperature: -5, MaxTemperature: -1},
				}, nil
			},
		}
		spot := &mockSpotify{playlistID: "456"}

		uc := NewBeerStyle(repo, spot, spot, nil)
		res, err := uc.Read(ctx, entity.BeerTemperature{Temperature: -3})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res.BeerStyle != "IPA" {
			t.Errorf("expected IPA, got %s", res.BeerStyle)
		}
	})

	t.Run("Logic: Alphabetic Tie-Breaking", func(t *testing.T) {
		uc := &beerStyle{}
		beers := []entity.BeerStyle{
			{Style: "Weiss", MinTemperature: 0, MaxTemperature: 4}, // Average 2
			{Style: "Bock", MinTemperature: 0, MaxTemperature: 4},  // Average 2
		}

		res := uc.selectBeerStyle(beers, entity.BeerTemperature{Temperature: 2})
		if res.Style != "Bock" {
			t.Errorf("expected Bock (alphabetic), got %s", res.Style)
		}
	})

	t.Run("Delete: Success", func(t *testing.T) {
		uc := NewBeerStyle(&mockRepo{}, nil, nil, nil)
		err := uc.Delete(ctx, "Stout")
		if err != nil {
			t.Errorf("error in delete: %v", err)
		}
	})
}
