package service

import (
	"supermarket/testutils"
	. "supermarket/types"
	"sync"
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

// dao.ProduceGetter mock
type MockProduceGetter struct {
	data []ProduceItem
}

func NewMockProduceGetter(data []ProduceItem) *MockProduceGetter {
	return &MockProduceGetter{
		data: data,
	}
}

func (m *MockProduceGetter) GetProduce() []ProduceItem {
	return m.data
}

func TestServiceGetProduce(t *testing.T) {
	testcases := []struct {
		name         string
		mockResponse []ProduceItem
		expected     []ProduceItem
		errs         error
	}{
		{
			name:         "Simple test",
			mockResponse: []ProduceItem{apple},
			expected:     []ProduceItem{apple},
			errs:         nil,
		},
		{
			name:         "More items",
			mockResponse: []ProduceItem{apple, orange, pear},
			expected:     []ProduceItem{pear, apple, orange}, // Different order
			errs:         nil,
		},
		{
			name:         "Empty db test",
			mockResponse: []ProduceItem{},
			expected:     []ProduceItem{},
			errs:         nil,
		},
	}

	for _, test := range testcases {
		svc := NewProduceService()
		svc.produceGetter = NewMockProduceGetter(test.mockResponse)

		result := svc.GetProduce()
		equal, msg := testutils.Equivalent(result, test.expected)
		if !equal {
			t.Errorf("%s", test.name)
			t.Errorf("Expected and actual %s", msg)
			t.Errorf("Expected: %v", test.expected)
			t.Errorf("Actual: %v", result)
			t.Fail()
		}
	}
}

// dao.ProducePoster mock
type MockProducePoster struct {
	response []bool
	count    int
	mu       sync.Mutex
}

func NewMockProducePoster(res []bool) *MockProducePoster {
	return &MockProducePoster{
		response: res,
		count:    0,
	}
}

func (m *MockProducePoster) PostProduce(item ProduceItem) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	res := m.response[m.count]
	m.count += 1
	return res
}

func TestServicePostProduce(t *testing.T) {
	testcases := []struct {
		name          string
		mockResponse  []bool
		arg           []ProduceItem
		expected      []ProduceItem
		checkContents bool
	}{
		{
			name:          "Single object successful post",
			mockResponse:  []bool{true},
			arg:           []ProduceItem{apple},
			expected:      []ProduceItem{apple},
			checkContents: true,
		},
		{
			name:          "Single object unsuccessful post",
			mockResponse:  []bool{false},
			arg:           []ProduceItem{apple},
			expected:      []ProduceItem{},
			checkContents: true,
		},
		{
			name:          "Multiple objects successful post",
			mockResponse:  []bool{true, true, true},
			arg:           []ProduceItem{apple, pear, orange},
			expected:      []ProduceItem{apple, pear, orange},
			checkContents: true,
		},
		{
			name:          "Multiple objects unsuccessful post",
			mockResponse:  []bool{false, false, false},
			arg:           []ProduceItem{apple, pear, orange},
			expected:      []ProduceItem{},
			checkContents: true,
		},
		{
			name:          "Multiple objects one failure",
			mockResponse:  []bool{true, false, true},
			arg:           []ProduceItem{apple, pear, orange},
			expected:      []ProduceItem{apple, orange}, //is usually pear, orange
			checkContents: false,
		},
		// Partial failures cannot easily be fully tested using this method
		// because execution order cannot be guaranteed
	}

	for _, test := range testcases {
		svc := NewProduceService()
		svc.producePoster = NewMockProducePoster(test.mockResponse)

		result := svc.PostProduce(test.arg)

		if len(result) != len(test.expected) {
			t.Errorf("Expected and actual lengths differ. Expected: %d, Actual: %d",
				len(test.expected), len(result))
			t.Fail()
		}
		if test.checkContents {
			equal, msg := testutils.Equivalent(result, test.expected)
			if !equal {
				t.Errorf("%s", test.name)
				t.Errorf("Expected and actual %s", msg)
				t.Errorf("Expected: %v", test.expected)
				t.Errorf("Actual: %v", result)
				t.Fail()
			}
		}
	}
}

// dao.ProduceDeleter mock
type MockProduceDeleter struct {
	response bool
}

func NewMockProduceDeleter(res bool) *MockProduceDeleter {
	return &MockProduceDeleter{
		response: res,
	}
}

func (m *MockProduceDeleter) DeleteProduce(code string) bool {
	return m.response
}

func TestServiceDeleteProduce(t *testing.T) {
	testcases := []struct {
		name         string
		mockResponse bool
		arg          string
		expected     bool
	}{
		{
			name:         "Successful test",
			mockResponse: true,
			arg:          "3434-3333-4444-4343",
			expected:     true,
		},
		{
			name:         "Unsuccessful test",
			mockResponse: false,
			arg:          "3434-3333-4444-4343",
			expected:     false,
		},
	}

	for _, test := range testcases {
		svc := NewProduceService()
		svc.produceDeleter = NewMockProduceDeleter(test.mockResponse)

		result := svc.DeleteProduce(test.arg)
		if result != test.expected {
			t.Errorf("%s", test.name)
			t.Errorf("TestServiceDeleteProduce: Result does not match expected result")
			t.Errorf("Expected: %v", test.expected)
			t.Errorf("Actual: %v", result)
			t.Fail()
		}
	}
}
