package models

import "time"

type Link struct {
	ID            string    `json:"id"`
	FullLink      string    `json:"full_link"`
	ShortenerLink string    `json:"shortener_link"`
	Visits        int       `json:"visitors"`
	LastVisitedAt time.Time `json:"last_visited_at"`
	CreatedAt     time.Time `json:"created_at"`
}

type LinkRequest struct {
	Link string `json:"url"`
}

type GetStatsResponse struct {
	StatsCount    int       `json:"stats_count"`
	LastVisitedAt time.Time `json:"last_visited_at"`
}
