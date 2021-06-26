package models

import "time"

type User struct {
	Id       int64
	Name     string    `json:"name"`
	Username string    `xorm:"varchar(12) not null unique" json:"username"`
	Password string    `xorm:"varchar(32)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}
