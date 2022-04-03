package config

import (
	"github.com/joho/godotenv"
	"os"
)

const envName = "deploy/.env"

type Config struct{}

func NewConfig() *Config {
	if err := godotenv.Load(envName); err != nil {
		panic(err)
	}

	return &Config{}
}

func (c *Config) GetSelfPort() string {
	return os.Getenv("SELF_PORT")
}

func (c *Config) GetHost() string {
	return os.Getenv("MONGO_HOST")
}

func (c *Config) GetPort() string {
	return os.Getenv("MONGO_PORT")
}

func (c *Config) GetDbName() string {
	return os.Getenv("MONGO_DATABASE")
}

func (c *Config) GetUser() string {
	return os.Getenv("MONGO_USER")
}

func (c *Config) GetPassword() string {
	return os.Getenv("MONGO_PWD")
}