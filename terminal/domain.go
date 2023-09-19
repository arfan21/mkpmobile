package terminal

import (
	"time"

	"github.com/arfan21/mkpmobile/errors"
	"github.com/google/uuid"
)

type Terminal struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterTerminalRequest struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (r *RegisterTerminalRequest) Validate() error {
	validationErrors := errors.ValidationError{}
	if r.Name == "" {
		validationErrors.Errors = append(validationErrors.Errors, "name is required")
	}

	if r.Longitude == 0 {
		validationErrors.Errors = append(validationErrors.Errors, "longitude is required")
	}

	if r.Latitude == 0 {
		validationErrors.Errors = append(validationErrors.Errors, "latitude is required")
	}
	if len(validationErrors.Errors) > 0 {
		return &validationErrors
	}

	return nil
}
