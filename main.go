package main

import (
	"log"
	"net/http"
)

func main() {
	brHandler := func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "error", http.StatusBadRequest)
		return
	}

	http.HandleFunc("/", brHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
