package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

func GetPlayerScore(name string) string {
	if name == "Rebeca" {
		return "20"
	}
	if name == "Biles" {
		return "10"
	}
	return ""
}

type PlayerServer struct {
	store PlayerStore

	// We need to change our tests to instead
	// create a new instance of our PlayerServer and then call its method ServeHTTP. in the test file
}

// w is the writer, and r is the body of the request
// For our PlayerServer to be able to use a PlayerStore, it will need a reference to one.
// Now feels like the right time to change our architecture so that our PlayerServer is now a struct.
// Finally, we will now implement the Handler interface by adding a method
// to our new struct and putting in our existing handler code.

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
