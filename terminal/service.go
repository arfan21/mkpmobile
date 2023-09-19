package terminal

import (
	"context"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

type Service interface {
	RegisterTerminal(ctx context.Context, req RegisterTerminalRequest) error
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) RegisterTerminal(ctx context.Context, req RegisterTerminalRequest) error {
	err := req.Validate()
	if err != nil {
		return err
	}

	u := &Terminal{
		ID:        uuid.New(),
		Name:      req.Name,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	}

	err = s.repo.RegisterTerminal(ctx, u)
	if err != nil {
		return err
	}

	return nil
}
