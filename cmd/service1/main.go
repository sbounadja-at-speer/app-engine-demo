package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi, I am service 1")
	})

	http.HandleFunc("/_ah/warmup", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Warmup: %s \n", "Service 1")
	})

	log.Printf("Service 1 running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
