package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	dbname   = "remotetododb"
	port     = "5432"
	user     = "postgres"
	password = "8055"
)

var Data *sqlx.DB

type SSLMode string

const (
	//SSLModeEnable  SSLMode = "enable"
	SSLModeDisable SSLMode = "disable"
)

func Connect() error {

	start := fmt.Sprintf("host=%s port=%s user=%s password=%s  dbname=%s sslmode=%s", host, port, user, password, dbname, SSLModeDisable)
	db, err := sqlx.Open("postgres", start)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	Data = db
	migerr := migrateStart(db)
	if migerr != nil {
		fmt.Println("migration err: +v", migerr)
	}

	return nil
}

func migrateStart(db *sqlx.DB) error {

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, NewErr := migrate.NewWithDatabaseInstance("file://database/migration", "postgres", driver)
	if NewErr != nil {
		return NewErr
	}
	if MigrateErr := m.Up(); MigrateErr != nil && MigrateErr != migrate.ErrNoChange { //up(): will migrate all the way up
		return MigrateErr
	}
	return nil
}
