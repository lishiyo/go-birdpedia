package main

import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)
	"encoding/json"
	"fmt"
	"net/http"
)

// Bird is the main model
type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

// GET all burritos
func GetBirdsHandler(w http.ResponseWriter, r *http.Request) {
	// Convert the "birds" variable to json
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// success - write JSON list
	w.Write(birdListBytes)
}

func CreateBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new Bird instance
	bird := Bird{}

	// We send all our data as HTML form data
	// The `ParseForm` method of the request parses the form values
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from the form info
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	birds = append(birds, bird)

	// Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
