package server

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/souvik03-136/Fam-Go/internal/controllers"
	"github.com/souvik03-136/Fam-Go/internal/services"
)

func InitRoutes(router *gin.Engine) {
	err := godotenv.Load() // Load environment variables from .env file
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	youtubeService := services.NewYouTubeService(os.Getenv("YOUTUBE_API_KEY"), db)
	videoController := controllers.NewVideoController(db, youtubeService)

	router.GET("/v1/videos", videoController.GetVideos)
	router.POST("/v1/videos/fetch", videoController.FetchAndSaveVideos)
}

func setupDatabase() (*sql.DB, error) {
	connStr := "user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"
	return sql.Open("postgres", connStr)
}
