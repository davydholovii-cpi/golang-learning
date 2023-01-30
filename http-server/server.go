package main

import (
	"net/http"
)

type Server interface {
	ListenAdnServe(addr string, handler Handler) error
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
