package domain

import "time"

type URLpack struct {
	ID          int
	OriginalURL string
	ShortURL    string
	CreatedAt   time.Time
}
