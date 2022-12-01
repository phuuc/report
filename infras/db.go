package infras

import (
	"database/sql"
	"log"
	"time"

	"github.com/finnpn/overview/config"

	"github.com/finnpn/overview/pkg/helpers"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DB struct {
	config *config.Config
}

func NewDB(config *config.Config) *DB {
	return &DB{
		config: config,
	}
}

func (d *DB) SetupDB() *sql.DB {
	var err error
	db, err := sql.Open(d.config.DB.DriverName, d.config.DB.DataSource)
	if err != nil {
		log.Fatalf("db could not open with err=%v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("db could not ping with err=%v", err)
	}
	db.SetMaxIdleConns(d.config.DB.MaxIdleConns)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Duration(d.config.DB.ConnMaxLifeTimeMiliseconds))

	log.Println("connected...")
	return db
}

func (d *DB) RunMigration() {
	m, err := migrate.New("file://"+helpers.RootDir()+"/db/migration", d.config.DB.MigrationConnURL)
	if err != nil {
		log.Fatalf("failed run migration with err=%v", err)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed run migration up with err=%v", err)
		return
	}
	log.Println("migration completed!")
}
