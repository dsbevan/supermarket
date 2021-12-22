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

// Implements dao.ProduceGetter
// Gets all the produce items.
// Returns a slice of produce item pointers.
func (d *ProduceDao) GetProduce() []*ProduceItem {
	// TODO get
	panic("Not implemented")
}

// Implements dao.ProducePoster
// Inserts the passed ProduceItem into the storage array,
// if there is space and the item does not already exist.
// Returns true if successful, else false.
func (d *ProduceDao) PostProduce(item ProduceItem) bool {
	// TODO insert
	panic("Not implemented")
}

// Implements dao.ProduceDeleter
// Deletes the produce item matching the given produce code.
// Returns true if successful, else false.
func (d *ProduceDao) DeleteProduce(code string) bool {
	// TODO delete
	panic("Not implemented")
}
