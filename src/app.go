package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		message := "Hello from " + hostname
		w.Write([]byte(message))
	})

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
