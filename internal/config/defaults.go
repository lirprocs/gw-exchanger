package config

type DefaultConfig struct {
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

var Defaults = DefaultConfig{
	Env:              "prod",
	PostgresUser:     "postgres",
	PostgresPassword: "password",
	PostgresDB:       "exchange_db",
	DBHost:           "localhost",
	DBPort:           "5432",
	GRPCPort:         "50051",
	Timeout:          "5",
}
