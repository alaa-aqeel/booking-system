package services

import (
	"github.com/alaa-aqeel/booking-system/app/services/services_service"
	"github.com/alaa-aqeel/booking-system/app/services/user_service"
	"github.com/alaa-aqeel/table"
)

type Services struct {
	User     *user_service.UserService
	Services *services_service.ServicesService
}

func NewServices(db table.IDatabase) *Services {
	user := user_service.NewUserService(db)
	services := services_service.NewServicesService(db, user)
	user.SetServices(services)

	return &Services{
		User:     user,
		Services: services,
	}
}
