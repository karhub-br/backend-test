package beerstyle

import (
	"context"
	"karhub/internal/entity"
	"karhub/pkg/telemetry"
	"time"
)

type beerStyleDecorator struct {
	bs *beerStyle
}

func (b *beerStyleDecorator) Create(ctx context.Context, beerStyle entity.BeerStyle) (entity.BeerStyle, error) {

	defer telemetry.TrackDuration(ctx, "create_beer", time.Now())

	return b.bs.Create(ctx, beerStyle)

}

func (b *beerStyleDecorator) Read(ctx context.Context, temperature entity.BeerTemperature) (entity.BeerPlaylistResponse, error) {

	defer telemetry.TrackDuration(ctx, "read_beer", time.Now())

	return b.bs.Read(ctx, temperature)

}

func (b *beerStyleDecorator) Update(ctx context.Context, beerStyle entity.BeerStyle) (entity.BeerStyle, error) {

	defer telemetry.TrackDuration(ctx, "update_beer", time.Now())

	return b.bs.Update(ctx, beerStyle)

}

func (b *beerStyleDecorator) Delete(ctx context.Context, style string) error {

	defer telemetry.TrackDuration(ctx, "delete_beer", time.Now())

	return b.bs.Delete(ctx, style)

}

func NewBeerStyleDecorator(beerStyle *beerStyle) *beerStyleDecorator {
	return &beerStyleDecorator{bs: beerStyle}
}
