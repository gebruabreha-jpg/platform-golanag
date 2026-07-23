package model

import "time"

type URL struct {
	Code      string    `json:"code"`
	LongURL   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
}