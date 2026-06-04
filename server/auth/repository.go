package auth

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

// variable for
var (
	ErrUsernameExists = errors.New("username exists")
	ErrEmailExists    = errors.New("email exists")
)

// function for new repository
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

// repository for testing
func (r *Repository) CheckVersion(ctx context.Context) (string, error) {
	var version string
	if err := r.db.QueryRow(ctx, "SELECT VERSION()").Scan(&version); err != nil {
		return "", err
	}
	return version, nil

}

// repository for login
func (r *Repository) GetUserByUsername(ctx context.Context, email, password string) (*User, error) {

	var user User

	err := r.db.QueryRow(
		ctx,
		"SELECT id,username FROM users WHERE email=$1 AND password_hash = crypt($2, password_hash)",
		email, password,
	).Scan(&user.ID, &user.Username)

	return &user, err
}

func (r *Repository) RegisterUser(
	ctx context.Context,
	username, email, password string,
) (string, error) {

	var id string

	err := r.db.QueryRow(
		ctx,
		`
        INSERT INTO users(username, email, password_hash)
        VALUES($1, $2, crypt($3, gen_salt('bf')))
        RETURNING id
        `,
		username,
		email,
		password,
	).Scan(&id)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) &&
			pgErr.Code == "23505" {

			switch pgErr.ConstraintName {
			case "users_username_key":
				return "", ErrUsernameExists

			case "users_email_key":
				return "", ErrEmailExists
			}
		}

		return "", err
	}

	return id, nil
}
