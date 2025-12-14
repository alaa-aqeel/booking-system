package domain

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID          uuid.UUID `json:"id"`
	Fullname    string    `json:"fullname"`
	PhoneNumber string    `json:"phone_number"`
	ServiceID   uuid.UUID `json:"service_id"`
	TimeSlotID  uuid.UUID `json:"time_slot_id"`

	Status    string    `json:"status"` // pending, confirmed, canceled
	Notes     string    `json:"notes,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
