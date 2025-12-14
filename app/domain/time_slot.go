package domain

import (
	"time"

	"github.com/google/uuid"
)

type TimeSlot struct {
	ID        uuid.UUID `json:"id"`
	ServiceID uuid.UUID `json:"service_id"` // Service
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	IsActive  string    `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
