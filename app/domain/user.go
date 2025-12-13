package domain

import (
	"time"

	"github.com/alaa-aqeel/optional-value"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // omit in JSON responses
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserCommand struct {
	Username optional.Optional[string]
	Password optional.Optional[string]
}

type UpdateUserCommand struct {
	Username optional.Optional[string]
	Password optional.Optional[string]
	IsActive optional.Optional[bool]
}

type UserQuery struct {
	Limit    optional.Optional[int]
	Page     optional.Optional[int]
	Username optional.Optional[string]
	IsActive optional.Optional[bool]
}
