package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
