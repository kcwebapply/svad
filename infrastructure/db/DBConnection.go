package db

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

var conn *dbr.Connection

//var filePath = "config.toml"

// db initialization
func InitDb(filePath string) {

	var config Config = getConfig(filePath)

	var user = config.DB.User
	var password = config.DB.Password
	var host = config.DB.Host
	var dbname = config.DB.DbName
	var driver = config.DB.Driver

	configIsEnough := true
	if user == "" {
		fmt.Println("please configurate \"user\" param on config file.")
		configIsEnough = false
	}

	if password == "" {
		fmt.Println("please configurate \"password\" param on config file.")
		configIsEnough = false
	}

	if host == "" {
		fmt.Println("please configurate \"host\" param on config file.")
		configIsEnough = false
	}

	if dbname == "" {
		fmt.Println("please configurate \"dbname\" param on config file.")
		configIsEnough = false
	}

	if driver == "" {
		fmt.Println("please configurate \"driver\" param on config file.")
		configIsEnough = false
	}

	if !configIsEnough {
		os.Exit(0)
	}

	connection, err := dbr.Open(driver, "user="+user+" password="+password+" dbname="+dbname+" host="+host+" sslmode=disable", nil)
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
