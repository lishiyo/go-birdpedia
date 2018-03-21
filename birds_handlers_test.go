package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

// Test GET birds
func TestBirdsHandler(t *testing.T) {
	// overwrites the global `birds` in birds_handler
	birds = []Bird{
		{"sparrow", "A small harmless bird"},
	}

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	// run the test
	hf := http.HandlerFunc(GetBirdsHandler)
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Bird{"sparrow", "A small harmless bird"}

	b := []Bird{}
	err = json.NewDecoder(recorder.Body).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	// get the bird that was decoded into it
	actual := b[0]
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestCreateBirdHandler(t *testing.T) {
	// overwrites the global `birds` in birds_handler
	birds = []Bird{
		{"sparrow", "A small harmless bird"},
	}

	form := newCreateBirdForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	// run the test
	hf := http.HandlerFunc(CreateBirdHandler)
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFound)
	}

	expected := []Bird{
		{"sparrow", "A small harmless bird"},
		{"eagle", "A bird of prey"},
	}

	// should be in our global birds array now
	if !reflect.DeepEqual(birds, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", birds, expected)
	}
}

// Return a pointer to a form
func newCreateBirdForm() *url.Values {
	// Values maps a string key to a list of values - map[string][]string
	// It is typically used for query parameters and form values.
	form := url.Values{}
	form.Set("species", "eagle")
	form.Set("description", "A bird of prey")
	return &form
}
