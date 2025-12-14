package domain

import (
	"time"

	"github.com/alaa-aqeel/optional-value"
	"github.com/google/uuid"
)

type UsersMap map[uuid.UUID]*User

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // omit in JSON responses
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// relations
	Services []Services `json:"services"`
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
