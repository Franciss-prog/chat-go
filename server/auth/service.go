package auth

import (
	"api/ws/middleware"
	"context"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Login(ctx context.Context, email, password string) (string, *User, error) {

	user, err := s.repo.GetUserByUsername(ctx, email, password)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// generate the token
	token, err := middleware.GenerateToken(user.ID, user.Username)

	if err != nil {
		return "", nil, err
	}

	return token, user, nil

}

func (s *Service) Register(ctx context.Context, request RegisterRequest) (string, string, error) {
	id, err := s.repo.RegisterUser(ctx, request.Username, request.Email, request.Password)
	if err != nil {
		return "", "", err
	}
	token, err := middleware.GenerateToken(id, request.Username)
	if err != nil {
		return "", "", err
	}
	return token, id, nil
}
