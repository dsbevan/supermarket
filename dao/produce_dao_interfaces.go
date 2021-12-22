package dao

import . "supermarket/types"

// Returns slice of pointers to stored produce items.
type ProduceGetter interface {
	GetProduce() []ProduceItem
}

// Inserts a produce item into storage.
// Returns true if successful, else false.
type ProducePoster interface {
	PostProduce(item ProduceItem) bool
}

// Deletes the produce item matching the given produce code.
// Returns true if successful, else false.
type ProduceDeleter interface {
	DeleteProduce(code string) bool
}
