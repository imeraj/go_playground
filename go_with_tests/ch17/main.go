package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{&InMemoryPlayStore{}}

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listen on port 500 %v", err)
	}
}
