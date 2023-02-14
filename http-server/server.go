package main

import (
	"fmt"
	"net/http"
)

func PlayerServe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

type TestFunc func()

func (f TestFunc) AnotherFunc() {
	fmt.Println("Hello World")
}
