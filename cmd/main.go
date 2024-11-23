package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gw-exchanger/internal/app"
	"gw-exchanger/internal/config"
	"gw-exchanger/internal/services"
	"gw-exchanger/internal/storages/postgres"
	logging "gw-exchanger/pkg/logs"
)

func init() {

}

func main() {
	cfg := config.New()
	log := logging.SetupLogger(cfg.Env)

	db, err := sql.Open("postgres", cfg.DatabaseURL())
	if err != nil {
		log.Error("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Error("Не удалось подключиться к базе данных: %v", err)
	}

	storage := postgres.NewPostgresStorage(db)
	exchangeService := services.NewExchangeService(storage)

	application := app.New(log, cfg.GRPCPort, exchangeService)
	application.MustRun()
}
