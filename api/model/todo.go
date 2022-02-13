package model

import (
	"time"
)

type Todo struct {
	ID uint64
	Content string
	Done bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
