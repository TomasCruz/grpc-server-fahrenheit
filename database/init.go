package database

import (
	"database/sql"
	"log"

	"github.com/TomasCruz/grpc-server-fahrenheit/model"

	"github.com/golang-migrate/migrate/v4"
	database "github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// I need this, linter
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type postgresDb struct {
	db *sql.DB
}

var pDb postgresDb

// InitializeDatabase does DB migrations and verifies DB accessibility
func InitializeDatabase(dbString string) model.Database {
	var err error
	var db *sql.DB

	if db, err = sql.Open("postgres", dbString); err != nil {
		log.Fatal(err)
	}

	var driver database.Driver
	if driver, err = postgres.WithInstance(db, &postgres.Config{}); err != nil {
		log.Fatal(err)
	}

	var m *migrate.Migrate
	if m, err = migrate.NewWithDatabaseInstance("file://database/migrations", "", driver); err != nil {
		log.Fatal(err)
	}

	if err = m.Up(); err == migrate.ErrNoChange {
		err = nil
	}

	sErr, dbErr := m.Close()
	if sErr != nil {
		log.Fatal(sErr)
	} else if dbErr != nil {
		log.Fatal(dbErr)
	}

	if db, err = sql.Open("postgres", dbString); err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	pDb = postgresDb{db: db}
	return pDb
}
