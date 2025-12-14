package services_service

import (
	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/table"
)

func (s *ServicesService) toServices(row table.IRow) (domain.Services, error) {
	var service domain.Services
	err := row.Scan(
		&service.ID,
		&service.Name,
		&service.Description,
		&service.Price,
		&service.IsActive,
		&service.CreatedBy,
		&service.CreatedAt,
		&service.UpdatedAt,
	)

	return service, err
}
