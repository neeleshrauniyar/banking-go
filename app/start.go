package app

import "net/http"

func Start() {
	mux:= http.NewServeMux()

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getallcustomers)

	http.ListenAndServe(":8000", mux)
}
