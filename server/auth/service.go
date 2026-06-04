package auth

import (
	"context"
	"errors"
)

type Service struct {
	repo *Repository
}

func (s *Service) Login(ctx context.Context, email, password string) (string, *User, error) {

	user, err := s.repo.GetUserByUsername(ctx, email, password)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// generate the token

}
