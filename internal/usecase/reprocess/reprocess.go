package reprocess

import (
	"context"
	"karhub/internal/entity"
)

type reprocess struct {
	reprocessRepository Repository
}

func (r *reprocess) Reprocess(ctx context.Context, rep entity.Reprocess) error {

	switch rep.QueryType {
	case "update":
		_, err := r.reprocessRepository.Update(ctx, rep.BeerStyle)
		return err
	case "insert":
		_, err := r.reprocessRepository.Insert(ctx, rep.BeerStyle)
		return err
	case "delete":
		return r.reprocessRepository.Delete(ctx, rep.BeerStyle.Style)
	}

	return nil
}

func NewReprocess(repo Repository) *reprocess {
	return &reprocess{reprocessRepository: repo}
}
