package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/pingcap/log"
)

type Config struct {
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort int
}

func NewConfig() *Config {
	return &Config{
		DBName: "demo",
		DBUser: "demo",
		DBPass: "123456",
		DBHost: "localhost",
		DBPort: 3306,
	}
}

func NewDBC(conf *Config) (*sql.DB, error) {
	dbCfg := mysql.Config{
		User:                 conf.DBUser,
		Passwd:               conf.DBUser,
		Addr:                 conf.DBHost,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dbc, err := sql.Open("mysql", dbCfg.FormatDSN())
	if err != nil {
		fmt.Printf("Database connect error")
	}
	return dbc, nil
}

func main() {
	dbc, err := NewDBC(NewConfig())
	if err != nil {
		// log.Fatal("Database connect faile")
		fmt.Sprintln(err)
	}

	pingErr := dbc.Ping()
	if pingErr != nil {
		fmt.Sprintln(err)
	}

	log.Info("conected")

	// Example usage of Config
}
