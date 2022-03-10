package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching records found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type EventType string

const (
	Concert  EventType = "concert"
	Festival EventType = "festival"
	Film     EventType = "film"
)

type Event struct {
	ID        int
	Name      string
	Type      EventType
	CreatedAt time.Time
	StartAt   time.Time
	EndAt     time.Time
	Venue     string
}

type Artist struct {
	ID   int
	Name string
}

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Email     string
	Password  []byte
	Active    bool
}
