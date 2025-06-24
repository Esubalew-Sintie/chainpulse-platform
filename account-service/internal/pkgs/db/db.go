package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chainpulse/backend/account/internal/models"
	"github.com/chainpulse/backend/account/internal/pkgs/config"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDSN generates a PostgreSQL DSN from the GlobalConfig.
func GetDSN(cfg *config.GlobalConfig) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DATABASE.USER,
		cfg.DATABASE.PASSWORD,
		cfg.DATABASE.NAME,
		cfg.DATABASE.HOST,
		cfg.DATABASE.PORT,
	)
}

// OpenDB initializes the database connection, runs migrations, and returns a cleanup function.
func OpenDB(dsn string) (*gorm.DB, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, func() {}, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, func() {}, fmt.Errorf("failed to get SQL DB instance: %w", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, func() {}, fmt.Errorf("failed to ping database: %w", err)
	}

	// Auto-migrate models
	if err := migrate(db); err != nil {
		log.Fatal("Database migration failed:", err)
	}

	// Return cleanup function
	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			log.Println("Error closing database:", err)
		}
	}
	return db, cleanup, nil
}

// migrate handles schema migrations for your models.
func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Buyer{},
		&models.Settings{},
	)
}
