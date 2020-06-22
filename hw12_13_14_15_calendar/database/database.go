package database

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq" // Init Database Driver
)

func GetDatabase(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.Fatalln(err)
	}

	migrate := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Fatalln(err)
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		db.MustExec(string(content))
		return nil
	}

	// Auto Migrate
	err = filepath.Walk("./migrations", migrate)
	if err != nil {
		logrus.Fatalln(err)
	}
	return db
}
