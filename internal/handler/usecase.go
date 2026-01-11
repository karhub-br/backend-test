package handlers

import (
	"context"
	"karhub/internal/entity"
)

type BeerUsecase interface {
	Create(ctx context.Context, beerStyle entity.BeerStyle) (entity.BeerStyle, error)
	Read(ctx context.Context, temperature entity.BeerTemperature) (entity.BeerPlaylistResponse, error)
	Update(ctx context.Context, beerStyle entity.BeerStyle) (entity.BeerStyle, error)
	Delete(ctx context.Context, style string) error
}
