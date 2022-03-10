package postgres

import (
	"database/sql"
	"errors"

	"github.com/ollema/sersophane/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	query := `INSERT INTO users (name, email, password_hash, active) VALUES ($1, $2, $3, true)`
	args := []interface{}{name, email, hashedPassword}

	_, err = m.DB.Exec(query, args...)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return models.ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_name_key"`:
			return models.ErrDuplicateName
		default:
			return err
		}
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	query := `SELECT id, password_hash FROM users WHERE email = $1 AND active = TRUE`
	args := []interface{}{email}

	row := m.DB.QueryRow(query, args...)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	u := &models.User{}
	query := `SELECT id, name, created_at, email, active FROM users WHERE id = $1`
	args := []interface{}{id}

	err := m.DB.QueryRow(query, args...).Scan(&u.ID, &u.Name, &u.CreatedAt, &u.Email, &u.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return u, nil
}
