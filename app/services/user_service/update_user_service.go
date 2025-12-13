package user_service

import (
	"context"

	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/booking-system/shared"
)

func (s *UserService) Update(id string, dto domain.UpdateUserCommand) error {

	data := map[string]any{}

	if dto.Username.IsSet {
		data["username"] = dto.Username.Value
	}

	if dto.Password.IsSet {
		hashed, err := shared.MakeHash(dto.Password.Value)
		if err != nil {
			return err
		}
		data["password"] = hashed
	}

	if dto.IsActive.IsSet {
		data["is_active"] = dto.IsActive
	}

	if len(data) == 0 {
		return nil
	}

	return s.UpdatePk(context.Background(), id, data)
}
