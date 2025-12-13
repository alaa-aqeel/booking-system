package services

import (
	"github.com/alaa-aqeel/booking-system/app/services/user_service"
	"github.com/alaa-aqeel/table"
)

type Services struct {
	User *user_service.UserService
}

func NewServices(db table.IDatabase) *Services {
	return &Services{
		User: user_service.NewUserService(db),
	}
}
