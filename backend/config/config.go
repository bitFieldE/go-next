package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port              string `toml:"port"`
	ClientUrl         string `toml:"client_url"`
	AccessLogfile     string `toml:"accessLogfile"`
	AppLogfile        string `toml:"appLogfile"`
	UploadMaxFileSize int64  `toml:"uploadMaxFileSize"`
}

type ConfigList struct {
	Server ServerConfig `toml:"server"`
	Db     DbConfig     `toml:"database"`
}

type DbConfig struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Name     string `toml:"db"`
}

var Config ConfigList
var Db DbConfig
var Server ServerConfig

func LoadDotEnv() {
	err := godotenv.Load(fmt.Sprintf("config/env/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}
}

func init() {
	LoadDotEnv()
	confPath := fmt.Sprintf("config/env/%s.toml", os.Getenv("GO_ENV"))
	_, err := toml.DecodeFile(confPath, &Config)
	Db.Password = os.Getenv("DATABASE_PASSWORD")
	Db.User = os.Getenv("DATABASE_USER")
	Db.Port = os.Getenv("DATABASE_PORT")
	Db.Name = os.Getenv("DATABASE_NAME")
	pass := os.Getenv("DATABASE_PASSWORD")
	fmt.Println("pass: ", pass)
	if err != nil {
		panic(err)
	}
	Db = Config.Db
	Server = Config.Server
}
