package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/souvik03-136/Fam-Go/internal/database"
	"github.com/souvik03-136/Fam-Go/internal/merrors"
	"github.com/souvik03-136/Fam-Go/internal/services"
)

type VideoController struct {
	DB             *database.Queries
	YouTubeService *services.YouTubeService
}

func NewVideoController(db *database.Queries, youtubeService *services.YouTubeService) *VideoController {
	return &VideoController{
		DB:             db,
		YouTubeService: youtubeService,
	}
}

// GetVideos retrieves videos from the database with pagination.
func (vc *VideoController) GetVideos(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := int32(10)
	offset := (int32(page) - 1) * limit

	videos, err := vc.DB.ListVideos(c, database.ListVideosParams{Limit: limit, Offset: offset})
	if err != nil {
		merrors.InternalServer(c, "Could not fetch videos from the database")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"videos": videos,
		"page":   page,
		"limit":  limit,
	})
}

func (vc *VideoController) FetchAndSaveVideos(c *gin.Context) {
	channelID := c.Query("channelId")
	if channelID == "" {
		merrors.BadRequest(c, "channelId is required")
		return
	}

	videos, err := vc.YouTubeService.FetchVideos(channelID)
	if err != nil {
		merrors.InternalServer(c, "Could not fetch videos from YouTube")
		return
	}

	if err := vc.YouTubeService.SaveVideosToDB(c, videos); err != nil {
		merrors.InternalServer(c, "Could not save videos to the database")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Videos fetched and saved successfully"})
}
