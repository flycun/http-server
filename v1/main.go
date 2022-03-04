package main

import (
	"log"
	"net/http"
)

type ImMemoryPlayerStore struct {
}

func (i ImMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", &PlayerServer{&ImMemoryPlayerStore{}}))
}
