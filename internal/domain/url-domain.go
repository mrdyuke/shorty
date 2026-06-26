package domain

import "time"

type URLPack struct {
	ID          int
	OriginalURL string
	ShortURL    string
	CreatedAt   time.Time
}
