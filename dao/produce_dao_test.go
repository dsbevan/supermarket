package dao

import (
	"supermarket/testutils"
	. "supermarket/types"
	"testing"
)

// Test data
var lowerApple ProduceItem = ProduceItem{
	Name:  "apple",
	Code:  "sn3j-3398-2222-1111",
	Price: 12.30,
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
var lemon ProduceItem = ProduceItem{
	Name:  "lemon",
	Code:  "8888-AAAA-BBBB-OOOO",
	Price: 2.00,
}

func TestDaoGetProduce(t *testing.T) {
	largeSlice := make([]ProduceItem, 0, 20)
	largeSlice = append(largeSlice, apple)
	testcases := []struct {
		name          string
		storedProduce []ProduceItem
		expected      []ProduceItem
	}{
		{
			name:          "Test empty array",
			storedProduce: make([]ProduceItem, 0, 10),
			expected:      make([]ProduceItem, 0, 10),
		},
		{
			name:          "Test full array",
			storedProduce: []ProduceItem{apple, pear, orange},
			expected:      []ProduceItem{pear, apple, orange},
		},
		{
			name:          "Test partially full array",
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
			t.Errorf("%s", test.name)
			t.Errorf("Expected and actual produce array %s", msg)
			t.Errorf("Expected: %v\n", test.expected)
			t.Errorf("Actual: %v\n", dao.produce)
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
		name            string
		initialContents []ProduceItem
		itemToAdd       ProduceItem
		expected        []ProduceItem
	}{
		{
			name:            "Test add to empty array (exceeding capacity)",
			initialContents: []ProduceItem{},
			itemToAdd:       apple,
			expected:        []ProduceItem{apple},
		},
		{
			name:            "Test add to full array",
			initialContents: []ProduceItem{apple, pear},
			itemToAdd:       orange,
			expected:        []ProduceItem{apple, orange, pear}, // Different order
		},
		{
			name:            "Test add to array with space",
			initialContents: partiallyFullSlice,
			itemToAdd:       orange,
			expected:        []ProduceItem{apple, orange, pear},
		},
		{
			name:            "Test add item that is already in the array",
			initialContents: partiallyFullSlice,
			itemToAdd:       apple,
			expected:        []ProduceItem{apple, pear},
		},
		{
			name:            "Test add item with a duplicate item code but different name and price.",
			initialContents: []ProduceItem{apple, pear, orange},
			itemToAdd:       lemon, // Same code as orange
			expected:        []ProduceItem{apple, pear, orange},
		},
		{
			name:            "Test add item that will be inserted into the middle of the array",
			initialContents: partiallyFullSlice,
			itemToAdd:       apple,
			expected:        []ProduceItem{apple, pear},
		},
		{
			name:            "Test case insensitivity",
			initialContents: []ProduceItem{},
			itemToAdd:       lowerApple,           // Insert lowercase
			expected:        []ProduceItem{apple}, // Contains uppercase
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
			t.Errorf("%s", test.name)
			t.Errorf("Expected and actual produce array %s", msg)
			t.Errorf("Expected: %v\n", test.expected)
			t.Errorf("Actual: %v\n", dao.produce)
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
		name            string
		initialContents []ProduceItem
		codeToDelete    string
		expected        []ProduceItem
	}{
		{
			name:            "Test delete last item",
			initialContents: []ProduceItem{apple, pear},
			codeToDelete:    "1112-3334-5556-7778", //pear
			expected:        []ProduceItem{apple},
		},
		{
			name:            "Test delete first item",
			initialContents: []ProduceItem{apple, pear},
			codeToDelete:    "SN3J-3398-2222-1111", //apple
			expected:        []ProduceItem{pear},
		},
		{
			name:            "Test delete invalid code",
			initialContents: []ProduceItem{apple, pear},
			codeToDelete:    "0000-0000-2222-1111",      //invalid
			expected:        []ProduceItem{pear, apple}, // Different order
		},
		{
			name:            "Test delete last remaining item",
			initialContents: []ProduceItem{apple},
			codeToDelete:    "SN3J-3398-2222-1111", //apple
			expected:        []ProduceItem{},
		},
		{
			name:            "Test delete from empty array",
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
			t.Errorf("%s", test.name)
			t.Errorf("Expected and actual produce array %s", msg)
			t.Errorf("Expected: %v\n", test.expected)
			t.Errorf("Actual: %v\n", dao.produce)
			t.Fail()
		}
	}
}

// Test for unintended state change behavior
func TestDaoAllMutatingMethods(t *testing.T) {
	testcases := []struct {
		name            string
		initialContents []ProduceItem
		firstOp         string
		firstArg        interface{}
		secondOp        string
		secondArg       interface{}
		expected        []ProduceItem
	}{
		{
			name:            "Test add then add the same again",
			initialContents: []ProduceItem{apple},
			firstOp:         "post",
			firstArg:        pear,
			secondOp:        "post",
			secondArg:       pear,
			expected:        []ProduceItem{apple, pear},
		},
		{
			name:            "Test delete then add different item",
			initialContents: []ProduceItem{apple, pear},
			firstOp:         "delete",
			firstArg:        "1112-3334-5556-7778", //pear
			secondOp:        "post",
			secondArg:       orange,
			expected:        []ProduceItem{orange, apple}, // different order
		},
		{
			name:            "Test delete the same item twice",
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
			t.Errorf("%s", test.name)
			t.Errorf("Expected and actual produce array %s", msg)
			t.Errorf("Expected: %v\n", test.expected)
			t.Errorf("Actual: %v\n", dao.produce)
			t.Fail()
		}
	}

}
