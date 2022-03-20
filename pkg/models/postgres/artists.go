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

type ArtistModel struct {
	DB *pgxpool.Pool
}

func (m *ArtistModel) Insert(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO artists (artist_name) VALUES ($1)`
	args := []interface{}{name}

	_, err := m.DB.Exec(ctx, query, args...)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "artists_artist_name_key" (SQLSTATE 23505)`:
			return models.ErrDuplicateName
		default:
			return err
		}
	}

	return nil
}

func (m *ArtistModel) Get(id int) (*models.Artist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	artist := &models.Artist{}
	query := `SELECT artist_id, artist_name FROM artists WHERE artist_id = $1`
	args := []interface{}{id}

	err := m.DB.QueryRow(ctx, query, args...).Scan(&artist.ID, &artist.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return artist, nil
}

func (m *ArtistModel) GetPage(filters *models.Filters) ([]*models.Artist, *models.Metadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	artists := []*models.Artist{}
	totalRecords := 0
	query := fmt.Sprintf(
		`SELECT count(*) OVER(), artist_id, artist_name FROM artists
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

func (m *ArtistModel) GetAll(filters *models.Filters) ([]*models.Artist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	artists := []*models.Artist{}
	query := fmt.Sprintf(
		`SELECT artist_id, artist_name FROM artists
		ORDER BY %s %s`,
		filters.SortBy,
		filters.SortDirection,
	)
	args := []interface{}{}

	rows, err := m.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var artist models.Artist
		err := rows.Scan(&artist.ID, &artist.Name)
		if err != nil {
			return nil, err
		}
		artists = append(artists, &artist)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return artists, nil
}
