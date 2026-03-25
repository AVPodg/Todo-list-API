package domain

import "time"

type Note struct {
	ID       string    `json:"id" db:"id"`
	Title    string    `json:"title" db:"title"`
	Content  string    `json:"content" db:"content"`
	CreateAt time.Time `json:"created_at" db:"created_at"`
}
