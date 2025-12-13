package user_service

import (
	"time"

	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/booking-system/shared"
	"github.com/alaa-aqeel/table"
	"github.com/google/uuid"
)

func (s *UserService) NewUserMap(Username, Password string, IsActive bool) map[string]any {
	hahsed, _ := shared.MakeHash(Password)

	return map[string]any{
		"id":         uuid.NewString(),
		"username":   Username,
		"password":   hahsed,
		"is_active":  IsActive,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
}

func (s *UserService) toUser(row table.IRow) (domain.User, error) {
	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.IsActive,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}
