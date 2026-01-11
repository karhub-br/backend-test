package beerstyle

import (
	"context"
	"database/sql"
)

type Querier interface {
	QueryContext(ctx context.Context, query string, args []any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args []any) error
}
