package main

import (
	"log/slog"
	"os"
	"stupid-news/backend/internal/config"
	"stupid-news/backend/internal/database"
	// "stupid-news/backend/internal/api/routes"
	// "stupid-news/backend/pkg/logger"
)

func setupLogger() *slog.Logger {
	var handler slog.Handler

	if os.Getenv("ENV") == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}
	return slog.New(handler)
}

func main() {
	// Initialize logger
	logger := setupLogger()

	logger.Info("Application starting", "name", "stupid-news-backend")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	logger.Info("Configuration loaded",
		"server_port", cfg.Server.Port,
		"server_host", cfg.Server.Host,
		"db_host", cfg.Database.Host,
		"db_port", cfg.Database.Port,
		"db_max_open_conns", cfg.Database.MaxOpenConns,
		"db_max_idle_conns", cfg.Database.MaxIdleConns,
	)

	// Initialize database
	db, err := database.NewConnection(&cfg.Database)
	if err != nil {
		logger.Error("Failed to connect to database",
			"error", err,
			"db_host", cfg.Database.Host,
			"db_port", cfg.Database.Port,
		)
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("Database connection established")

	// Start server
	logger.Info("Server starting",
		"port", cfg.Server.Port,
		"host", cfg.Server.Host,
		"read_timeout_ms", cfg.Server.ReadTimeout.Milliseconds(),
		"write_timeout_ms", cfg.Server.WriteTimeout.Milliseconds(),
	)
	
	// Note: You'll need to uncomment and implement router
	// log.Fatal(router.ListenAndServe(":" + strconv.Itoa(cfg.Server.Port)))
}
