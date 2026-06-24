package chat

import (
	"context"
	"encoding/json"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SaveMessage(ctx context.Context, msg *Message) error {
	return s.repo.SaveMessage(ctx, msg)
}

func (s *Service) GetMessages(ctx context.Context, room string, limit int) ([]Message, error) {
	return s.repo.GetMessages(ctx, room, limit)
}

func parseWSMessage(data []byte) (*WSMessage, error) {
	var msg WSMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
