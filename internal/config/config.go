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
		log.Println("YouTube API validation failed: " + err.Error())
		return nil
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Println("DATABASE_URL environment variable is required") // Log error
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
		log.Println("Failed to connect to the database: " + err.Error()) // Log error
		return nil
	}

	if err := db.Ping(); err != nil {
		log.Println("Failed to ping the database: " + err.Error()) // Log error
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
			if err != nil {
				log.Printf("Failed to validate API key %s: %v", key, err)
				continue
			}
			defer resp.Body.Close()

			// Log the response status code for debugging
			if resp.StatusCode == http.StatusOK {
				log.Println("Successfully validated YouTube API key")
				return nil
			} else {
				log.Printf("API key %s returned status code %d", key, resp.StatusCode)
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
