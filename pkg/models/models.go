package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching records found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrDuplicateName      = errors.New("models: duplicate name")
	ErrInvalidFilters     = errors.New("models: invalid query filters")
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
	Cancelled bool

	// relationships
	Artists []Artist // many-to-many
	Users   []User   // many-to-many
	Venue   Venue    // many-to-one
}

type Artist struct {
	ID   int
	Name string
}

type Venue struct {
	ID   int
	Name string
	City string
}

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Email     string
	Password  []byte
	Activated bool
}
