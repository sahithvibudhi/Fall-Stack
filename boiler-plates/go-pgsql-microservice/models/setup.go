package models

import (
	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/go-xorm/xorm/migrate"
	_ "github.com/lib/pq"
	"github.com/sahithvibudhi/full-stack/boiler-plates/go-pgsql-microservice/helpers"
)

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
	Engine   *xorm.Engine
}

var DBInstance *Config

func Setup() error {
	DBInstance = &Config{
		host:     "127.0.0.1",
		port:     "5432",
		user:     "postgres",
		password: "postgres",
		dbName:   "boilerplate",
	}
	err := DBInstance.Connect()
	if err != nil {
		return err
	}

	err = DBInstance.PingDB()
	if err != nil {
		return err
	}

	err = DBInstance.Migrate()

	return err
}

func (c *Config) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.host, c.port, c.user, c.password, c.dbName)

	var err error
	c.Engine, err = xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		helpers.Log.Error(fmt.Sprintf("Failed to connect to the db. Err: %v", err))
	}

	return err
}

func (c *Config) PingDB() error {
	err := c.Engine.Ping()
	if err != nil {
		helpers.Log.Error(fmt.Sprintf("Ping to DB Failed. Err: %v", err))
	}

	return err
}

func (c *Config) Close() {
	c.Engine.Close()
}

func (c *Config) Migrate() error {
	m := migrate.New(c.Engine, &migrate.Options{
		TableName:    "migrations",
		IDColumnName: "id",
	}, migrations)

	err := m.Migrate()
	if err != nil {
		helpers.Log.Error(fmt.Sprintf("Error running migrations. Err: %v", err))
	}

	return err
}
