package users

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo Repository
}

type Service interface {
	RegisterUser(ctx context.Context, req RegisterUserRequest) error
	Login(ctx context.Context, req LoginUserRequest) (LoginUserResponse, error)
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) RegisterUser(ctx context.Context, req RegisterUserRequest) error {
	err := req.Validate()
	if err != nil {
		err = fmt.Errorf("error validate request: %w", err)
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("error generate password: %w", err)
		return err
	}

	u := &User{
		ID:       uuid.New(),
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}

	err = s.repo.RegisterUser(ctx, u)
	if err != nil {
		err = fmt.Errorf("error register user: %w", err)
		return err
	}

	return nil
}

func (s *service) Login(ctx context.Context, req LoginUserRequest) (LoginUserResponse, error) {
	err := req.Validate()
	if err != nil {
		return LoginUserResponse{}, err
	}

	user, err := s.repo.GetByEmail(ctx, req.Email)

	if err != nil {
		return LoginUserResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return LoginUserResponse{}, err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"id": user.ID,
	}
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return LoginUserResponse{}, err
	}

	return LoginUserResponse{
		User:  user,
		Token: tokenString,
	}, nil
}
