package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/myshkins/gopetwatch/logger"
)

var Database *sql.DB

func Connect() {
	var err error
	Database, err = sql.Open("mysql", "chancho:raisin@(gopetwatch_mysql)/gopetwatch?parseTime=true")
	if err != nil {
		logger.Log.Fatal("db connection failed. ", err)
	} else {
		logger.Log.Info("db connection succeeded.")
	}

	if err := Database.Ping(); err != nil {
		logger.Log.Fatal("db not responding ", err)
	} else {
		logger.Log.Info("db is responding")
	}

	logger.Log.Info("about to drop table if exists")
	_, err = Database.Exec(`drop table if exists readings`)
	if err != nil {
		logger.Log.Warn(err)
	} else {
		logger.Log.Info("drop table suceeded")
	}
}

func CreateTable() {
	query := `
		create table readings (
			id int auto_increment,
			temperature float not null,
			reading_timestamp datetime,
			primary key (id)
		)`

	logger.Log.Info("about to create table")
	_, err := Database.Exec(query)
	if err != nil {
		logger.Log.Fatal("table creation failed", err)
	} else {
		logger.Log.Info("table creation succeeded")
	}
}

func SeedDatabase() {
	query := `
		insert into readings (temperature, reading_timestamp)
		values(?, ?)`
	t1 := time.Date(2023, time.June, 2, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, time.June, 2, 0, 1, 0, 0, time.UTC)
	t3 := time.Date(2023, time.June, 2, 0, 2, 0, 0, time.UTC)
	_, err := Database.Exec(query, 65.5, t1)
	_, err = Database.Exec(query, 66.1, t2)
	_, err = Database.Exec(query, 65.3, t3)
	if err != nil {
		logger.Log.Warn("database seeding failed, err: ", err)
	} else {
		logger.Log.Info("database seeding suceeded.")
	}
}

func Query24hr() []Reading {
	query := `
		select temperature, reading_timestamp
		from readings
	  where reading_timestamp >= now() - interval 1 day`
	rows, err := Database.Query(query)
	if err != nil {
		logger.Log.Warn(err)
	}
	defer rows.Close()
	readings := make([]Reading, 0)

	for rows.Next() {
		var r Reading
		if err := rows.Scan(&r.Temperature, &r.ReadingTimestamp); err != nil {
			logger.Log.Warn(err)
		}
		readings = append(readings, r)
	}
	// check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		logger.Log.Warn(err)
	}
	logger.Log.Infof("readings: %v", readings)

	return readings
}

func InsertRow(reading Reading) (int64, error) {
	logger.Log.Info(reading.ReadingTimestamp)
  result, err := Database.Exec(
    `INSERT into readings (temperature, reading_timestamp)
      VALUES (?, ?)`, reading.Temperature, reading.ReadingTimestamp,
	)

	if err != nil {
		return 0, fmt.Errorf("InsertRow() error: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("InsertRow() error: %v", err)
	}
	return id, nil
}
