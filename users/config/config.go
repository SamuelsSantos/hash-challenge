package config

import "fmt"

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
	Driver   string
}

type serverConfig struct {
	Port string
}

// Config struct
type Config struct {
	Db *dbConfig
}

//NewConfig new struct to configurations enviroments
func NewConfig() *Config {
	return &Config{
		Db: &dbConfig{
			Host:     GetenvString("DB_HOST", "127.0.0.1"),
			Port:     GetenvString("DB_PORT", "5433"),
			User:     GetenvString("DB_USER", "postgres"),
			Name:     GetenvString("DB_NAME", "users"),
			Password: GetenvString("DB_PASSWORD", "hash"),
			Driver:   GetenvString("DB_DRIVER", "postgres"),
		},
	}
}

func (cfg *dbConfig) ToURL() string {
	url := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable "
	return fmt.Sprintf(url, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
}
