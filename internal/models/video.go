package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Video struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	PublishedAt  time.Time `json:"published_at"`
	ThumbnailURL string    `json:"thumbnail_url"`
	VideoURL     string    `json:"video_url"`
}

func NewVideo(title, description string, publishedAt time.Time, thumbnailURL, videoURL string) (Video, error) {
	if title == "" {
		return Video{}, errors.New("title cannot be empty")
	}
	if description == "" {
		return Video{}, errors.New("description cannot be empty")
	}
	return Video{
		ID:           uuid.New(),
		Title:        title,
		Description:  description,
		PublishedAt:  publishedAt,
		ThumbnailURL: thumbnailURL,
		VideoURL:     videoURL,
	}, nil
}
