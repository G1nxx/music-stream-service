package main

import (
	"log/slog"
	"music-stream-service/controllers"
	"music-stream-service/internal/config"
	"music-stream-service/internal/lib/logger/sl"
	_ "music-stream-service/internal/repositories"
	"music-stream-service/internal/repositories/postgresql"
	"music-stream-service/service"
	"net/http"
	"os"

	"github.com/rs/cors"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	os.Setenv("CONFIG_PATH", "config/local.yaml")

	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	// TODO: init logger: slog
	log := setupLogger(cfg.Env)

	log.Info("starting music-stream-service", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	db, err := postgresql.New(cfg.DBServer)
	if err != nil {
		log.Error("failed to init database", sl.Err(err))
		os.Exit(1)
	} else {
		log.Info("database successfully initiated", slog.String("env", cfg.Env))
	}

	// TODO: init router: chi, "chi render"

	// TODO: run server
	//!!!
	repos := service.NewRepository(db, cfg, log)
	serv, err := service.NewService(*repos, log)
	if err != nil {
		log.Error("failed to init service", sl.Err(err))
	}

	controllers := controllers.NewController(serv)

	mux := http.NewServeMux()

	frontFS := http.FileServer(http.Dir("front"))
	mux.Handle("/", frontFS)

	mediaFS := http.FileServer(http.Dir("files"))
	mux.Handle("/files/", http.StripPrefix("/files/", mediaFS))

	mux.HandleFunc("/music", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "front/music.html")
	})

	mux.HandleFunc("/subs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "front/subs.html")
	})

	// API routes
	controllers.RegisterRoutes(mux)

	muxHandler := cors.New(
		cors.Options{
			AllowedOrigins: []string{"http://localhost:3243"},
			AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
			AllowedMethods: []string{"GET", "POST", "PUT", "UPDATE", "DELETE", "OPTIONS"},
		},
	).Handler(mux)

	addr := "localhost:8080"
	log.Info("Starting server on " + addr)
	if err := http.ListenAndServe(addr, muxHandler); err != nil {
		log.Error("failed to start server", sl.Err(err))
		os.Exit(1)
	}

	// fs = http.FileServer(http.Dir("front"))
	// http.Handle("/", fs)
	// sl.Err(http.ListenAndServe(":8080", nil))
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
