package migration

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/kaibling/iggy/pkg/config"
)

//go:embed migration_data/*.sql
var migrations embed.FS

func SelfMigrate(cfg config.DBConfig) error {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&x-migrations-table=\"public\".\"iggy_schema_migrations\"&x-migrations-table-quoted=1",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBDatabase,
	)

	p := &postgres.Postgres{}
	driver, err := p.Open(databaseURL)
	if err != nil {
		return err
	}

	defer driver.Close()
	d, err := iofs.New(migrations, "migration_data")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil

}

// func MigrateDB(prefix string, path string, cfg config.Configuration) error {

// 	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&x-migrations-table=\"public\".\"%s_schema_migrations\"&x-migrations-table-quoted=1",
// 		cfg.DBUser,
// 		cfg.DBPassword,
// 		cfg.DBHost,
// 		cfg.DBPort,
// 		cfg.DBDatabase,
// 		prefix,
// 	)

// 	p := &postgres.Postgres{}
// 	driver, err := p.Open(databaseURL)
// 	if err != nil {
// 		return err
// 	}
// 	defer driver.Close()
// 	// Create a new migration instance
// 	m, err := migrate.NewWithDatabaseInstance(
// 		"file://"+path,
// 		"postgres",
// 		driver,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("could not create migration instance: %v", err)
// 	}

// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		return fmt.Errorf("could not apply migrations: %v", err)
// 	}

// 	return nil
// }
