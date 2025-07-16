package documents

import (
	"time"
)

type Document struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	UploadedAt time.Time `json:"uploaded_at"`
	Path       string    `json:"path"`
}

// api response model
type DocumentResource struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	URL        string    `json:"url"`
	UploadedAt time.Time `json:"uploaded_at"`
}
