package user_service

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/table"
)

func (s *UserService) Get(ids ...any) ([]domain.User, error) {
	rows, err := s.ITable.Get(context.Background(), squirrel.Eq{"id": ids})
	if err != nil {
		return nil, err
	}

	return table.ScanRows(rows, s.toUser)
}

func (s *UserService) Find(ids any, loads ...table.LoaderOne[domain.User]) (domain.User, error) {
	row, err := s.ITable.Find(context.Background(), "id", ids)
	if err != nil {
		return domain.User{}, err
	}

	user, err := s.toUser(row)
	if err != nil {
		return domain.User{}, err
	}

	if len(loads) == 0 {
		return user, nil
	}

	for _, load := range loads {
		user, err = load(user)
		if err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}

func (s *UserService) GetAll(dto domain.UserQuery) ([]domain.User, error) {
	page := dto.Page.ValueOrDefault(1)
	limit := dto.Limit.ValueOrDefault(10)

	filters := squirrel.And{}
	if dto.Username.IsSet && !dto.Username.IsEmpty() {
		filters = append(filters, squirrel.Like{"username": "%" + dto.Username.Value + "%"})
	}

	if dto.IsActive.IsSet {
		filters = append(filters, squirrel.Eq{"is_active": dto.IsActive.Value})
	}

	rows, err := s.Paginate(context.Background(),
		limit,
		(page-1)*limit,
		filters,
	)
	if err != nil {
		return nil, err
	}

	return table.ScanRows(rows, s.toUser)
}
