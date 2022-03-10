package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ollema/sersophane/pkg/models"
)

type EventModel struct {
	DB *sql.DB
}

func (m *EventModel) Insert(name string, eventType models.EventType, startAt time.Time, endAt time.Time) error {
	query := `INSERT INTO events (name, type, start_at, end_at) VALUES($1, $2, $3, $4, $5)`
	args := []interface{}{name, eventType, startAt, endAt}

	_, err := m.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (m *EventModel) Get(id int) (*models.Event, error) {
	e := &models.Event{}
	query := `SELECT id, name, type, created_at, start_at, end_at, FROM events WHERE id = $1`
	args := []interface{}{id}

	err := m.DB.QueryRow(query, args...).Scan(&e.ID, &e.Name, &e.Type, &e.StartAt, &e.EndAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return e, nil
}
