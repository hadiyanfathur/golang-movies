package main

import (
	"net/http"

	"github.com/hadiyanfathur/golang-movies/helper"
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    "localhost:8088",
		Handler: handler,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
