package config

import (
	"starter-go-gin/libs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormMysqlConnection() (*gorm.DB, error) {
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
	conn, err := gorm.Open(mysql.Open(consStr), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		return nil, err
	}

	return conn, nil
}
