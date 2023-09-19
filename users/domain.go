package users

import (
	"time"

	"github.com/arfan21/mkpmobile/errors"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (r *RegisterUserRequest) Validate() error {
	validationErrors := errors.ValidationError{}
	if r.Username == "" {
		validationErrors.Errors = append(validationErrors.Errors, "username is required")
	}

	if r.Password == "" {
		validationErrors.Errors = append(validationErrors.Errors, "password is required")
	}

	if r.Email == "" {
		validationErrors.Errors = append(validationErrors.Errors, "email is required")
	}

	if len(validationErrors.Errors) > 0 {
		return &validationErrors
	}

	return nil
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginUserRequest) Validate() error {
	validationErrors := errors.ValidationError{}
	if r.Email == "" {
		validationErrors.Errors = append(validationErrors.Errors, "email is required")
	}

	if r.Password == "" {
		validationErrors.Errors = append(validationErrors.Errors, "password is required")
	}
	if len(validationErrors.Errors) > 0 {
		return &validationErrors
	}

	return nil
}

type LoginUserResponse struct {
	User
	Token string `json:"token"`
}
