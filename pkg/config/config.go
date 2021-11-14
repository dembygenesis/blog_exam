package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// Config hold app config variables
type Config struct {
	Database Database `json:"database"`
	AppPort  int
}

// Database holds database variables
type Database struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"username"`
	Pass   string `json:"password"`
	Schema string `json:"schema"`
	Driver string `json:"driver"`
}

/*func (c *Config) ReadConfig(f string) (*Config, error) {

}*/

// ReadConfig reads the config file
func ReadConfig(f string) (*Config, error) {
	err := godotenv.Load(f)
	if err != nil {
		return nil, fmt.Errorf("error loading %v file", f)
	}
	return &Config{}, nil
}

func (c *Config) SetConfig() error {
	// Set database config
	envDBPort := os.Getenv("DB_PORT")
	envDBHost := os.Getenv("DB_HOST")
	envDBUser := os.Getenv("DB_USERNAME")
	envDBPass := os.Getenv("DB_PASSWORD")
	envDBSchema := os.Getenv("DB_SCHEMA")
	envDBDriver := os.Getenv("DB_DRIVER")

	if envDBHost == "" {
		return fmt.Errorf("env: DB_HOST is missing")
	}
	dbPort, err := strconv.Atoi(envDBPort)
	if err != nil {
		return fmt.Errorf("error trying to parse %v (DB_PORT) into an int", envDBPort)
	}
	if envDBHost == "" {
		return fmt.Errorf("env: DB_HOST is missing")
	}
	if envDBUser == "" {
		return fmt.Errorf("env: DB_USERNAME is missing")
	}
	if envDBPass == "" {
		return fmt.Errorf("env: DB_PASSWORD is missing")
	}
	if envDBSchema == "" {
		return fmt.Errorf("env: DB_SCHEMA is missing")
	}
	if envDBDriver == "" {
		return fmt.Errorf("env: DB_DRIVER is missing")
	}

	c.Database = Database{
		Host:   envDBHost,
		Port:   dbPort,
		User:   envDBUser,
		Pass:   envDBPass,
		Schema: envDBSchema,
		Driver: envDBDriver,
	}

	// Set app port
	port := os.Getenv("PORT")
	if port == "" {
		return fmt.Errorf("env: PORT is missing")
	}
	appPort, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("error trying to parse %v (DB_PORT) into an int", port)
	}
	c.AppPort = appPort

	return nil
}
