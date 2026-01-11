package consumer

import (
	"context"
	"karhub/internal/entity"
)

type Reprocess interface {
	Reprocess(ctx context.Context, rep entity.Reprocess) error
}
