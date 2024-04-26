package config

import (
	"fmt"
	"os"
	"sync"
)

var singleton Config
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		singleton = new()
	})

	return singleton
}

type Config struct {
	DbConfig
	FsConfig
	ServerConfig
}

type DbConfig struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbName     string `env:"DB_NAME"`
	DbSSLMode  string `env:"DB_SSL"`
}

type FsConfig struct {
	FsEndPoint string `env:"FS_END_POINT"`
	FsUser     string `env:"FS_USER"`
	FsPassword string `env:"FS_PASSWORD"`
}

type ServerConfig struct {
	ServerPort string `env:"SERVER_PORT"`
	ServerHost string `env:"SERVER_HOST"`
}

func (c *Config) GetDbConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DbHost, c.DbPort, c.DbUser, c.DbPassword, c.DbName, c.DbSSLMode)
}

func new() Config {
	dbUser, isSet := os.LookupEnv("DB_USER")
	if !isSet {
		panic("DB_USER isnt set")
	}

	dbPassword, isSet := os.LookupEnv("DB_PASSWORD")
	if !isSet {
		panic("DB_PASSWORD isnt set")
	}

	dbHost, isSet := os.LookupEnv("DB_HOST")
	if !isSet {
		panic("DB_HOST isnt set")
	}

	dbPort, isSet := os.LookupEnv("DB_PORT")
	if !isSet {
		panic("DB_PORT isnt set")
	}

	dbName, isSet := os.LookupEnv("DB_NAME")
	if !isSet {
		panic("DB_NAME isnt set")
	}

	dbSSLMode, isSet := os.LookupEnv("DB_SSL")
	if !isSet {
		panic("DB_SSL isnt set")
	}

	fsUser, isSet := os.LookupEnv("FS_USER")
	if !isSet {
		panic("FS_USER isnt set")
	}

	fsPassword, isSet := os.LookupEnv("FS_PASSWORD")
	if !isSet {
		panic("FS_PASSWORD isnt set")
	}

	fsEndPoint, isSet := os.LookupEnv("FS_END_POINT")
	if !isSet {
		panic("FS_END_POINT isnt set")
	}

	serverPort, isSet := os.LookupEnv("SERVER_PORT")
	if !isSet {
		panic("SERVER_PORT isnt set")
	}

	serverHost, isSet := os.LookupEnv("SERVER_HOST")
	if !isSet {
		panic("SERVER_HOST isnt set")
	}

	return Config{
		DbConfig: DbConfig{
			DbUser:     dbUser,
			DbPassword: dbPassword,
			DbHost:     dbHost,
			DbPort:     dbPort,
			DbName:     dbName,
			DbSSLMode:  dbSSLMode,
		},
		FsConfig: FsConfig{
			FsUser:     fsUser,
			FsPassword: fsPassword,
			FsEndPoint: fsEndPoint,
		},
		ServerConfig: ServerConfig{
			ServerPort: serverPort,
			ServerHost: serverHost,
		},
	}
}
