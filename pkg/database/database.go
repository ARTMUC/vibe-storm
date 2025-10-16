package database

import (
	"fmt"
	"log"
	"vibe-storm/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Initialize initializes the database connection
func Initialize(dbConfig config.DatabaseConfig) *gorm.DB {
	dsn := dbConfig.GetDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Get underlying sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connected successfully")
	DB = db
	return db
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return fmt.Errorf("failed to get database instance: %w", err)
		}
		return sqlDB.Close()
	}
	return nil
}
