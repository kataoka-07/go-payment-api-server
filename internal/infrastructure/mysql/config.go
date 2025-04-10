package mysql

import "os"

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewConfigFromEnv() Config {
	return Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
