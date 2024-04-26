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
		singleton = config()
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
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName, c.DbSSLMode)
}

func config() Config {
	return Config{
		DbConfig: DbConfig{
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbName:     os.Getenv("DB_NAME"),
			DbSSLMode:  os.Getenv("DB_SSL"),
		},
		FsConfig: FsConfig{
			FsUser:     os.Getenv("FS_USER"),
			FsPassword: os.Getenv("FS_PASSWORD"),
			FsEndPoint: os.Getenv("FS_END_POINT"),
		},
		ServerConfig: ServerConfig{
			ServerPort: os.Getenv("SERVER_PORT"),
			ServerHost: os.Getenv("SERVER_HOST"),
		},
	}
}
