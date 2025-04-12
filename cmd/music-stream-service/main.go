package main

import (
	"fmt"
	"log/slog"
	_ "music-stream-service/domain/entities"
	"music-stream-service/internal/config"
	"music-stream-service/internal/lib/logger/sl"
	_ "music-stream-service/internal/storage"
	"music-stream-service/internal/storage/postgresql"
	"os"
)


const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "G1nxx"
	password = "password1"
	dbname   = "music-stream-service"
)


func main() {
	os.Setenv("CONFIG_PATH", "config/local.yaml")

	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	// TODO: init logger: slog
	log := setupLogger(cfg.Env)

	log.Info("starting music-stream-service", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// TODO: init storage: postgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

	storage, err := postgresql.New(psqlInfo)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage

	//var user e.User = e.User{Id: 2, Login: "G2", Email: "g@2", Paswdhash: "32345678"}

	// err = storage.CreateUser(user)
	// if err != nil {
	// 	log.Error("failed to create user", sl.Err(err))
	// 	os.Exit(1)
	// }

	// TODO: init router: chi, "chi render"

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}