package services_service

import (
	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/table"
)

func (s *ServicesService) LoadUser(services []domain.Services) ([]domain.Services, error) {

	if len(services) == 0 {
		return nil, nil
	}

	return table.OneToOne(
		services,
		func(service domain.Services) any {
			return service.CreatedBy
		},
		func(ids []any) ([]domain.User, error) {
			return s.userService.Get(ids...)
		},
		func(u domain.User) any {
			return u.ID
		},
		func(service *domain.Services, user domain.User) {
			service.Creator = user
		},
	)
}

func (s *ServicesService) LoadUserOne(service domain.Services) (domain.Services, error) {

	services, err := s.LoadUser([]domain.Services{service})
	if err != nil {
		return domain.Services{}, err
	}

	return services[0], nil
}
