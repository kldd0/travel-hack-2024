package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kldd0/travel-hack-2024/internal/entity"
	"github.com/kldd0/travel-hack-2024/internal/pkg/postgres"
	"github.com/kldd0/travel-hack-2024/internal/repository/repoerrs"
)

type TourRepository struct {
	*postgres.Postgres
}

func NewTourRepository(pg *postgres.Postgres) *TourRepository {
	return &TourRepository{pg}
}

func (r *TourRepository) GetById(ctx context.Context, id int) (entity.Tour, error) {
	sql, args, _ := r.Builder.
		Select("*").
		From("tours").
		Where("id = ?", id).
		ToSql()

	var tour entity.Tour
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&tour.Id,
		&tour.Title,
		&tour.Location,
		&tour.Category,
		&tour.Tags, /* from []text */
		&tour.Desc,
		&tour.NightsCount,
		&tour.Program,     /* from []text */
		&tour.Included,    /* from []text */
		&tour.NotIncluded, /* from []text */
		&tour.DifficultyLevel,
		&tour.ComfortLevel,
		&tour.Dates, /* from []text */
		&tour.ImportantInfo,
		&tour.HeadMedia,         /* from []text */
		&tour.ProgramMedia,      /* from []text */
		&tour.AccomodationMedia, /* from []text */
		&tour.MapSrc,
		&tour.Faq,
		&tour.Rating,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Tour{}, repoerrs.ErrNotFound
		}
		return entity.Tour{}, fmt.Errorf("TourRepository.GetTourById - r.Pool.QueryRow: %v", err)
	}

	return tour, nil
}

func (r *TourRepository) GetMany(ctx context.Context, filters map[string]interface{}) ([]entity.Tour, error) {
	builder := r.Builder.
		Select("*").
		From("tours")

	/*
		if len(filters) != 0 {
			for key, value := range filters {
				switch key {
				case "when", "from_name", "adults", "childrens", "price_from", "price_to", "guaranteed", "food_id":
					continue
				case "to_name":
					builder = builder.Where(squirrel.Eq{"location": value})
				case "tags":
					tags, _ := value.([]string)
					builder = builder.Where(squirrel.Expr("tags::text[] @> ARRAY["+squirrel.Placeholders(len(tags))+"]", value))
				case "with_flight", "with_acc", "with_food", "day_off", "low_cost", "age_group":
					continue
				default:
					builder = builder.Where(squirrel.Eq{key: value})
				}
			}
		}
	*/

	sql, args, _ := builder.ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("TourRepository.GetMany - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var tours []entity.Tour
	for rows.Next() {
		var tour entity.Tour
		err := rows.Scan(
			&tour.Id,
			&tour.Title,
			&tour.Location,
			&tour.Category,
			&tour.Tags,
			&tour.Desc,
			&tour.NightsCount,
			&tour.Program,
			&tour.Included,
			&tour.NotIncluded,
			&tour.DifficultyLevel,
			&tour.ComfortLevel,
			&tour.Dates,
			&tour.ImportantInfo,
			&tour.HeadMedia,
			&tour.ProgramMedia,
			&tour.AccomodationMedia,
			&tour.MapSrc,
			&tour.Faq,
			&tour.Rating,
		)
		if err != nil {
			return nil, fmt.Errorf("TourRepository.GetAllProducts - rows.Scan: %v", err)
		}
		tours = append(tours, tour)
	}

	return tours, nil
}
