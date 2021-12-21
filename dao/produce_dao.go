package dao

import (
	"supermarket/data"
	. "supermarket/types"
)

type ProduceDao struct {
	produce []ProduceItem
}

func NewProduceDao() *ProduceDao {
	return &ProduceDao{
		produce: data.Produce[:],
	}
}

// Inserts the passed ProduceItem into the storage array,
// if there is space and the item does not already exist.
// Returns true if successful, else false.
func (d *ProduceDao) InsertProduce(item ProduceItem) bool {
	// TODO insert
	panic("Not implemented")
}

// Gets all the produce items.
// Returns a pointer to the array of produce items.
func (d *ProduceDao) GetProduce(code string) []*ProduceItem {
	// TODO get
	panic("Not implemented")
}

// Deletes the produce item matching the given produce code.
// Returns true if successful, else false.
func (d *ProduceDao) DeleteProduce(code string) bool {
	// TODO delete
	panic("Not implemented")
}
