package main

import "time"

const (
	// URLItemType represents single URL items.
	URLItemType = iota

	// TextItemType is a text blob, possibly Markdown or Code snippets.
	TextItemType
)

// Item represents a single item stored and accessible via a short URL.
type Item struct {
	Code      string    `jsonapi:"primary,items"`
	Type      int       `jsonapi:"attr,type"`
	Content   string    `jsonapi:"attr,content"`
	CreatedAt time.Time `jsonapi:"attr,created_at"`
}
