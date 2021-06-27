package models

import (
	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/go-xorm/xorm/migrate"
	_ "github.com/lib/pq"
	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/helpers"
)

var DB *xorm.Engine

func Setup() error {
	err := connect()
	if err != nil {
		return err
	}

	err = PingDB()
	if err != nil {
		return err
	}

	err = Migrate()

	return err
}

func connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", helpers.Config.DBHost, helpers.Config.DBPort, helpers.Config.DBUser, helpers.Config.DBPass, helpers.Config.DBName)

	var err error
	DB, err = xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		helpers.Log.Error(fmt.Sprintf("Failed to connect to the db. Err: %v", err))
	}

	return err
}

func PingDB() error {
	err := DB.Ping()
	if err != nil {
		helpers.Log.Error(fmt.Sprintf("Ping to DB Failed. Err: %v", err))
	}

	return err
}

func Close() {
	DB.Close()
}

func Migrate() error {
	m := migrate.New(DB, &migrate.Options{
		TableName:    "migrations",
		IDColumnName: "id",
	}, migrations)

	err := m.Migrate()
	if err != nil {
		helpers.Log.Error(fmt.Sprintf("Error running migrations. Err: %v", err))
	}

	return err
}
