package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ollema/sersophane/pkg/models"
)

type EventModel struct {
	DB *pgxpool.Pool
}

func (m *EventModel) Insert(name string, eventType models.EventType, startAt time.Time, endAt time.Time, artistId, venueId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	var eventId int
	event_query := `INSERT INTO events (name, type, start_at, end_at, cancelled) VALUES ($1, $2, $3, $4, FALSE) RETURNING id`
	event_args := []interface{}{name, eventType, startAt, endAt}

	err = tx.QueryRow(ctx, event_query, event_args...).Scan(&eventId)
	if err != nil {
		return err
	}

	event_artist_query := `INSERT INTO event_artist (event_id, artist_id) VALUES ($1, $2)`
	event_artist_args := []interface{}{eventId, artistId}

	_, err = tx.Exec(ctx, event_artist_query, event_artist_args...)
	if err != nil {
		return err
	}

	event_venue_query := `INSERT INTO event_venue (event_id, venue_id) VALUES ($1, $2)`
	event_venue_args := []interface{}{eventId, venueId}

	_, err = tx.Exec(ctx, event_venue_query, event_venue_args...)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *EventModel) Get(id int) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	event := &models.Event{}
	query := `SELECT id, name, type, created_at, start_at, end_at, cancelled FROM events WHERE id = $1`
	args := []interface{}{id}

	err := m.DB.QueryRow(ctx, query, args...).Scan(&event.ID, &event.Name, &event.Type, &event.StartAt, &event.EndAt, &event.Cancelled)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return event, nil
}

func (m *EventModel) GetPage(filters *models.Filters) ([]*models.Event, *models.Metadata, error) {
	query := fmt.Sprintf(
		`SELECT
		    count(*) OVER(), events.id, events.name, events.type, events.created_at, events.start_at, events.end_at, events.cancelled,
		    ARRAY_AGG(artists.id) AS artist_ids, ARRAY_AGG(artists.name) AS artist_names,
		    venues.id AS venue_id, venues.name AS venue_name
		FROM events
		INNER JOIN event_artist ON events.id = event_artist.event_id
		INNER JOIN artists ON event_artist.artist_id = artists.id
		INNER JOIN event_venue ON events.id = event_venue.event_id
		INNER JOIN venues ON event_venue.venue_id = venues.id
		GROUP BY
			events.id, events.name, events.type, events.created_at, events.start_at, events.end_at, events.cancelled,
			venues.id, venues.name
		ORDER BY %s %s
		LIMIT $1 OFFSET $2`,
		filters.SortBy,
		filters.SortDirection,
	)
	args := []interface{}{filters.Limit(), filters.Offset()}

	return m.getPage(query, args, filters.Page, filters.PageSize)
}

func (m *EventModel) GetPageForArtist(artistId int, filters *models.Filters) ([]*models.Event, *models.Metadata, error) {
	query := fmt.Sprintf(
		`SELECT
		    count(*) OVER(), events.id, events.name, events.type, events.created_at, events.start_at, events.end_at, events.cancelled,
		    ARRAY_AGG(artists.id) AS artist_ids, ARRAY_AGG(artists.name) AS artist_names,
		    venues.id AS venue_id, venues.name AS venue_name
		FROM events
		INNER JOIN event_artist ON events.id = event_artist.event_id
		INNER JOIN artists ON event_artist.artist_id = artists.id
		INNER JOIN event_venue ON events.id = event_venue.event_id
		INNER JOIN venues ON event_venue.venue_id = venues.id
		GROUP BY
			events.id, events.name, events.type, events.created_at, events.start_at, events.end_at, events.cancelled,
			venues.id, venues.name
		HAVING $1 = ANY(ARRAY_AGG(artists.id))
		ORDER BY %s %s
		LIMIT $2 OFFSET $3`,
		filters.SortBy,
		filters.SortDirection,
	)
	args := []interface{}{artistId, filters.Limit(), filters.Offset()}

	return m.getPage(query, args, filters.Page, filters.PageSize)
}

func (m *EventModel) GetPageForVenue(venueId int, filters *models.Filters) ([]*models.Event, *models.Metadata, error) {
	query := fmt.Sprintf(
		`SELECT
		    count(*) OVER(), events.id, events.name, events.type, events.created_at, events.start_at, events.end_at, events.cancelled,
		    ARRAY_AGG(artists.id) AS artist_ids, ARRAY_AGG(artists.name) AS artist_names,
		    venues.id AS venue_id, venues.name AS venue_name
		FROM events
		INNER JOIN event_artist ON events.id = event_artist.event_id
		INNER JOIN artists ON event_artist.artist_id = artists.id
		INNER JOIN event_venue ON events.id = event_venue.event_id
		INNER JOIN venues ON event_venue.venue_id = venues.id
		WHERE venues.id = $1
		GROUP BY
			events.id, events.name, events.type, events.created_at, events.start_at, events.end_at, events.cancelled,
			venues.id, venues.name
		ORDER BY %s %s
		LIMIT $2 OFFSET $3`,
		filters.SortBy,
		filters.SortDirection,
	)
	args := []interface{}{venueId, filters.Limit(), filters.Offset()}

	return m.getPage(query, args, filters.Page, filters.PageSize)
}

func (m *EventModel) getPage(query string, args []interface{}, page, pageSize int) ([]*models.Event, *models.Metadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	events := []*models.Event{}
	totalRecords := 0

	rows, err := m.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, &models.Metadata{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		var artistIds []int
		var artistNames pgtype.TextArray
		var artists []models.Artist

		err := rows.Scan(
			&totalRecords,
			&event.ID, &event.Name, &event.Type, &event.CreatedAt, &event.StartAt, &event.EndAt, &event.Cancelled,
			&artistIds, &artistNames,
			&event.Venue.ID, &event.Venue.Name,
		)
		if err != nil {
			return nil, &models.Metadata{}, err
		}

		for i := range artistIds {
			artists = append(artists, models.Artist{ID: artistIds[i], Name: artistNames.Elements[i].String})
		}
		// sort according to junction table insertion order, might be hard. probably need another column
		// sort.Slice(artists, func(i, j int) bool { return ... < ... })
		event.Artists = artists

		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, &models.Metadata{}, err
	}

	metadata := models.CalculateMetadata(totalRecords, page, pageSize)

	return events, metadata, nil
}
