package dbConfig

import (
	"github.com/BurntSushi/toml"
)

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

func GetConfig() Config {
	var config Config
	toml.DecodeFile("config.toml", &config)
	return config
}
