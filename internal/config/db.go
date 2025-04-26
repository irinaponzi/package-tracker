package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

var (
	err    error
	config *SQLConfig
	once   sync.Once

	getConnectionProvider = getLocalConnection
)

// SQLConfig is the struct that represents the application database configuration
type SQLConfig struct {
	DBHost        string
	DBName        string
	DBUser        string
	DBPassword    string
	ServerAddress string
	DBPort        string
}

// GetConnection is a function that returns a connection to the database
func GetConnection() (*sql.DB, error) {
	return getConnectionProvider()
}

// getLocalConnection get a local connection to the database
func getLocalConnection() (*sql.DB, error) {
	cfg, err := loadDBConfigLocal()
	if err != nil {
		return nil, err
	}

	localDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=100ms&readTimeout=1000ms&writeTimeout=100ms&parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	conn, err := sql.Open("mysql", localDSN)
	if err != nil {
		log.Fatal(err)
	}

	// health check
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	return conn, nil
}

// loadDBConfigLocal loads the configuration from environment variables
func loadDBConfigLocal() (*SQLConfig, error) {
	once.Do(func() {
		if err = godotenv.Load(); err != nil {
			err = fmt.Errorf("error loading .env file: %w", err)
			return
		}

		config = &SQLConfig{
			ServerAddress: os.Getenv("SERVER_ADDRESS"),
			DBHost:        os.Getenv("DB_HOST"),
			DBPort:        os.Getenv("DB_PORT"),
			DBUser:        os.Getenv("DB_USER"),
			DBPassword:    os.Getenv("DB_PASSWORD"),
			DBName:        os.Getenv("DB_NAME"),
		}
	})

	return config, err
}
