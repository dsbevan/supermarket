package service

import . "supermarket/types"

// Gets all stored produce, returning the list of produce
// and nil if successful.
type ProduceGetter interface {
	GetProduce() ([]*ProduceItem, error)
}

// Saves the given item, returning the item inserted if successful.
type ProduceInserter interface {
	InsertProduce(item *ProduceItem) (*ProduceItem, error)
}

// Deletes the item with the given produce code.
// Returns true if successful, else false.
type ProduceDeleter interface {
	DeleteProduce(code string) (bool, error)
}
