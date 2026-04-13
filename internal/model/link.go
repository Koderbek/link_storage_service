package model

import "time"

type Link struct {
	Url    string `json:"url" db:"original_url"`
	Visits int    `json:"visits" db:"visits"`
}

type LinkStats struct {
	Link
	Code      string    `json:"short_code" db:"short_code"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
