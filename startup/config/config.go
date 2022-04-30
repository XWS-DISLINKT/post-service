package config

type Config struct {
	Port       string
	PostDBHost string
	PostDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       "8002",      //os.Getenv("POST_SERVICE_PORT"),
		PostDBHost: "localhost", //os.Getenv("POST_DB_HOST"),
		PostDBPort: "27017",     //os.Getenv("POST_DB_PORT"),
	}
}
