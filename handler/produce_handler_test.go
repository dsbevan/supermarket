package handler

import (
	"encoding/json"
	"net/http"
	"testing"

	"supermarket/service"
	"supermarket/testutils"
	. "supermarket/types"
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
var missingChar ProduceItem = ProduceItem{
	Name:  "orange",
	Code:  "888-AAAA-BBBB-OOOO", //Invalid
	Price: 2.99,
}
var missingDash ProduceItem = ProduceItem{
	Name:  "orange",
	Code:  "888aAAAA-BBBB-OOOO", //Invalid
	Price: 2.99,
}
var numName ProduceItem = ProduceItem{
	Name:  "0RANG3",
	Code:  "8888-AAAA-BBBB-OOOO",
	Price: 2.99,
}
var invalidName ProduceItem = ProduceItem{
	Name:  "0RAN+E",
	Code:  "8888-AAAA-BBBB-OOOO",
	Price: 2.99,
}
var nonAlphaCode ProduceItem = ProduceItem{
	Name:  "orange",
	Code:  "8888-AAAA-B/BB-OOOO", //Invalid
	Price: 2.99,
}
var extraSet ProduceItem = ProduceItem{
	Name:  "orange",
	Code:  "8888-AAAA-BBBB-OOOO-3333", //Invalid
	Price: 2.99,
}
var extraDigit ProduceItem = ProduceItem{
	Name:  "orange",
	Code:  "8888-AAAA-BBBB-OOOO",
	Price: 2.998, //Invalid
}
var noDecimal ProduceItem = ProduceItem{
	Name:  "orange",
	Code:  "8888-AAAA-BBBB-OOOO",
	Price: 2, // Valid
}

// Helper function for mocking request bodies
func jsonBodyBytes(i interface{}) []byte {
	bytes, _ := json.Marshal(i)
	return bytes
}

// Mock http request
func makeHttpRequest(method string, body []byte) *http.Request {
	b := Body{body}
	return &http.Request{
		Method: method,
		Body:   b,
	}
}

// Mock http.Request.Body
// Implements io.ReadCloser
type Body struct {
	body []byte
}

func (b Body) Read(p []byte) (int, error) {
	for i := range b.body {
		if i < 2048 {
			p[i] = b.body[i]
		}
	}
	return len(b.body), nil
}
func (b Body) Close() error {
	return nil
}

// Mock response writer
// Implements http.ResponseWriter interface
type MockWriter struct {
	http.ResponseWriter
	buffer     *[]byte
	statusCode *int
}

func (m MockWriter) Write(bytes []byte) (int, error) {
	if *m.statusCode == 0 {
		*m.statusCode = 200
	}
	for i := 0; i < len(bytes); i++ {
		*m.buffer = append(*m.buffer, bytes[i])
	}
	//fmt.Println(string(m.buffer[:]))
	return len(bytes), nil
}
func (m MockWriter) WriteHeader(code int) {
	*m.statusCode = code
}

// Mock service.ProduceGetter
type MockProduceGetter struct {
	produce []ProduceItem
}

func (m MockProduceGetter) GetProduce() []ProduceItem {
	return m.produce
}

func TestHandleGet(t *testing.T) {
	testcases := []struct {
		name           string
		responseWriter MockWriter
		request        *http.Request
		produceGetter  service.ProduceGetter
		expected       []ProduceItem
		expectedCode   int
	}{
		{
			name:           "Single item GET",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("GET", []byte("")),
			produceGetter:  MockProduceGetter{[]ProduceItem{apple}},
			expected:       []ProduceItem{apple},
			expectedCode:   200,
		},
		{
			name:           "Multi-item GET",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("GET", []byte("")),
			produceGetter:  MockProduceGetter{[]ProduceItem{apple, pear, orange}},
			expected:       []ProduceItem{apple, orange, pear}, // Different order
			expectedCode:   200,
		},
		{
			name:           "Empty GET",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("GET", []byte("")),
			produceGetter:  MockProduceGetter{[]ProduceItem{}},
			expected:       []ProduceItem{}, // Different order
			expectedCode:   200,
		},
	}

	for _, test := range testcases {
		handler := NewProduceHandler()
		handler.produceGetter = test.produceGetter

		handler.HandleProduce(test.responseWriter, test.request)

		// Check status code
		if *test.responseWriter.statusCode != test.expectedCode {
			t.Errorf("%s", test.name)
			t.Errorf("Expected %d status code. Actual %d", test.expectedCode, *&test.responseWriter.statusCode)
			t.Fail()
		}

		res := GetProduceResponse{}
		json.Unmarshal(*test.responseWriter.buffer, &res)

		// Compare result and expected result
		same, msg := testutils.Equivalent(res.Produce, test.expected)
		if !same {
			t.Errorf("%s", test.name)
			t.Errorf("Actual and expected %s.\n Expected: %v\n Actual: %v",
				msg, res.Produce, test.expected)
			t.Fail()
		}
	}

}

// Mock service.ProducePoster
type MockProducePoster struct {
	produce []ProduceItem
}

func (m MockProducePoster) PostProduce(i []ProduceItem) []ProduceItem {
	return m.produce
}

func TestHandlePost(t *testing.T) {
	testcases := []struct {
		name           string
		responseWriter MockWriter
		request        *http.Request
		producePoster  service.ProducePoster
		expected       []ProduceItem
		expectedCode   int
	}{
		{
			name:           "Test single post success",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{apple}})),
			producePoster:  MockProducePoster{[]ProduceItem{apple}}, // Success
			expected:       []ProduceItem{apple},
			expectedCode:   200,
		},
		{
			name:           "Test single post failed to post",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{apple}})),
			producePoster:  MockProducePoster{[]ProduceItem{}}, // Empty because it failed to post (already exists or full)
			expected:       []ProduceItem{},
			expectedCode:   200,
		},
		{
			name:           "Test multiple post with some failures",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{apple, pear, orange}})),
			producePoster:  MockProducePoster{[]ProduceItem{apple, pear}}, // Orange failed
			expected:       []ProduceItem{apple, pear},
			expectedCode:   200,
		},
		{
			name:           "Test valid produce name with numbers",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{apple, numName}})), //Name with nums
			producePoster:  MockProducePoster{[]ProduceItem{apple, numName}},
			expected:       []ProduceItem{apple, numName},
			expectedCode:   200,
		},
		{
			name:           "Test invalid name",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{apple, invalidName}})), //Invalid name
			producePoster:  MockProducePoster{[]ProduceItem{apple}},                                                       // Unused
			expected:       []ProduceItem{},
			expectedCode:   400, // Bad request
		},
		{
			name:           "Test malformed request body",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes("{hello, this is a bad body}")),
			producePoster:  MockProducePoster{[]ProduceItem{apple, pear}}, // Unused
			expected:       []ProduceItem{},
			expectedCode:   400, // Bad request
		},
		{
			name:           "Test invalid produce code missing char",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{apple, missingChar}})), //missingChar
			producePoster:  MockProducePoster{[]ProduceItem{apple, pear}},                                                 // Unused
			expected:       []ProduceItem{},
			expectedCode:   400, // Bad request
		},
		{
			name:           "Test invalid produce code missing dash",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{apple, missingDash}})), //missingChar
			producePoster:  MockProducePoster{[]ProduceItem{apple}},                                                       // Unused
			expected:       []ProduceItem{},
			expectedCode:   400, // Bad request
		},
		{
			name:           "Test non-alphanumeric code",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{nonAlphaCode}})),
			producePoster:  MockProducePoster{[]ProduceItem{apple}}, // Unused
			expected:       []ProduceItem{},
			expectedCode:   400, // Bad request
		},
		{
			name:           "Test extra code set",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{extraSet}})),
			producePoster:  MockProducePoster{[]ProduceItem{apple}}, // Unused
			expected:       []ProduceItem{},
			expectedCode:   400, // Bad request
		},
		{
			name:           "Test extra decimal places in price",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{extraDigit}})),
			producePoster:  MockProducePoster{[]ProduceItem{apple}}, // Unused
			expected:       []ProduceItem{},
			expectedCode:   400, // Bad request
		},
		{
			name:           "Test no decimal places in price",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", jsonBodyBytes(PostProduceRequest{[]ProduceItem{noDecimal}})),
			producePoster:  MockProducePoster{[]ProduceItem{noDecimal}},
			expected:       []ProduceItem{noDecimal},
			expectedCode:   200,
		},
		{
			name:           "Test post with no body",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("POST", nil),             // No body
			producePoster:  MockProducePoster{[]ProduceItem{orange}}, // Unused
			expected:       []ProduceItem{},
			expectedCode:   400, // Bad request
		},
	}

	for _, test := range testcases {
		handler := NewProduceHandler()
		handler.producePoster = test.producePoster

		handler.HandleProduce(test.responseWriter, test.request)

		// Check status code
		if *test.responseWriter.statusCode != test.expectedCode {
			t.Errorf("%s", test.name)
			t.Errorf("Expected %d status code. Actual %d", test.expectedCode, *test.responseWriter.statusCode)
			t.Fail()
		}

		res := PostProduceResponse{}
		json.Unmarshal(*test.responseWriter.buffer, &res)

		// Compare result and expected result
		same, msg := testutils.Equivalent(res.Produce, test.expected)
		if !same {
			t.Errorf("%s", test.name)
			t.Errorf("Expected and actual %s.\n Expected: %v\n Actual: %v",
				msg, test.expected, res.Produce)
			t.Errorf("%s", string(*test.responseWriter.buffer))
			t.Fail()
		}
	}
}

// Mock service.ProduceDeleter
type MockProduceDeleter struct {
	returnVal bool
}

func (m MockProduceDeleter) DeleteProduce(code string) bool {
	return m.returnVal
}

func TestHandleDelete(t *testing.T) {
	testcases := []struct {
		name           string
		responseWriter MockWriter
		request        *http.Request
		produceDeleter service.ProduceDeleter
		expected       bool
		expectedCode   int
	}{
		{
			name:           "Simple delete test",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("DELETE", []byte(`{"code":"aaaa-bbbb-cccc-dddd"}`)),
			produceDeleter: MockProduceDeleter{true},
			expected:       true,
			expectedCode:   200,
		},
		{
			name:           "Simple delete failure test/code not found",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("DELETE", []byte(`{"code":"aaaa-bbbb-cccc-dddd"}`)),
			produceDeleter: MockProduceDeleter{false},
			expected:       false,
			expectedCode:   200,
		},
		// The following test uses the same validation mechanism as POST,
		//  so other error cases are covered.
		{
			name:           "Test delete with an invalid code",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("DELETE", []byte(`{"code":"aaaa-bbbb-cccc.dddd"}`)),
			produceDeleter: MockProduceDeleter{false}, // Unused
			expected:       false,                     // Unused
			expectedCode:   400,
		},
		{
			name:           "Test delete with invalid json request body",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("DELETE", []byte(`{"code":"aaaa-bbbb-cccc.dddd"`)),
			produceDeleter: MockProduceDeleter{false}, // Unused
			expected:       false,                     // Unused
			expectedCode:   400,
		},
		{
			name:           "Test delete with no body",
			responseWriter: MockWriter{buffer: new([]byte), statusCode: new(int)},
			request:        makeHttpRequest("DELETE", nil),
			produceDeleter: MockProduceDeleter{false}, // Unused
			expected:       false,                     // Unused
			expectedCode:   400,
		},
	}

	for _, test := range testcases {
		handler := NewProduceHandler()
		handler.produceDeleter = test.produceDeleter

		handler.HandleProduce(test.responseWriter, test.request)

		// Check status code
		if *test.responseWriter.statusCode != test.expectedCode {
			t.Errorf("%s", test.name)
			t.Errorf("Expected %d status code. Actual %d", test.expectedCode, *test.responseWriter.statusCode)
			t.Fail()
		}

		res := DeleteProduceResponse{}
		json.Unmarshal(*test.responseWriter.buffer, &res)

		// Compare expected and actual result
		if res.Success != test.expected {
			t.Errorf("%s", test.name)
			t.Errorf("Expected and actual results differ. Expected: %v, Actual: %v", test.expected, res.Success)
			t.Errorf("%s", string(*test.responseWriter.buffer))
			t.Fail()
		}

	}
}
