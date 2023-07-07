package database

import (
	"database/sql"

	"time"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func Connect() {
	Database, err := sql.Open("mysql", "chancho:raisin@(gopetwatch_mysql)/gopetwatch?parseTime=true")
	if err != nil {
		log.Fatal("db connection failed. ", err)
	}

	if err := Database.Ping(); err != nil {
		log.Fatal("db not responding ", err)
	}
}

func createTable() (error) {
	Database.Exec(`drop table if exists readings`)
	query := `
		create table readings (
			id int auto_increment,
			temperature float not null,
			reading_timestamp datetime,
			primary key (id)
		)`

	_, err := Database.Exec(query)
	if err != nil {
		log.Fatal("table creation failed", err)
	}
	return err
}

func seedDatabase() (error) {
	query := `
		insert into readings (temperature, reading_timestamp)
		values(?, ?)`
	t1 := time.Date(2023, time.June, 2, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, time.June, 2, 0, 1, 0, 0, time.UTC)
	t3 := time.Date(2023, time.June, 2, 0, 2, 0, 0, time.UTC)
	_, err := Database.Exec(query, 70.5, t1)
	_, err = Database.Exec(query, 70.1, t2)
	_, err = Database.Exec(query, 71.3, t3)
	return err
}

func QueryReadings() ([]Reading) {
	rows, err := Database.Query("select temperature, reading_timestamp from readings")
	if err != nil {
		log.Warn(err)
	}
	defer rows.Close()
	readings := make([]Reading, 0)

	for rows.Next() {
		var r Reading
		if err := rows.Scan(&r.Temperature, &r.Reading_timestamp); err != nil {
			log.Warn(err)
		}
		readings = append(readings, r)
	}
	// check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Warn(err)
	}

	return readings
}

