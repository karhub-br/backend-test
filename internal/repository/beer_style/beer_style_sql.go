package beerstyle

import (
	"context"
	"karhub/internal/entity"
)

type repo struct {
	repoSql Querier
}

func (r *repo) Insert(ctx context.Context, beer entity.BeerStyle) (entity.BeerStyle, error) {

	err := r.repoSql.ExecContext(ctx, `INSERT INTO beer_styles (style, min_temperature, max_temperature) VALUES ($1, $2, $3) `, []any{beer.Style, beer.MinTemperature, beer.MaxTemperature})
	if err != nil {
		return entity.BeerStyle{}, err
	}

	return beer, nil
}

func (r *repo) Get(ctx context.Context, temperature int) ([]entity.BeerStyle, error) {

	rows, err := r.repoSql.QueryContext(ctx, `SELECT style, min_temperature, max_temperature FROM beer_styles
		 WHERE  $1 >= min_temperature  AND  $1 <= max_temperature  ORDER BY style ASC`,
		[]any{temperature})
	if err != nil {
		return nil, err
	}

	var beers []entity.BeerStyle
	for rows.Next() {
		var b entity.BeerStyle
		if err := rows.Scan(&b.Style, &b.MinTemperature, &b.MaxTemperature); err != nil {
			return nil, err
		}
		beers = append(beers, b)
	}

	return beers, nil
}

func (r *repo) Update(ctx context.Context, beer entity.BeerStyle) (entity.BeerStyle, error) {

	err := r.repoSql.ExecContext(ctx, `UPDATE beer_styles SET min_temperature=$2, max_temperature=$3 WHERE style=$1`,
		[]any{beer.Style, beer.MinTemperature, beer.MaxTemperature})
	if err != nil {
		return entity.BeerStyle{}, err
	}

	return beer, nil
}

func (r *repo) Delete(ctx context.Context, style string) error {
	err := r.repoSql.ExecContext(ctx, `DELETE FROM beer_styles WHERE style = $1`,
		[]any{style})

	return err
}

func NewRepo(pg Querier) *repo {
	return &repo{repoSql: pg}
}
