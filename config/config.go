package config

type Config struct {
	Port     int `json:"port" env:"SERVICE_PORT"`
	Database struct {
		User     string `json:"user" env:"DB_USER"`
		Password string `json:"password" env:"DB_PASSWORD"`
		Schema   string `json:"schema" env:"DB_SCHEMA"`
		Address  string `json:"address" env:"DB_ADDRESS"`
	} `json:"database"`
}
