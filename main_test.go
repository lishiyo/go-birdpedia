package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test that /hello goes to the correct route
func TestHandler(t *testing.T) {
	// Instantiate the router from main.go
	r := newRouter()

	// Create a new server using the "httptest" libraries `NewServer` method
	// Documentation : https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)

	// The mock server we created runs a server and exposes its location in the URL attribute
	// We make a GET request to the "hello" route we defined in the router
	resp, err := http.Get(mockServer.URL + "/hello")

	// Handle any unexpected error
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// the response body is read, and converted to a string
	defer resp.Body.Close()
	// read the body into a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// convert the bytes into string
	respString := string(b)
	expected := "Hello World!"

	// We want our response to match the one defined in our handler.
	// If it does happen to be "Hello world!", then it confirms, that the
	// route is correct
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	// Handle any unexpected error
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 404
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be StatusMethodNotAllowed, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	// We want our response to match the one defined in our handler.
	// If it does happen to be "Hello world!", then it confirms, that the
	// route is correct
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// We want to hit the `GET /assets/` route to get the index.html file response
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	// It isn't wise to test the entire content of the HTML file.
	// Instead, we test that the content-type header is "text/html; charset=utf-8" so that we know that an html file has been served
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"
	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}

// Test that /api/birds goes to the correct route
func TestApiGetBirdsHandler(t *testing.T) {
	// Instantiate the router from main.go
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/api/birds")
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// the response body is read, and converted to a string
	defer resp.Body.Close()
	// read the body into a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// convert the bytes into string
	respString := string(b)

	// We want our response to match the one defined in our handler.
	// If it does happen to be "Hello world!", then it confirms, that the
	// route is correct
	if respString != "null" {
		t.Errorf("Response should be null, got %s", respString)
	}
}
