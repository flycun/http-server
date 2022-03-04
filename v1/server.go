package main

import (
	"fmt"
	"net/http"
)

// PlayerServer currently returns Hello, world given _any_ request.
func PlayerServer(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, getPlayerScore(player))

}

func getPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}

	if player == "Floyd" {
		return "10"
	}
	return ""
}
