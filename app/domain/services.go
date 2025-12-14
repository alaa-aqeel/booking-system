package domain

import (
	"time"

	"github.com/alaa-aqeel/optional-value"
	"github.com/google/uuid"
)

type Services struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	IsActive    bool      `json:"is_active"`
	CreatedBy   uuid.UUID `json:"create_by"` // User
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Creator User
}

type CreateServicesCommand struct {
	Name        optional.Optional[string]  `json:"name"`
	Description optional.Optional[string]  `json:"description"`
	Price       optional.Optional[float64] `json:"price"`
	IsActive    optional.Optional[bool]    `json:"is_active"`
	CreatedBy   optional.Optional[string]  `json:"create_by"` // User
}

type UpdateServicesCommand struct {
	Name        optional.Optional[string]  `json:"name"`
	Description optional.Optional[string]  `json:"description"`
	Price       optional.Optional[float64] `json:"price"`
	IsActive    optional.Optional[bool]    `json:"is_active"`
	CreatedBy   optional.Optional[string]  `json:"create_by"` // User
}

type ServicesQuery struct {
	Limit optional.Optional[int]
	Page  optional.Optional[int]

	Name      optional.Optional[string]  `json:"name"`
	IsActive  optional.Optional[bool]    `json:"is_active"`
	Price     optional.Optional[float64] `json:"price"`
	CreatedBy optional.Optional[string]  `json:"create_by"`
}
