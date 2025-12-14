package services_service

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/table"
)

func (s *ServicesService) GetAll(dto domain.ServicesQuery, loads ...table.LoaderMany[domain.Services]) ([]domain.Services, error) {
	page := dto.Page.ValueOrDefault(1)
	limit := dto.Limit.ValueOrDefault(10)

	filters := squirrel.And{}
	if dto.Name.IsSet && !dto.Name.IsEmpty() {
		filters = append(filters, squirrel.Like{"name": "%" + dto.Name.Value + "%"})
	}

	if dto.IsActive.IsSet {
		filters = append(filters, squirrel.Eq{"is_active": dto.IsActive.Value})
	}

	if dto.Price.IsSet {
		filters = append(filters, squirrel.Eq{"price": dto.Price.Value})
	}

	rows, err := s.Paginate(context.Background(),
		limit,
		(page-1)*limit,
		filters,
	)
	if err != nil {
		return nil, err
	}

	services, err := table.ScanRows(rows, s.toServices)

	if err != nil {
		return nil, err
	}

	if len(loads) == 0 {
		return services, nil
	}

	for _, load := range loads {
		services, err = load(services)
		if err != nil {
			return nil, err
		}
	}

	return services, nil
}

func (s *ServicesService) FindBy(key, id string, loads ...table.LoaderOne[domain.Services]) (domain.Services, error) {
	row, err := s.Find(context.Background(), key, id)
	if err != nil {
		return domain.Services{}, err
	}

	services, err := s.toServices(row)

	if err != nil {
		return domain.Services{}, err
	}

	if len(loads) == 0 {
		return services, nil
	}

	for _, load := range loads {
		services, err = load(services)
		if err != nil {
			return domain.Services{}, err
		}
	}

	return services, nil
}

func (s *ServicesService) GetBy(key string, ids []any) ([]domain.Services, error) {
	row, err := s.Get(context.Background(), map[string]any{
		key: ids,
	})
	if err != nil {
		return nil, err
	}

	return table.ScanRows(row, s.toServices)
}
