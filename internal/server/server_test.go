package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	InitRoutes(router)
	return router
}

func TestGetVideos(t *testing.T) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/v1/videos", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %v", response.Code)
	}
}

func TestFetchAndSaveVideos(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("POST", "/v1/videos/fetch", nil) // Adjust request body as needed
	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %v", response.Code)
	}
}
