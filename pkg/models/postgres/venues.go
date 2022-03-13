package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ollema/sersophane/pkg/models"
)

type VenueModel struct {
	DB *sql.DB
}

func (m *VenueModel) Insert(name, city string) error {
	query := `INSERT INTO venues (name, city) VALUES ($1, $2)`
	args := []interface{}{name, city}

	_, err := m.DB.Exec(query, args...)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "venues_name_key"`:
			return models.ErrDuplicateName
		default:
			return err
		}
	}

	return nil
}

func (m *VenueModel) Get(id int) (*models.Venue, error) {
	venue := &models.Venue{}
	query := `SELECT id, name, city FROM venues WHERE id = $1`
	args := []interface{}{id}

	err := m.DB.QueryRow(query, args...).Scan(&venue.ID, &venue.Name, &venue.City)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return venue, nil
}

func (m *VenueModel) GetAll(filters *models.Filters) ([]*models.Venue, *models.Metadata, error) {
	venues := []*models.Venue{}
	totalRecords := 0
	query := fmt.Sprintf(
		`SELECT count(*) OVER(), id, name, city FROM venues
		ORDER BY %s %s LIMIT $1 OFFSET $2`,
		filters.SortBy,
		filters.SortDirection,
	)
	args := []interface{}{filters.Limit(), filters.Offset()}

	rows, err := m.DB.Query(query, args...)
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
