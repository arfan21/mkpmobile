package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

type Repository interface {
	RegisterUser(ctx context.Context, u *User) error
	GetByEmail(ctx context.Context, email string) (User, error)
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db}
}

func (r *repository) RegisterUser(ctx context.Context, u *User) error {
	db, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}

	defer db.Release()

	_, err = db.Exec(
		ctx,
		"INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4)",
		u.ID,
		u.Username,
		u.Password,
		u.Email,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (User, error) {
	db, err := r.db.Acquire(ctx)
	if err != nil {
		return User{}, err
	}

	defer db.Release()

	var user User

	err = db.QueryRow(
		ctx,
		"SELECT id, username, password, email, created_at, updated_at FROM users WHERE email = $1",
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
