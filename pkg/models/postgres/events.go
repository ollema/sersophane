package postgres

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ollema/sersophane/pkg/models"
)

type EventModel struct {
	DB *pgxpool.Pool
}

func (m *EventModel) Insert(name string, eventType models.EventType, start time.Time, end time.Time, artistIds []int, venueId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := m.DB.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	var eventId int
	event_query := `INSERT INTO events (event_name, event_type, event_start, event_end) VALUES ($1, $2, $3, $4) RETURNING event_id`
	event_args := []interface{}{name, eventType, start, end}

	err = tx.QueryRow(ctx, event_query, event_args...).Scan(&eventId)
	if err != nil {
		return err
	}

	for runningOrder, artistId := range artistIds {
		event_artist_query := `INSERT INTO event_artist (event_id, artist_id, event_artist_running_order) VALUES ($1, $2, $3)`
		event_artist_args := []interface{}{eventId, artistId, runningOrder}

		_, err = tx.Exec(ctx, event_artist_query, event_artist_args...)
		if err != nil {
			return err
		}
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
	query := `
		SELECT
			event_id,
			event_name,
			event_type,
			event_created_at,
			event_start,
			event_end,
			event_cancelled
		FROM events
		WHERE event_id = $1`
	args := []interface{}{id}

	err := m.DB.QueryRow(ctx, query, args...).Scan(&event.ID, &event.Name, &event.Type, &event.Start, &event.End, &event.Cancelled)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
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
		    count(*) OVER(),
			events.event_id,
			events.event_name,
			events.event_type,
			events.event_created_at,
			events.event_start,
			events.event_end,
			events.event_cancelled,
			ARRAY_AGG(artists.artist_id) AS artist_ids,
			ARRAY_AGG(artists.artist_name) AS artist_names,
			ARRAY_AGG(event_artist.event_artist_running_order) as running_order,
			venues.venue_id,
			venues.venue_name,
			venues.venue_city
		FROM events
		INNER JOIN event_artist ON events.event_id = event_artist.event_id
		INNER JOIN artists ON event_artist.artist_id = artists.artist_id
		INNER JOIN event_venue ON events.event_id = event_venue.event_id
		INNER JOIN venues ON event_venue.venue_id = venues.venue_id
		GROUP BY
			events.event_id,
			events.event_name,
			events.event_type,
			events.event_created_at,
			events.event_start,
			events.event_end,
			events.event_cancelled,
			venues.venue_id,
			venues.venue_name,
			venues.venue_city
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
		    count(*) OVER(),
			events.event_id,
			events.event_name,
			events.event_type,
			events.event_created_at,
			events.event_start,
			events.event_end,
			events.event_cancelled,
			ARRAY_AGG(artists.artist_id) AS artist_ids,
			ARRAY_AGG(artists.artist_name) AS artist_names,
			ARRAY_AGG(event_artist.event_artist_running_order) as running_order,
			venues.venue_id,
			venues.venue_name,
			venues.venue_city
		FROM events
		INNER JOIN event_artist ON events.event_id = event_artist.event_id
		INNER JOIN artists ON event_artist.artist_id = artists.artist_id
		INNER JOIN event_venue ON events.event_id = event_venue.event_id
		INNER JOIN venues ON event_venue.venue_id = venues.venue_id
		GROUP BY
			events.event_id,
			events.event_name,
			events.event_type,
			events.event_created_at,
			events.event_start,
			events.event_end,
			events.event_cancelled,
			venues.venue_id,
			venues.venue_name,
			venues.venue_city
		HAVING $1 = ANY(ARRAY_AGG(artists.artist_id))
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
		    count(*) OVER(),
			events.event_id,
			events.event_name,
			events.event_type,
			events.event_created_at,
			events.event_start,
			events.event_end,
			events.event_cancelled,
			ARRAY_AGG(artists.artist_id) AS artist_ids,
			ARRAY_AGG(artists.artist_name) AS artist_names,
			ARRAY_AGG(event_artist.event_artist_running_order) as running_order,
			venues.venue_id,
			venues.venue_name,
			venues.venue_city
		FROM events
		INNER JOIN event_artist ON events.event_id = event_artist.event_id
		INNER JOIN artists ON event_artist.artist_id = artists.artist_id
		INNER JOIN event_venue ON events.event_id = event_venue.event_id
		INNER JOIN venues ON event_venue.venue_id = venues.venue_id
		WHERE venues.venue_id = $1
		GROUP BY
			events.event_id,
			events.event_name,
			events.event_type,
			events.event_created_at,
			events.event_start,
			events.event_end,
			events.event_cancelled,
			venues.venue_id,
			venues.venue_name,
			venues.venue_city
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
		var runningOrder []int
		var artists []models.Artist

		err := rows.Scan(
			&totalRecords,
			&event.ID, &event.Name, &event.Type, &event.CreatedAt, &event.Start, &event.End, &event.Cancelled,
			&artistIds, &artistNames, &runningOrder,
			&event.Venue.ID, &event.Venue.Name, &event.Venue.City,
		)
		if err != nil {
			return nil, &models.Metadata{}, err
		}

		for i := range artistIds {
			artists = append(artists, models.Artist{
				ID:                artistIds[i],
				Name:              artistNames.Elements[i].String,
				EventRunningOrder: runningOrder[i],
			})
		}
		sort.Slice(artists, func(i, j int) bool { return artists[i].EventRunningOrder < artists[j].EventRunningOrder })
		event.Artists = artists

		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, &models.Metadata{}, err
	}

	metadata := models.CalculateMetadata(totalRecords, page, pageSize)

	return events, metadata, nil
}
