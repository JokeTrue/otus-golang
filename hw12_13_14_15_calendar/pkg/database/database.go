package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq" // Init Database Driver
)

func GetDatabase(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.Fatalln(err)
	}
	return db
}
