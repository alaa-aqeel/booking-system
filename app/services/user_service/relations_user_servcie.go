package user_service

import (
	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/table"
)

func (s *UserService) LoadServices(users []domain.User) ([]domain.User, error) {

	return table.OneToMany(
		users,
		func(user domain.User) any {
			return user.ID
		}, // fk
		func(ids []any) ([]domain.Services, error) {
			return s.services.GetBy("created_by", ids)
		}, // load user id
		func(services domain.Services) any {
			return services.CreatedBy
		}, // ref id ref
		func(user *domain.User, services []domain.Services) {
			user.Services = services
		},
	)
}

func (s *UserService) LoadServicesOne(user domain.User) (domain.User, error) {

	users, err := s.LoadServices([]domain.User{user})
	if err != nil {
		return domain.User{}, err
	}

	return users[0], nil
}
