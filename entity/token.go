package entity

import (
	"time"
)

type Token struct {
	ID          int
	UserID      int
	Name        string
	TokenNumber int
	TokenDate   time.Time
}
