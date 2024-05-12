package config

import (
	"fmt"
	"github.com/caarlos0/env"
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
	DbUser       string `env:"DB_USER" envDefault:"user"`
	DbPassword   string `env:"DB_PASSWORD" envDefault:"<password>"`
	DbHost       string `env:"DB_HOST" envDefault:"localhost"`
	DbPort       string `env:"DB_PORT" envDefault:"8080"`
	DbName       string `env:"DB_NAME" envDefault:"student_records"`
	DbSSLMode    string `env:"DB_SSL" envDefault:"disable"`
	FsEndPoint   string `env:"FS_END_POINT" envDefault:"9000"`
	FsUser       string `env:"FS_USER" envDefault:"user"`
	FsPassword   string `env:"FS_PASSWORD" envDefault:"password"`
	ServerPort   string `env:"SERVER_PORT" envDefault:"8080"`
	ServerHost   string `env:"SERVER_HOST" envDefault:"localhost"`
	JwtSecretKey string `env:"JWT_SECRET_KEY" envDefault:"mysecretkey"`
}

type DbConfig struct {
}

type FsConfig struct {
}

type ServerConfig struct {
}

func (c Config) GetDbConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName, c.DbSSLMode)
}

func (c Config) GetJwtSecretKey() []byte {
	return []byte(c.JwtSecretKey)
}

func config() Config {
	var result Config
	err := env.Parse(&result)
	if err != nil {
		panic("Config cant download")
	}

	return result
}
