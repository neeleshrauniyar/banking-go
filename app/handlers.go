package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	ZipCode int `json:"zipCode" xml:"zipCode"`
}

func getallcustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"John", "London", 12345},
		{"Jane", "Paris", 67890},
	}

	if r.Header.Get("Content-Type") != "application/xml" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
		return
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}
}