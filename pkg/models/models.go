package models

import (
	"errors"
	"time"
)

// ErrNoRecord : Enable logging
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet model
type Snippet struct {
	ID		int
	Title	string
	Content	string
	Created	time.Time
	Expires	time.Time 
}
