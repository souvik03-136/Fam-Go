package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/souvik03-136/Fam-Go/internal/merrors"
)

type Config struct {
	APIKeys       []string
	DatabaseURL   string
	FetchInterval int
	DB            *sql.DB
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	apiKeys := []string{
		os.Getenv("YOUTUBE_API_KEY_1"),
		os.Getenv("YOUTUBE_API_KEY_2"),
		os.Getenv("YOUTUBE_API_KEY_3"),
	}

	if err := validateYouTubeAPIKeys(apiKeys); err != nil {
		merrors.InternalServer(nil, "YouTube API validation failed: "+err.Error())
		return nil
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		merrors.BadRequest(nil, "DATABASE_URL environment variable is required")
		return nil
	}

	fetchInterval := 10
	if interval := os.Getenv("FETCH_INTERVAL"); interval != "" {
		if value, err := strconv.Atoi(interval); err == nil {
			fetchInterval = value
		} else {
			log.Printf("Invalid FETCH_INTERVAL value, defaulting to 10: %v", err)
		}
	}

	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		merrors.InternalServer(nil, "Failed to connect to the database: "+err.Error())
		return nil
	}

	if err := db.Ping(); err != nil {
		merrors.InternalServer(nil, "Failed to ping the database: "+err.Error())
		return nil
	}

	log.Println("Successfully connected to the database")

	return &Config{
		APIKeys:       apiKeys,
		DatabaseURL:   databaseURL,
		FetchInterval: fetchInterval,
		DB:            db,
	}
}

func validateYouTubeAPIKeys(apiKeys []string) error {
	for _, key := range apiKeys {
		if key != "" {
			resp, err := http.Get("https://www.googleapis.com/youtube/v3/videos?part=id&key=" + key)
			if err == nil && resp.StatusCode == http.StatusOK {
				return nil
			}
		}
	}
	return fmt.Errorf("no valid YouTube API keys found")
}

func (c *Config) CloseDB() {
	if c.DB != nil {
		if err := c.DB.Close(); err != nil {
			log.Println("Error closing database connection:", err)
		}
	}
}
