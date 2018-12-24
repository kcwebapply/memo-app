package dbConfig

import (
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

var db *sql.DB

// db initialization
func init() {
	var config = getConfig()
	var err error
	db, err = sql.Open("postgres", "postgres://"+config.DB.User+":"+config.DB.Password+"@"+config.DB.Host+"/"+config.DB.DbName+"?sslmode=disable")
	if err != nil {
		fmt.Println("error connection:", err)
		panic(err)
	}
}

// get DatabaseConnection
func GetConnection() *sql.DB {
	return db
}

// get config prom toml property
func getConfig() Config {
	var config Config
	toml.DecodeFile("config.toml", &config)
	return config
}

type Config struct {
	DB DBConfig `toml:DB`
}

type DBConfig struct {
	DbName   string `toml:dbName`
	Port     uint   `toml:port`
	User     string `toml:user`
	Password string `toml:password`
	Host     string `toml:host`
}
