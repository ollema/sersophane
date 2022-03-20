package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ollema/sersophane/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(name, email, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	query := `INSERT INTO users (user_name, user_email, user_password_hash, user_activated) VALUES ($1, $2, $3, TRUE)`
	args := []interface{}{name, email, hashedPassword}

	_, err = m.DB.Exec(ctx, query, args...)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "users_user_name_key" (SQLSTATE 23505)`:
			return models.ErrDuplicateName
		case err.Error() == `ERROR: duplicate key value violates unique constraint "users_user_email_key" (SQLSTATE 23505)`:
			return models.ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var hashedPassword []byte
	query := `SELECT user_id, user_password_hash FROM users WHERE user_email = $1 AND user_activated = TRUE`
	args := []interface{}{email}

	row := m.DB.QueryRow(ctx, query, args...)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	u := &models.User{}
	query := `SELECT user_id, user_name, user_created_at, user_email, user_activated FROM users WHERE user_id = $1`
	args := []interface{}{id}

	err := m.DB.QueryRow(ctx, query, args...).Scan(&u.ID, &u.Name, &u.CreatedAt, &u.Email, &u.Activated)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return u, nil
}
