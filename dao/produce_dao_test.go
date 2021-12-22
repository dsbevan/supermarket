package dao

import (
	. "supermarket/types"
	"testing"
)

func contains(slice []ProduceItem, produce ProduceItem) bool {
	for _, item := range slice {
		if item.Name == produce.Name && item.Code == produce.Code && item.Price == produce.Price {
			return true
		}
	}
	return false
}

func equivalent(first []ProduceItem, second []ProduceItem) (bool, string) {
	if len(first) != len(second) {
		return false, "lengths differ"
	}
	for _, item := range first {
		if !contains(second, item) {
			return false, "contents differ"
		}
	}
	for _, item := range second {
		if !contains(first, item) {
			return false, "contents differ"
		}
	}
	return true, ""
}

var apple ProduceItem = ProduceItem{
	Name:  "apple",
	Code:  "SN3J-3398-2222-1111",
	Price: 12.30,
}
var pear ProduceItem = ProduceItem{
	Name:  "pear",
	Code:  "1112-3334-5556-7778",
	Price: 3.3,
}
var orange ProduceItem = ProduceItem{
	Name:  "orange",
	Code:  "8888-AAAA-BBBB-OOOO",
	Price: 2.99,
}

func TestGetProduce(t *testing.T) {
	largeSlice := make([]ProduceItem, 0, 20)
	largeSlice = append(largeSlice, apple)
	testcases := []struct {
		storedProduce []ProduceItem
		expected      []ProduceItem
	}{
		{ // Test empty array
			storedProduce: make([]ProduceItem, 0, 10),
			expected:      make([]ProduceItem, 0, 10),
		},
		{ // Test full array
			storedProduce: []ProduceItem{apple, pear, orange},
			expected:      []ProduceItem{pear, apple, orange},
		},
		{ // Test partially full array
			storedProduce: largeSlice,
			expected:      largeSlice,
		},
	}

	for _, test := range testcases {
		dao := NewProduceDao()
		dao.produce = test.storedProduce
		produce := dao.GetProduce()
		equal, msg := equivalent(produce, test.expected)
		if !equal {
			t.Errorf("Expected and actual %s", msg)
			t.Fail()
		}
	}
}

func TestPostProduce(t *testing.T) {
	// Create data structures for tests
	emptySlice := []ProduceItem{}
	fullSlice := []ProduceItem{apple, pear}
	partiallyFullSlice := make([]ProduceItem, 0, 10)
	partiallyFullSlice = append(partiallyFullSlice, apple)
	partiallyFullSlice = append(partiallyFullSlice, pear)

	testcases := []struct {
		initialContents []ProduceItem
		itemToAdd       ProduceItem
		expected        []ProduceItem
	}{
		{ // Test add to empty array
			initialContents: emptySlice,
			itemToAdd:       apple,
			expected:        []ProduceItem{apple},
		},
		{ // Test add to full array
			initialContents: fullSlice,
			itemToAdd:       orange,
			expected:        []ProduceItem{apple, orange, pear},
		},
		{ // Test add to array with space
			initialContents: partiallyFullSlice,
			itemToAdd:       orange,
			expected:        []ProduceItem{apple, orange, pear},
		},
		{ // Test add item that is already in the array
			initialContents: partiallyFullSlice,
			itemToAdd:       apple,
			expected:        []ProduceItem{apple, pear},
		},
	}

	for _, test := range testcases {
		dao := NewProduceDao()
		dao.produce = test.initialContents

		dao.PostProduce(test.itemToAdd)

		equal, msg := equivalent(dao.produce, test.expected)
		if !equal {
			t.Errorf("Expected and actual %s", msg)
			t.Fail()
		}

		// Reset test data in case it has been changed
		emptySlice = []ProduceItem{}
		fullSlice = []ProduceItem{apple, pear}
		partiallyFullSlice = make([]ProduceItem, 0, 10)
		partiallyFullSlice = append(partiallyFullSlice, apple)
		partiallyFullSlice = append(partiallyFullSlice, pear)
	}
}
