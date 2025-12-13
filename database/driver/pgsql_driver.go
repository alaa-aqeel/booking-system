package database

import (
	"context"
	"database/sql"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	database *Database
	onceDb   sync.Once
)

type Database struct {
	sqlDb *sql.DB
}

func NewDatabase() *Database {
	onceDb.Do(func() {
		database = &Database{}
	})

	return database
}

func (p *Database) Connect(ctx context.Context, dsn string) error {
	sql, err := sql.Open("pgx", dsn)
	if err != nil {
		return MapPgError(err)
	}
	p.sqlDb = sql

	return nil
}

func (s *Database) Close() {

	s.sqlDb.Close()
}

func (s *Database) Db() *sql.DB {

	return s.sqlDb
}

func (s *Database) QueryRow(ctx context.Context, sql string, args ...any) *sql.Row {

	return s.Db().QueryRowContext(ctx, sql, args...)
}

func (s *Database) Query(ctx context.Context, sql string, args ...any) (*sql.Rows, error) {
	row, err := s.Db().QueryContext(ctx, sql, args...)

	return row, MapPgError(err)
}

// Exec executes a statement with arguments
func (s *Database) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := s.Db().ExecContext(ctx, sql, args...)

	return MapPgError(err)
}
