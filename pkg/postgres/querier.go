package postgres

import (
	"context"
	"database/sql"
)

type querier struct {
	db *sql.DB
}

func (d *querier) QueryContext(ctx context.Context, query string, args []any) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, args...)

}

func (d *querier) ExecContext(ctx context.Context, query string, args []any) error {

	_, err := d.db.ExecContext(ctx, query, args...)

	return err
}

func NewQuerier(db *sql.DB) *querier {
	return &querier{db: db}
}
