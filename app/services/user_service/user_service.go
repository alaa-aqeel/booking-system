package user_service

import (
	"github.com/alaa-aqeel/table"
)

type UserService struct {
	table.ITable
}

func NewUserService(db table.IDatabase) *UserService {
	return &UserService{
		table.
			Table(db, "users", "id").
			Column("id, username, is_active, password, created_at, updated_at"),
	}
}
