package user_service

import (
	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/table"
)

type IServicesService interface {
	GetBy(key string, ids []any) ([]domain.Services, error)
}

type UserService struct {
	table.ITable
	services IServicesService
}

func NewUserService(db table.IDatabase) *UserService {
	return &UserService{
		ITable: table.
			Table(db, "users", "id").
			Column("id, username, is_active, password, created_at, updated_at"),
	}
}

func (s *UserService) SetServices(services IServicesService) {
	s.services = services
}
