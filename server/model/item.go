package model

import "time"

const (
	// URLItemType represents single URL items.
	URLItemType = iota

	// TextItemType is a text blob, possibly Markdown or Code snippets.
	TextItemType
)

// Item represents a single item stored and accessible via a short URL.
type Item struct {
	ID        string `jsonapi:"-"`
	Type      int
	Content   string
	CreatedAt time.Time `jsonapi:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (i Item) GetID() string {
	return i.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (i *Item) SetID(id string) error {
	i.ID = id
	return nil
}
