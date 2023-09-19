package terminal

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

type Repository interface {
	RegisterTerminal(ctx context.Context, t *Terminal) error
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db}
}

func (r *repository) RegisterTerminal(ctx context.Context, t *Terminal) error {
	db, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}

	defer db.Release()

	_, err = db.Exec(
		ctx,
		"INSERT INTO terminals (id, name, longitude, latitude) VALUES ($1, $2, $3, $4)",
		t.ID,
		t.Name,
		t.Longitude,
		t.Latitude,
	)

	if err != nil {
		return err
	}

	return nil
}
