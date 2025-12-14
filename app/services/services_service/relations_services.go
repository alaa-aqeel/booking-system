package services_service

import (
	"context"

	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/table"
)

func (s *ServicesService) GetBy(key string, ids []any) ([]domain.Services, error) {
	row, err := s.Get(context.Background(), map[string]any{
		key: ids,
	})
	if err != nil {
		return nil, err
	}

	return table.ScanRows(row, s.toServices)
}

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
