package postgres

import (
	"database/sql"
	"errors"

	"github.com/ollema/sersophane/pkg/models"
)

type ArtistModel struct {
	DB *sql.DB
}

func (m *ArtistModel) Insert(name string) error {
	query := `INSERT INTO artists (name) VALUES ($1)`
	args := []interface{}{name}

	_, err := m.DB.Exec(query, args...)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "artists_name_key"`:
			return models.ErrDuplicateName
		default:
			return err
		}
	}

	return nil
}

func (m *ArtistModel) Get(id int) (*models.Artist, error) {
	a := &models.Artist{}
	query := `SELECT id, name FROM artists WHERE id = $1`
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
