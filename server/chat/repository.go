package chat

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	db *pgxpool.Pool
}

// function for new repository
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
