package service

import . "supermarket/types"

// Gets all stored produce, returning the list of produce
// and nil if successful.
type ProduceGetter interface {
	GetProduce() []ProduceItem
}

// Saves the given items.
// Returns a slice of successfully inserted items.
type ProducePoster interface {
	PostProduce(items []ProduceItem) []ProduceItem
}

// Deletes the item with the given produce code.
// Returns true if successful, else false.
type ProduceDeleter interface {
	DeleteProduce(code string) bool
}
