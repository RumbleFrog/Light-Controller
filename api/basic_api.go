package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rumblefrog/light-controller/light"
	rpio "github.com/stianeikeland/go-rpio"
)

// Payload - Basic color changing payload
type Payload struct {
	R rpio.State `json:"r"`
	G rpio.State `json:"g"`
	B rpio.State `json:"b"`
}

// Register - Starts the basic API server
func Register() {
	log.Print("Test")

	http.HandleFunc("/", ColorChange)

	log.Fatal(http.ListenAndServe(":8080", nil))

	log.Print("Serving at 8080")
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

	light.R.Write(Content.R)
	light.G.Write(Content.G)
	light.B.Write(Content.B)

	w.WriteHeader(http.StatusNoContent)
}
