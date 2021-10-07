package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func check(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	fmt.Printf("%s %d\n", r.RemoteAddr, http.StatusOK)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", check)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
