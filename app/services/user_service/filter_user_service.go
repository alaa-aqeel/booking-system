package user_service

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/table"
)

func (s *UserService) Find(ids ...string) (domain.User, error) {
	row, err := s.One(context.Background(), "id", ids)
	if err != nil {
		return domain.User{}, err
	}

	return s.toUser(row)
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

	rows, err := s.Filter(context.Background(),
		limit,
		(page-1)*limit,
		filters,
	)
	if err != nil {
		return nil, err
	}

	return table.ScanRows(rows, s.toUser)
}
