package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rumblefrog/light-controller/light"
)

// Payload - Basic color changing payload
type Payload struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// Register - Starts the basic API server
func Register() {
	APIMux := http.NewServeMux()

	APIMux.HandleFunc("/", ColorChange)

	log.Print("Serving at 8080")

	log.Fatal(http.ListenAndServe(":8080", APIMux))
}

// ColorChange - Handles default route
func ColorChange(w http.ResponseWriter, r *http.Request) {
	Content := &Payload{}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Panic("Unable to read JSON body")
	}

	if err := r.Body.Close(); err != nil {
		log.Panic("Unable to close Body handle")
	}

	if err := json.Unmarshal(body, &Content); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Panic("Unable to unmarshal JSON data")
	}

	light.WriteAll(
		float64(Content.R)/255.0,
		float64(Content.G)/255.0,
		float64(Content.B)/255.0,
	)

	w.WriteHeader(http.StatusNoContent)
}
