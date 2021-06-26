package models

import (
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/xorm/migrate"
)

var migrations = []*migrate.Migration{
	{
		ID: "201608301400",
		Migrate: func(tx *xorm.Engine) error {
			return tx.Sync2(&User{})
		},
		Rollback: func(tx *xorm.Engine) error {
			return tx.DropTables(&User{})
		},
	},
}
