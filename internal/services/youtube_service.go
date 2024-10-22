package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/souvik03-136/Fam-Go/internal/database"
)

type YouTubeService struct {
	APIKey string
	DB     *database.Queries
}

func NewYouTubeService(apiKey string, db *database.Queries) *YouTubeService {
	return &YouTubeService{
		APIKey: apiKey,
		DB:     db,
	}
}

type YouTubeVideo struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	PublishedAt  time.Time `json:"publishedAt"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	VideoUrl     string    `json:"videoUrl"`
}

func (s *YouTubeService) FetchVideos(channelID string) ([]YouTubeVideo, error) {
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&channelId=%s&maxResults=10&order=date&type=video&key=%s", channelID, s.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data struct {
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title       string `json:"title"`
				Description string `json:"description"`
				PublishedAt string `json:"publishedAt"`
				Thumbnails  struct {
					Default struct {
						URL string `json:"url"`
					} `json:"default"`
				} `json:"thumbnails"`
			} `json:"snippet"`
		} `json:"items"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	var videos []YouTubeVideo
	for _, item := range data.Items {
		publishedAt, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		video := YouTubeVideo{
			ID:           item.ID.VideoID,
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			PublishedAt:  publishedAt,
			ThumbnailUrl: item.Snippet.Thumbnails.Default.URL,
			VideoUrl:     fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.ID.VideoID),
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func (s *YouTubeService) SaveVideosToDB(ctx context.Context, videos []YouTubeVideo) error {
	for _, video := range videos {
		createParams := database.CreateVideoParams{
			Title:        video.Title,
			Description:  sql.NullString{String: video.Description, Valid: video.Description != ""},
			PublishedAt:  video.PublishedAt,
			ThumbnailUrl: sql.NullString{String: video.ThumbnailUrl, Valid: video.ThumbnailUrl != ""},
			VideoUrl:     video.VideoUrl,
		}

		err := s.DB.CreateVideo(ctx, createParams)
		if err != nil {
			return err
		}
	}
	return nil
}
