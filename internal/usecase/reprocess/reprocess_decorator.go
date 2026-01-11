package reprocess

import (
	"context"
	"karhub/internal/entity"
	"karhub/pkg/telemetry"
	"time"
)

type reprocessDecorator struct {
	rp *reprocess
}

func (r *reprocessDecorator) Reprocess(ctx context.Context, rep entity.Reprocess) error {

	defer telemetry.TrackDuration(ctx, "reprocess", time.Now())

	return r.rp.Reprocess(ctx, rep)
}

func NewReprocessDecorator(rep *reprocess) *reprocessDecorator {
	return &reprocessDecorator{rp: rep}
}
