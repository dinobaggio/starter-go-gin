package config

import (
	"database/sql"
	"flag"
	"log"
	"starter-go-gin/libs"
)

var pool *sql.DB

func init() {
	type DBConfig struct {
		hostname string
		port     string
		user     string
		password string
		database string
	}
	var dbConfig DBConfig
	dbConfig.hostname = libs.EnvVariable("DB_HOST")
	dbConfig.port = libs.EnvVariable("DB_PORT")
	dbConfig.user = libs.EnvVariable("DB_USERNAME")
	dbConfig.password = libs.EnvVariable("DB_PASSWORD")
	dbConfig.database = libs.EnvVariable("DB_DATABASE")

	consStr := dbConfig.user + ":" + dbConfig.password + "@tcp(" + dbConfig.hostname + ":" + dbConfig.port + ")/" + dbConfig.database + "?parseTime=true"
	dsn := flag.String("dsn", consStr, "connection data source name")
	flag.Parse()

	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	var err error

	// Opening a driver typically will not attempt to connect to the database.
	pool, err = sql.Open(libs.EnvVariable("DB_CONNECTION"), *dsn)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}
}

func SQLDBConn() *sql.DB {
	return pool
}
