package model

import "time"

// ContentType specifies what type of content an item contains.
type ContentType int

const (
	// URLItemType represents single URL items.
	URLItemType = iota

	// TextItemType is a text blob, possibly Markdown or Code snippets.
	TextItemType
)

// Item represents a single item stored and accessible via a short URL.
type Item struct {
	Code      string      `json:"code"`
	Type      ContentType `json:"type"`
	Content   string      `json:"content"`
	CreatedAt time.Time   `json:"created_at"`
}
