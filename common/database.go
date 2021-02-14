package common

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// CreateDBConnection returns DB connection
func CreateDBConnection(c *Conf) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", c.DB.Host, c.DB.Port),
		User:     c.DB.User,
		Password: c.DB.Pass,
		Database: c.DB.Name,
	})
}

// CreateTable creates a DB table from a model
func CreateTable(db *pg.DB, conf *Conf, model interface{}) error {
	err := db.Model(model).CreateTable(&orm.CreateTableOptions{
		Temp:          conf.DB.Test,
		IfNotExists:   true,
		FKConstraints: true,
	})
	return err
}
