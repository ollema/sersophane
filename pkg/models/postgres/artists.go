package postgres

import (
	"database/sql"
	"errors"
	"fmt"

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
	artist := &models.Artist{}
	query := `SELECT id, name FROM artists WHERE id = $1`
	args := []interface{}{id}

	err := m.DB.QueryRow(query, args...).Scan(&artist.ID, &artist.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return artist, nil
}

func (m *ArtistModel) GetAll(filters *models.Filters) ([]*models.Artist, *models.Metadata, error) {
	artists := []*models.Artist{}
	totalRecords := 0
	query := fmt.Sprintf(
		`SELECT count(*) OVER(), id, name FROM artists
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
		var artist models.Artist
		err := rows.Scan(&totalRecords, &artist.ID, &artist.Name)
		if err != nil {
			return nil, &models.Metadata{}, err
		}
		artists = append(artists, &artist)
	}

	if err = rows.Err(); err != nil {
		return nil, &models.Metadata{}, err
	}

	metadata := models.CalculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return artists, metadata, nil
}
