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

func (m *VenueModel) Insert(name string) error {
	query := `INSERT INTO venues (name) VALUES ($1)`
	args := []interface{}{name}

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
	a := &models.Venue{}
	query := `SELECT id, name FROM venues WHERE id = $1`
	args := []interface{}{id}

	err := m.DB.QueryRow(query, args...).Scan(&a.ID, &a.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return a, nil
}

func (m *VenueModel) GetAll(filters Filters) ([]*models.Venue, error) {
	query := fmt.Sprintf(
		`SELECT id, name FROM venues ORDER BY %s %s LIMIT $1 OFFSET $2`,
		filters.SortBy, filters.SortDirection)
	args := []interface{}{filters.limit(), filters.offset()}

	rows, err := m.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	venues := []*models.Venue{}
	for rows.Next() {
		var venue models.Venue
		err := rows.Scan(&venue.ID, &venue.Name)
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
