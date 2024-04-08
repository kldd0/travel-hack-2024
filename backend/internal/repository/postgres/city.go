package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/pkg/postgres"
)

type CityRepository struct {
	*postgres.Postgres
}

func NewCityRepository(pg *postgres.Postgres) *CityRepository {
	return &CityRepository{pg}
}

func (r *CityRepository) GetMany(ctx context.Context, prefix string, limit int) ([]entity.City, error) {
	sql, args, _ := r.Builder.
		Select("*").
		From("cities").
		Where(squirrel.Like{"name": prefix + "%"}).
		ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("CityRepository.GetMany - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var cities []entity.City
	for rows.Next() {
		var city entity.City
		err := rows.Scan(
			&city.Id,
			&city.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("CityRepository.GetAllProducts - rows.Scan: %v", err)
		}
		cities = append(cities, city)
	}

	return cities, nil
}
