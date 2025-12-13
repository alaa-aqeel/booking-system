package user_service

import (
	"context"
	"errors"

	"github.com/alaa-aqeel/booking-system/app/domain"
)

func (s *UserService) Create(dto domain.CreateUserCommand) error {

	if !dto.Password.IsSet {
		return errors.New("password must be not null")
	}

	if !dto.Username.IsSet {
		return errors.New("username must be not null")
	}

	return s.Insert(
		context.Background(),
		s.NewUserMap(dto.Username.Value, dto.Password.Value, true),
	)
}
