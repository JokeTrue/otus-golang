package database

import (
	"fmt"

	"github.com/JokeTrue/otus-golang/hw12_13_14_15_calendar/internal/config/calendar"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq" // Init Database Driver
)

var (
	schema = `
	CREATE TABLE IF NOT EXISTS events
	(
		id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_id         NUMERIC   NOT NULL,
		title           TEXT      NOT NULL,
		description     TEXT,
		start_date      TIMESTAMP NOT NULL,
		end_date        TIMESTAMP NOT NULL,
		notify_interval NUMERIC CHECK (notify_interval > 0)
	);
	`
)

func GetDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		calendar.Conf.Database.Host,
		calendar.Conf.Database.Port,
		calendar.Conf.Database.User,
		calendar.Conf.Database.Password,
		calendar.Conf.Database.Name,
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.Fatalln(err)
	}

	// Auto Migrate
	db.MustExec(schema)
	return db
}
