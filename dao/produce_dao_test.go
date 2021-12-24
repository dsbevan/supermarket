package dao

import (
	"fmt"
	"supermarket/testutils"
	. "supermarket/types"
	"testing"
)

// Test data
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

func TestDaoGetProduce(t *testing.T) {
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
		dao.produce = &test.storedProduce
		produce := dao.GetProduce()
		equal, msg := testutils.Equivalent(produce, test.expected)
		if !equal {
			t.Errorf("Expected and actual produce array %s", msg)
			t.Fail()
		}
	}
}

func TestDaoPostProduce(t *testing.T) {
	// Create data structures for tests
	partiallyFullSlice := make([]ProduceItem, 0, 10)
	partiallyFullSlice = append(partiallyFullSlice, apple)
	partiallyFullSlice = append(partiallyFullSlice, pear)

	testcases := []struct {
		initialContents []ProduceItem
		itemToAdd       ProduceItem
		expected        []ProduceItem
	}{
		{ // Test add to empty array (exceeding capacity)
			initialContents: []ProduceItem{},
			itemToAdd:       apple,
			expected:        []ProduceItem{apple},
		},
		{ // Test add to full array
			initialContents: []ProduceItem{apple, pear},
			itemToAdd:       orange,
			expected:        []ProduceItem{apple, orange, pear}, // Different order
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
		{ // Test add item that should be inserted in the middle of the array
			initialContents: partiallyFullSlice,
			itemToAdd:       apple,
			expected:        []ProduceItem{apple, pear},
		},
	}

	for _, test := range testcases {
		// Reset test data in case it has been changed
		partiallyFullSlice = make([]ProduceItem, 0, 10)
		partiallyFullSlice = append(partiallyFullSlice, apple)
		partiallyFullSlice = append(partiallyFullSlice, pear)

		// test
		dao := NewProduceDao()
		dao.produce = &test.initialContents

		dao.PostProduce(test.itemToAdd)

		equal, msg := testutils.Equivalent(*dao.produce, test.expected)
		if !equal {
			t.Errorf("Expected and actual produce array %s", msg)
			fmt.Printf("Expected: %v\n", test.expected)
			fmt.Printf("Actual: %v\n", dao.produce)
			t.Fail()
		}

	}
}

func TestDaoDeleteProduce(t *testing.T) {
	// Create data structures for tests
	partiallyFullSlice := make([]ProduceItem, 0, 10)
	partiallyFullSlice = append(partiallyFullSlice, apple)
	partiallyFullSlice = append(partiallyFullSlice, pear)

	testcases := []struct {
		initialContents []ProduceItem
		codeToDelete    string
		expected        []ProduceItem
	}{
		{ // Test delete last item
			initialContents: []ProduceItem{apple, pear},
			codeToDelete:    "1112-3334-5556-7778", //pear
			expected:        []ProduceItem{apple},
		},
		{ // Test delete first item
			initialContents: []ProduceItem{apple, pear},
			codeToDelete:    "SN3J-3398-2222-1111", //apple
			expected:        []ProduceItem{pear},
		},
		{ // Test delete invalid code
			initialContents: []ProduceItem{apple, pear},
			codeToDelete:    "0000-0000-2222-1111",      //invalid
			expected:        []ProduceItem{pear, apple}, // Different order
		},
		{ // Test delete last remaining item
			initialContents: []ProduceItem{apple},
			codeToDelete:    "SN3J-3398-2222-1111", //apple
			expected:        []ProduceItem{},
		},
		{ // Test delete from empty array
			initialContents: []ProduceItem{},
			codeToDelete:    "SN3J-3398-2222-1111", //apple
			expected:        []ProduceItem{},
		},
	}
	for _, test := range testcases {
		// Reset test data in case it has been changed
		partiallyFullSlice = make([]ProduceItem, 0, 10)
		partiallyFullSlice = append(partiallyFullSlice, apple)
		partiallyFullSlice = append(partiallyFullSlice, pear)

		// test
		dao := NewProduceDao()
		dao.produce = &test.initialContents

		dao.DeleteProduce(test.codeToDelete)

		equal, msg := testutils.Equivalent(*dao.produce, test.expected)
		if !equal {
			t.Errorf("Expected and actual produce array %s", msg)
			fmt.Printf("Expected: %v\n", test.expected)
			fmt.Printf("Actual: %v\n", dao.produce)
			t.Fail()
		}
	}
}

//var apple ProduceItem = ProduceItem{
//	Name:  "apple",
//	Code:  "SN3J-3398-2222-1111",
//	Price: 12.30,
//}
//var pear ProduceItem = ProduceItem{
//	Name:  "pear",
//	Code:  "1112-3334-5556-7778",
//	Price: 3.3,
//}
//var orange ProduceItem = ProduceItem{
//	Name:  "orange",
//	Code:  "8888-AAAA-BBBB-OOOO",
//	Price: 2.99,
//}

// Test for unintended state change behavior
func TestDaoAllMutatingMethods(t *testing.T) {
	testcases := []struct {
		initialContents []ProduceItem
		firstOp         string
		firstArg        interface{}
		secondOp        string
		secondArg       interface{}
		expected        []ProduceItem
	}{
		{ // Test add then add the same again
			initialContents: []ProduceItem{apple},
			firstOp:         "post",
			firstArg:        pear,
			secondOp:        "post",
			secondArg:       pear,
			expected:        []ProduceItem{apple, pear},
		},
		{ // Test delete then add different item
			initialContents: []ProduceItem{apple, pear},
			firstOp:         "delete",
			firstArg:        "1112-3334-5556-7778", //pear
			secondOp:        "post",
			secondArg:       orange,
			expected:        []ProduceItem{orange, apple}, // different order
		},
		{ // Test delete the same item twice
			initialContents: []ProduceItem{apple, pear},
			firstOp:         "delete",
			firstArg:        "1112-3334-5556-7778", //pear
			secondOp:        "delete",
			secondArg:       "1112-3334-5556-7778", //pear
			expected:        []ProduceItem{apple},
		},
	}

	for _, test := range testcases {
		dao := NewProduceDao()
		dao.produce = &test.initialContents

		switch test.firstOp {
		case "post":
			dao.PostProduce(test.firstArg.(ProduceItem))
		case "delete":
			dao.DeleteProduce(test.firstArg.(string))
		}

		switch test.secondOp {
		case "post":
			dao.PostProduce(test.secondArg.(ProduceItem))
		case "delete":
			dao.DeleteProduce(test.secondArg.(string))
		}

		equal, msg := testutils.Equivalent(*dao.produce, test.expected)
		if !equal {
			t.Errorf("Expected and actual produce array %s", msg)
			fmt.Printf("Expected: %v\n", test.expected)
			fmt.Printf("Actual: %v\n", dao.produce)
			t.Fail()
		}
	}

}
