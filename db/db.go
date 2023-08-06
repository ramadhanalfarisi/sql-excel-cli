package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func ConnectDB() *sql.DB {
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	host := viper.Get("SQL_EXCEL_CLI_HOST").(string)
	uname := viper.Get("SQL_EXCEL_CLI_USERNAME").(string)
	password := viper.Get("SQL_EXCEL_CLI_PASSWORD").(string)
	port := viper.Get("SQL_EXCEL_CLI_PORT").(string)
	dbname := viper.Get("SQL_EXCEL_CLI_NAME").(string)

	if len(host) <= 0 ||
		len(uname) <= 0 ||
		len(password) <= 0 ||
		len(port) <= 0 ||
		len(dbname) <= 0 {
		log.Fatal(fmt.Errorf("environment has not been inserted. please set environment in 'sql-excel-cli env'"))
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", uname, password, host, port, dbname)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	return db
}
