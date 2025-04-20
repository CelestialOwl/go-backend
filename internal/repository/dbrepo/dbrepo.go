package dbrepo

import (
	"database/sql"

	"github.com/tsawler/bookings-app/internal/config"
	"github.com/tsawler/bookings-app/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type TestDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, app *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: app,
		DB:  conn,
	}
}

func NewTestingRepo(app *config.AppConfig) repository.DatabaseRepo {
	return &TestDBRepo{
		App: app,
	}
}
