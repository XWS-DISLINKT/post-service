package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	PostDBHost string
	PostDBPort string
}

func NewConfig() *Config {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		fmt.Println("docker")

		return &Config{
			Port:       os.Getenv("POST_SERVICE_PORT"),
			PostDBHost: os.Getenv("POST_DB_HOST"),
			PostDBPort: os.Getenv("POST_DB_PORT"),
		}
	} else {
		fmt.Println("local")

		return &Config{
			Port:       "8002",
			PostDBHost: "localhost",
			PostDBPort: "27017",
		}
	}
}
