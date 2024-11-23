package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load("config.env"); err != nil {
		log.Print("No config.env file found")
	}
}

type Config struct {
	Env              string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	DBHost           string
	DBPort           string
	MigrationsPath   string
	GRPCPort         string
	Timeout          string
}

func New() *Config {
	return &Config{
		Env:              getEnv("ENV", Defaults.Env),
		PostgresUser:     getEnv("POSTGRES_USER", Defaults.PostgresUser),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", Defaults.PostgresPassword),
		PostgresDB:       getEnv("POSTGRES_DB", Defaults.PostgresDB),
		DBHost:           getEnv("DB_HOST", Defaults.DBHost),
		DBPort:           getEnv("DB_PORT", Defaults.DBPort),
		MigrationsPath:   getEnv("MIGRATIONS_PATH", Defaults.MigrationsPath),
		GRPCPort:         getEnv("GRPC_PORT", Defaults.GRPCPort),
		Timeout:          getEnv("TIMEOUT", Defaults.Timeout),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.PostgresUser, c.PostgresPassword, c.DBHost, c.DBPort, c.PostgresDB)
}
