package services_service

import (
	"context"
	"errors"
	"time"

	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/table"
	"github.com/google/uuid"
)

type UserService interface {
	Get(ids ...any) ([]domain.User, error)
}

type ServicesService struct {
	table.ITable
	userService UserService
}

func NewServicesService(db table.IDatabase, userService UserService) *ServicesService {
	return &ServicesService{
		ITable: table.
			Table(db, "services", "id").
			Column("id, name, description, price, is_active, created_by, created_at, updated_at"),
		userService: userService,
	}
}

func (s *ServicesService) Create(dto domain.CreateServicesCommand) error {

	if !dto.Name.IsSet || dto.CreatedBy.IsEmpty() {
		return errors.New("name must be not null")
	}

	if !dto.CreatedBy.IsSet || dto.CreatedBy.IsEmpty() {
		return errors.New("create by id must be not null")
	}

	return s.Insert(context.Background(), map[string]any{
		"id":          uuid.NewString(),
		"name":        dto.Name.Value,
		"created_by":  dto.CreatedBy.Value,
		"description": dto.Description.ValueOrDefault(""),
		"price":       dto.Price.ValueOrDefault(0.0),
		"is_active":   dto.IsActive.ValueOrDefault(true),
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	})
}

func (s *ServicesService) Update(id string, dto domain.UpdateServicesCommand) error {

	data := map[string]any{}

	if dto.Name.IsSet && !dto.Name.IsEmpty() {
		data["name"] = dto.Name.Value
	}

	if dto.Description.IsSet {
		data["description"] = dto.Description.Value
	}

	if dto.Price.IsSet {
		data["price"] = dto.Price.Value
	}

	if dto.IsActive.IsSet {
		data["is_active"] = dto.IsActive.Value
	}

	if len(data) == 0 {
		return nil
	}

	data["updated_at"] = time.Now()

	return s.UpdatePk(context.Background(), id, data)
}
