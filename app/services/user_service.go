package services

import "github.com/alaa-aqeel/table"

type UserService struct {
	db *table.SqlTable
}

func NewUserService(db table.IDatabase) *UserService {
	return &UserService{
		db: table.Table(db, "users", "id"),
	}
}
