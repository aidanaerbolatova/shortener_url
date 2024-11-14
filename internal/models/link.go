package models

import "time"

type Link struct {
	ID            string    `json:"id"`
	FullLink      string    `json:"full_link"`
	ShortenerLink string    `json:"shortener_link"`
	Visits        int       `json:"visitors"`
	LastVisitTime time.Time `json:"last_visit_time"`
	CreatedAt     time.Time `json:"created_at"`
}

type LinkRequest struct {
	Link string `json:"urls"`
}

type GetStatsResponse struct {
	StatsCount    int       `json:"stats_count"`
	LastVisitTime time.Time `json:"last_visit_time"`
}
