package models

import "time"

type Post struct {
	ID        int
	Title     string
	Content   string
	Published bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
