package main

import (
	"devsu/server"
	"errors"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "LOG: ", log.Lshortfile)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/DevOps", server.DevOpsHandler{})

	logger.Print("starting server in port 5000")
	err := http.ListenAndServe(":5000", mux)

	if errors.Is(err, http.ErrServerClosed) {
		logger.Print("server closed\n")
	} else if err != nil {
		logger.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
