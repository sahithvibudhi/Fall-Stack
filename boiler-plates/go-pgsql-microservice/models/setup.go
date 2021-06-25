package models

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/helpers"
)

type Config struct {
	db *pg.DB
}

var dbInstance Config

func Setup() error {
	dbInstance = Config{}
	dbInstance.Connect()
	err := dbInstance.Migrate()
	if err != nil {
		helpers.Log.Error(fmt.Scanf("Failed to start db. Err: %v", err))
	}
	return err
}

func (c Config) Connect() {
	c.db = pg.Connect(&pg.Options{
		User: "postgres",
	})
}

func (c Config) Close() {
	c.db.Close()
}

func modelList() []interface{} {
	return []interface{}{
		(*User)(nil),
	}
}

func (c Config) Migrate() error {
	models := modelList()

	for _, model := range models {
		err := c.db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
