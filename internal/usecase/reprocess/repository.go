package reprocess

import (
	"context"
	"karhub/internal/entity"
)

type Repository interface {
	Insert(ctx context.Context, beer entity.BeerStyle) (entity.BeerStyle, error)
	Update(ctx context.Context, beer entity.BeerStyle) (entity.BeerStyle, error)
	Delete(ctx context.Context, name string) error
}
