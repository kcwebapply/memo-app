package dbConfig

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

var conn *dbr.Connection

// db initialization
func init() {
	var config = getConfig()
	connection, err := dbr.Open("postgres", "postgres://"+config.DB.User+":"+config.DB.Password+"@"+config.DB.Host+"/"+config.DB.DbName+"?sslmode=disable", nil)
	if err != nil {
		fmt.Println("error happened in connection:", err)
	}
	conn = connection
}

// get DatabaseConnection
func GetConnection() *dbr.Connection {
	return conn
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
