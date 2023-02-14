package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServe)
	//d := TestFunc()
	//d.AnotherFunc()

	log.Fatal(http.ListenAndServe(":7070", handler))
}
