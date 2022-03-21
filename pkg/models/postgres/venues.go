package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ollema/sersophane/pkg/models"
)

type VenueModel struct {
	DB *pgxpool.Pool
}

func (m *VenueModel) Insert(name, city string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO venues (venue_name, venue_city) VALUES ($1, $2)`
	args := []interface{}{name, city}

	_, err := m.DB.Exec(ctx, query, args...)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "venues_venue_name_key" (SQLSTATE 23505)`:
			return models.ErrDuplicateName
		default:
			return err
		}
	}

	return nil
}

func (m *VenueModel) Get(id int) (*models.Venue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	venue := &models.Venue{}
	query := `SELECT venue_id, venue_name, venue_city FROM venues WHERE venue_id = $1`
	args := []interface{}{id}

	err := m.DB.QueryRow(ctx, query, args...).Scan(&venue.ID, &venue.Name, &venue.City)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return venue, nil
}

func (m *VenueModel) GetPage(filters *models.Filters) ([]*models.Venue, *models.Metadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	venues := []*models.Venue{}
	totalRecords := 0
	query := fmt.Sprintf(
		`SELECT count(*) OVER(), venue_id, venue_name, venue_city FROM venues
		ORDER BY %s %s LIMIT $1 OFFSET $2`,
		filters.SortBy,
		filters.SortDirection,
	)
	args := []interface{}{filters.Limit(), filters.Offset()}

	rows, err := m.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, &models.Metadata{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var venue models.Venue
		err := rows.Scan(&totalRecords, &venue.ID, &venue.Name, &venue.City)
		if err != nil {
			return nil, &models.Metadata{}, err
		}
		venues = append(venues, &venue)
	}

	if err = rows.Err(); err != nil {
		return nil, &models.Metadata{}, err
	}

	metadata := models.CalculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return venues, metadata, nil
}

func (m *VenueModel) GetAll() ([]*models.Venue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	venues := []*models.Venue{}
	query := `SELECT venue_id, venue_name, venue_city FROM venues ORDER BY venue_name ASC`

	args := []interface{}{}

	rows, err := m.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var venue models.Venue
		err := rows.Scan(&venue.ID, &venue.Name, &venue.City)
		if err != nil {
			return nil, err
		}
		venues = append(venues, &venue)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return venues, nil
}
