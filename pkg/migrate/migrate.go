package migrate

import (
	"fmt"
	"log"

	"github.com/edward-/four-in-a-row-game/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var migrationFolder = "file://./migrations"

func Up() error {
	m := build()

	err := m.Up()
	if err != nil {
		panic(err)
	}
	fmt.Println("migration UP completed")

	return err
}

func Down() error {
	m := build()
	
	err := m.Up()
	if err != nil {
		panic(err)
	}
	fmt.Println("migration DOWN completed")
	
	return err
}

func build() *migrate.Migrate {
	configuration := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		configuration.Db.Host,
		configuration.Db.User,
		configuration.Db.Password,
		configuration.Db.DBName,
		configuration.Db.Port,
		configuration.Db.SSLMode,
		configuration.Db.TimeZone,
	)

	db, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	instanceDB, err := db.DB()
	if err != nil {
		panic("failed to get database instance")
	}

	driver, err := postgres.WithInstance(instanceDB, &postgres.Config{})
	if err != nil {
		panic(fmt.Sprintf("Could not start sql migration... %v", err))
	}

	m, err := migrate.NewWithDatabaseInstance(migrationFolder, configuration.Db.DBName, driver)
	if err != nil {
		log.Fatal(err)
	}

	return m
}
