package main

import (
	"devsu/server"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

type config struct {
	apiKey       string
	jwtSigninKey []byte
	port         string
}

func init() {
	logger = log.New(os.Stderr, "LOG: ", log.Lshortfile)
}

func initConfig() config {
	apiKey, present := os.LookupEnv("DEVSU_API_KEY")
	if !present {
		logger.Println("DEVSU_API_KEY environment variable not set, using default")
		apiKey = "TEST"
	}

	jwtSigninKey, present := os.LookupEnv("DEVSU_JWT_KEY")
	if !present {
		logger.Println("DEVSU_JWT_KEY environment variable not set, using default")
		jwtSigninKey = "1234"
	}

	port, present := os.LookupEnv("DEVSU_PORT")
	if !present {
		port = "5000"
	}

	return config{apiKey, []byte(jwtSigninKey), port}
}

func main() {
	mux := http.NewServeMux()

	config := initConfig()

	apiKeyMiddleware := server.ApiKeyMiddleware(config.apiKey)
	jwtMiddleware := server.JWTMiddleware(config.jwtSigninKey)

	devOpsHandler := http.HandlerFunc(server.HandleDevOps)
	mux.Handle("/DevOps", jwtMiddleware(apiKeyMiddleware(devOpsHandler)))

	jwtHandler := http.HandlerFunc(server.HandleJWT(config.jwtSigninKey))
	mux.Handle("/token", jwtHandler)

	logger.Printf("starting server in port %s", config.port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.port), mux)

	if errors.Is(err, http.ErrServerClosed) {
		logger.Println("server closed")
	} else if err != nil {
		logger.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
