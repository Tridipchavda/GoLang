package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)
// Struct to test search Query for Wikipedia
type testQuery struct {
	arg    string
	expected bool
}

// Test Cases for Search
var testQueries = []testQuery{
	{"a", true},
	{" ", false},
	{" dsiodok dcoscdso  dsoicidscok cdsokcndsoc cdsokncdsokc csdocndso", false},
	{"usa", true},
	{"", false},
}

// Struct to test http Request
type testHTTP struct {
	argSearch      string
	expectedCode   int
	expectedResult int
}

// Test Cases for HTTP
var testHTTPs = []testHTTP{
	{"", 200, 0},
	{"usa", 200, 1},
	{"elkfekfe ferwklfcerlk vodfkfn ofndrsvnko inovsdfvon erfoeinvf", 200, 0},
}

// Function to test for Data from wikipedia API
func TestGetDataFromWikipediaAPI(t *testing.T) {
	// Iterating through Test Cases
	for _, test := range testQueries {
		actual_arr := test.arg
		// Check the response and error from GetWikipediaSearch function
		resp, err := GetWikipediaSearch(actual_arr)
		if err != nil {
			t.Error(test.arg, err.Error())
			return
		}
		// If results are not expected print test case and Error
		if len(resp.Name) == 0 && test.expected {
			t.Errorf("No Result found for '%v'", test.arg)
		}
	}
}

// Function to test the HandleSearch 
func TestHandleSearch(t *testing.T) {
	// Iterating Thorugh test cases
	for _, v := range testHTTPs {
		// Make new request
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal("failed to create a http request for ", v)
		}
		// Make new Response Writer
		r := httptest.NewRecorder()
		// Make Handler and pass request and response
		handler := http.HandlerFunc(HandleSearch)
		handler.ServeHTTP(r, req)
		// Check the expected Case
		if r.Code != v.expectedCode {
			t.Errorf("test failed as expected code is %d but actually code is %d", v.expectedCode, r.Code)
		}
	}
}

// Function to test the HandleWikipediaContent
func TestHandleWikiPediaContent(t *testing.T) {
	// Iterate over test cases
	for _, v := range testHTTPs {
		// Make new Request and ResponseWriter
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal("failed to create a http request for ", v)
		}
		r := httptest.NewRecorder()
		// Create new Handler and Implement ServerHTTP method with Request and Response
		handler := http.HandlerFunc(HandleWikiPediaContent)
		handler.ServeHTTP(r, req)
		// Check the code and error
		if r.Code != v.expectedCode {
			t.Errorf("test failed as expected code is %d but actually code is %d", v.expectedCode, r.Code)
		}
	}
}

// function to Handle Wikipedia Search
func TestGetWikipediaSearch(t *testing.T) {
	// Iterate over test cases
	for _, v := range testQueries {
		// Test GetWikipediaContent for arguments in Testcases
		_, err := GetWikipediaContent(v.arg)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}
