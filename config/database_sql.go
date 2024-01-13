package config

import (
	"database/sql"
	"flag"
	"log"
	"starter-go-gin/libs"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB

func init() {
	// Note: Prevent errors like "flag provided but not defined: -test.paniconexit0" from occurring in go test.
	testing.Init()

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
	pool, err = sql.Open(libs.EnvVariable("DB_CONNECTION"), *dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
}

func SQLDBConn() *sql.DB {
	return pool
}
