package dao

import (
	"strings"
	"supermarket/data"
	. "supermarket/types"
	"sync"
)

type ProduceDao struct {
	produce *[]ProduceItem
	mu      sync.Mutex
}

// Creates a new ProduceDao using data.Produce as the default
// data storage implementation.
func NewProduceDao() *ProduceDao {
	return &ProduceDao{
		produce: &data.ProduceSlice,
	}
}

// Implements dao.ProduceGetter
// Gets all the produce items.
// Returns a slice of produce item pointers.
func (d *ProduceDao) GetProduce() []ProduceItem {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Return a copy so that concurrently processed requests cannot change
	// the result of this get request.
	size := len(*d.produce)
	res := make([]ProduceItem, size, size)
	for i, val := range *d.produce {
		res[i] = val
	}
	return res
}

// Implements dao.ProducePoster
// Inserts the passed ProduceItem into the storage array
// if the item does not already exist.
// Returns true if successful, else false.
func (d *ProduceDao) PostProduce(item ProduceItem) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	if containsCode(*d.produce, item.Code) == -1 {
		// Format
		item.Code = strings.ToUpper(item.Code)
		*d.produce = append(*d.produce, item)
		return true
	}
	return false
}

// Implements dao.ProduceDeleter
// Deletes the produce item matching the given produce code,
// leaving all remaining items as a slice of consecutive indices.
// Order is not preserved.
// Returns true if successful, else false.
func (d *ProduceDao) DeleteProduce(code string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	idx := containsCode(*d.produce, code)
	if idx == -1 { // Code not found
		return false
	}
	var ok bool
	if *d.produce, ok = deleteProduce(*d.produce, idx); !ok {
		return false
	}
	return true
}

// Finds the produce item in the slice, returning the index of the item
// if it exists, else -1
func contains(slice []ProduceItem, produce *ProduceItem) int {
	for i, item := range slice {
		if item.Name == produce.Name && item.Code == produce.Code &&
			item.Price == produce.Price {
			return i
		}
	}
	return -1
}

// Finds the produce item in the slice by produce code (case insensitive)
// and returns the item's index or -1 if not found.
func containsCode(slice []ProduceItem, code string) int {
	code = strings.ToUpper(code)
	for i, item := range slice {
		if item.Code == code {
			return i
		}
	}
	return -1
}

func deleteProduce(slice []ProduceItem, idx int) ([]ProduceItem, bool) {
	if idx < 0 || idx > len(slice)-1 {
		return slice, false
	}
	slice[idx] = slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return slice, true
}
