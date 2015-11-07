package main

import "time"

// ItemType represents what kind of content is in an Item.
type ItemType int

const (
	// URLItemType represents single URL items.
	URLItemType ItemType = iota

	// TextItemType is a text blob, possibly Markdown or Code snippets.
	TextItemType
)

// Item represents a single item stored and accessible via a short URL.
type Item struct {
	Code      string
	Type      ItemType
	Content   string
	CreatedAt time.Time
}
