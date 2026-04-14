package model

import "time"

type Link struct {
	Id     uint   `json:"-" db:"id"`
	Url    string `json:"url" db:"original_url"`
	Visits uint   `json:"visits" db:"visits"`
}

type LinkStats struct {
	Link
	Code      string    `json:"short_code"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
