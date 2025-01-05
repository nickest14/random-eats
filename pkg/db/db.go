package db

import (
	"fmt"
	"os"

	"time"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Restaurant struct {
	gorm.Model
	Name      string    `gorm:"unique;not null"`
	City      string    `gorm:"not null"`
	District  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"index;autoCreateTime"`
}

func getDSN() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
}

// Connect initializes the database connection
func Connect() (*gorm.DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, fmt.Errorf("error loading env file: %v", err)
	}

	dsn := getDSN()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&Restaurant{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return DB, nil
}

// Close closes the database connection
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
