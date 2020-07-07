package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayStore())

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listen on port %v", err)
	}
}
