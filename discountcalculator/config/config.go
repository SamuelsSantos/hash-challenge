package config

import (
	"fmt"
	"time"
)

type serviceConfig struct {
	Host string
	Port string
}

// Config struct
type Config struct {
	UserConfig      *serviceConfig
	ProductConfig   *serviceConfig
	BlackFridayDate time.Time // Format Date: YYYY-MM-DD -> "2010-12-06"
}

//NewConfig new struct to config environments
func NewConfig() *Config {
	return &Config{
		UserConfig: &serviceConfig{
			Host: GetenvString("USER_SERVICE_HOST", "127.0.0.1"),
			Port: GetenvString("USER_SERVICE_PORT", "8485"),
		},
		ProductConfig: &serviceConfig{
			Host: GetenvString("PRODUCT_SERVICE_HOST", "127.0.0.1"),
			Port: GetenvString("PRODUCT_SERVICE_PORT", "8486"),
		},
		BlackFridayDate: GetenvDate("BLACK_FRIDAY", "2020-11-25"),
	}
}

// ToString ...
func (cfg *Config) ToString() string {
	description := "User Service >> %s \n Product Service >> %s \n"
	return fmt.Sprintf(description,
		cfg.UserConfig.ToURL(),
		cfg.ProductConfig.ToURL(),
	)
}

// ToURL return a string address [HOST:PORT]
func (c *serviceConfig) ToURL() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// ToStringISO - LayoutISO: YYYY-MM-DD
func (cfg *Config) ToStringISO() string {
	return cfg.BlackFridayDate.Format(LayoutISO)
}
