package chat

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveMessage(ctx context.Context, msg *Message) error {
	msg.CreatedAt = time.Now()
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO messages (room, sender_id, username, content)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, created_at`,
		msg.Room, msg.SenderID, msg.Username, msg.Content,
	).Scan(&msg.ID, &msg.CreatedAt)
	return err
}

func (r *Repository) GetMessages(ctx context.Context, room string, limit int) ([]Message, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, room, sender_id, username, content, created_at
		 FROM messages
		 WHERE room = $1
		 ORDER BY created_at DESC
		 LIMIT $2`,
		room, limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.Room, &msg.SenderID, &msg.Username, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
