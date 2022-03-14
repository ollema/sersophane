package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ollema/sersophane/pkg/models"
)

type EventModel struct {
	DB *pgxpool.Pool
}

func (m *EventModel) Insert(name string, eventType models.EventType, startAt time.Time, endAt time.Time) error {
	query := `INSERT INTO events (name, type, start_at, end_at) VALUES($1, $2, $3, $4, $5)`
	args := []interface{}{name, eventType, startAt, endAt}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (m *EventModel) Get(id int) (*models.Event, error) {
	event := &models.Event{}
	query := `SELECT id, name, type, created_at, start_at, end_at, cancelled FROM events WHERE id = $1`
	args := []interface{}{id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

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

func (m *EventModel) GetAll(filters *models.Filters) ([]*models.Event, *models.Metadata, error) {
	events := []*models.Event{}
	totalRecords := 0
	query := fmt.Sprintf(
		`SELECT count(*) OVER(), id, name, type, created_at, start_at, end_at, cancelled FROM events
		ORDER BY %s %s LIMIT $1 OFFSET $2`,
		filters.SortBy,
		filters.SortDirection,
	)
	args := []interface{}{filters.Limit(), filters.Offset()}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, &models.Metadata{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&totalRecords, &event.ID, &event.Name, &event.Type, &event.CreatedAt, &event.StartAt, &event.EndAt, &event.Cancelled)
		if err != nil {
			return nil, &models.Metadata{}, err
		}
		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, &models.Metadata{}, err
	}

	metadata := models.CalculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return events, metadata, nil
}
