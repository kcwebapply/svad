package db

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

var conn *dbr.Connection
var filePath = "config.toml"

// db initialization
func init() {

	var config Config = getConfig(filePath)

	var user = config.DB.User
	var password = config.DB.Password
	var host = config.DB.Host
	var dbname = config.DB.DbName
	var driver = config.DB.Driver

	if user == "" {

	}

	if password == "" {

	}

	if host == "" {

	}

	if dbname == "" {

	}

	if driver == "" {

	}

	connection, err := dbr.Open(driver, "postgres://"+user+":"+password+"@"+host+"/"+dbname+"?sslmode=disable", nil)
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
func getConfig(filePath string) Config {
	var config Config
	toml.DecodeFile(filePath, &config)
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
	Driver   string `toml:driver`
}
