package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port                  string
	PostDBHost            string
	PostDBPort            string
	ConnectionServiceHost string
	ConnectionServicePort string
	ProfileServiceHost    string
	ProfileServicePort    string
	Neo4jHost             string
	Neo4jPort             string
	Neo4jProtocol         string
	Neo4jUsername         string
	Neo4jPassword         string
}

func NewConfig() *Config {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		fmt.Println("docker")

		return &Config{
			Port:                  os.Getenv("POST_SERVICE_PORT"),
			PostDBHost:            os.Getenv("POST_DB_HOST"),
			PostDBPort:            os.Getenv("POST_DB_PORT"),
			ConnectionServiceHost: os.Getenv("CONNECTION_SERVICE_HOST"),
			ConnectionServicePort: os.Getenv("CONNECTION_SERVICE_PORT"),
			ProfileServiceHost:    os.Getenv("PROFILE_SERVICE_HOST"),
			ProfileServicePort:    os.Getenv("PROFILE_SERVICE_PORT"),
			Neo4jPort:             os.Getenv("POST_GRAPH_DB_PORT"),
			Neo4jHost:             os.Getenv("POST_GRAPH_DB_HOST"),
			Neo4jProtocol:         os.Getenv("POST_GRAPH_DB_PROTOCOL"),
			Neo4jUsername:         os.Getenv("POST_GRAPH_DB_USERNAME"),
			Neo4jPassword:         os.Getenv("POST_GRAPH_DB_PASSWORD"),
		}
	} else {
		fmt.Println("local")

		return &Config{
			Port:                  "8002",
			PostDBHost:            "localhost",
			PostDBPort:            "27017",
			ConnectionServiceHost: "localhost",
			ConnectionServicePort: "8004",
			ProfileServiceHost:    "localhost",
			ProfileServicePort:    "8001",
			Neo4jPort:             "7687",
			Neo4jHost:             "localhost",
			Neo4jProtocol:         "bolt",
			Neo4jUsername:         "neo4j",
			Neo4jPassword:         "password",
		}
	}
}
