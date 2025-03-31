package models

import "time"

type UrlPayload struct {
	Url            string    `json:"url" binding:"required"`
	ExpirationDate time.Time `json:"expiration_date" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
}
