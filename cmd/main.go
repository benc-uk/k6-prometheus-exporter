package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//const K6_METRICS = "http://localhost:6565/v1/metrics"
//const K6_GROUPS = "http://localhost:6565/v1/groups"

func metricsHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "Hi there, I love!")
}

func main() {
	serverPort := "8080"
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		serverPort = portEnv
	}

	http.HandleFunc("/metrics", metricsHandler)
	log.Printf("### Server listening on %v\n", serverPort)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", serverPort),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
