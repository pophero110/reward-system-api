package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"reward-system-api/internal/api"
	"reward-system-api/internal/model"
	"reward-system-api/internal/repository"
	"reward-system-api/internal/service"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logger.Warn("⚠️ No .env file found — relying on system environment variables")
	}

	// Connect to PostgreSQL using GORM
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		logger.Error("❌ DATABASE_URL is not set")
		os.Exit(1)
	}

	db, err := connectGormDB(dsn, logger)
	if err != nil {
		logger.Error("❌ Database connection failed", "error", err)
		os.Exit(1)
	}

	// ✅ Auto-migrate your models
	err = db.AutoMigrate(&model.Quest{}, &model.Board{}, &model.User{})
	if err != nil {
		logger.Error("❌ Migration failed", "error", err)
		os.Exit(1)
	}

	// Build the application
	questService := &service.QuestService{Logger: logger, Quests: &repository.QuestModel{DB: db}}
	userService := &service.UserService{Logger: logger, Users: &repository.UserModel{DB: db}}
	boardService := &service.BoardService{Logger: logger, Boards: &repository.BoardModel{DB: db}}
	app := &api.Application{
		Logger:       logger,
		QuestService: questService,
		UserService:  userService,
		BoardService: boardService,
	}

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(app.Routes())

	// Start http server
	port := parsePort(os.Getenv("PORT"), 4000)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		IdleTimeout:  45 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}
	logger.Info("✅ Server started", "addr", fmt.Sprintf(":%d", port))
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("❌ Server error", "error", err)
	}
}

// connectGormDB establishes a GORM connection and verifies it
func connectGormDB(dsn string, logger *slog.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic DB handle: %w", err)
	}

	// Optional connection pool tuning
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("✅ Database connection established")
	return db, nil
}

func parsePort(portStr string, defaultPort int) int {
	if portStr == "" {
		return defaultPort
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return defaultPort
	}
	return port
}
