package models

import "time"

type Link struct {
	ID            string    `json:"id"`
	FullLink      string    `json:"full_link"`
	ShortenerLink string    `json:"shortener_link"`
	Visits        int       `json:"visitors"`
	CreatedAt     time.Time `json:"created_at"`
}

type LinkRequest struct {
	Link string `json:"urls"`
}
