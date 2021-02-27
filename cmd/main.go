package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Metric from k6 API
type Metric struct {
	ID         string
	Attributes struct {
		Type     string
		Contains string
		Sample   map[string]float64
	}
}

var k6EndpointMetrics = "http://localhost:6565/v1/metrics"

//vark6EndpointGroups = "http://localhost:6565/v1/groups"

func metricsHandler(resp http.ResponseWriter, req *http.Request) {

	// Call the k6 metrics API
	apiRes, err := http.Get(k6EndpointMetrics)
	if err != nil || apiRes.StatusCode != 200 {
		resp.WriteHeader(500)
		var errorMsg string
		if err != nil {
			errorMsg = err.Error()
		} else {
			errorMsg = apiRes.Status
		}
		fmt.Fprintln(resp, errorMsg)
		return
	}

	// Hold API result
	metricData := struct {
		Data []Metric
	}{}

	// JSON unmarshall
	err = json.NewDecoder(apiRes.Body).Decode(&metricData)
	if err != nil {
		resp.WriteHeader(500)
		fmt.Fprintln(resp, err.Error())
		return
	}

	// Now build Prometheus exposition format
	expoResult := ""
	for _, metric := range metricData.Data {
		metricName := metric.ID
		for sampleName, sampleValue := range metric.Attributes.Sample {
			sampleName = strings.ReplaceAll(sampleName, "(", "")
			sampleName = strings.ReplaceAll(sampleName, ")", "")
			expoResult += fmt.Sprintf("k6_%s_%s %f\n", metricName, sampleName, sampleValue)
		}
		expoResult += "\n"
	}

	// Return as HTTP response
	_, _ = resp.Write([]byte(expoResult))
}

func main() {
	serverPort := "2112"
	if portEnv := os.Getenv("METRICS_PORT"); portEnv != "" {
		serverPort = portEnv
	}

	k6ApiEndpoint := "http://localhost:6565/v1"
	if endpoint := os.Getenv("K6_API_ENDPOINT"); endpoint != "" {
		k6ApiEndpoint = endpoint + "/metrics"
	}
	k6EndpointMetrics = k6ApiEndpoint + "/metrics"

	http.HandleFunc("/metrics", metricsHandler)

	log.Printf("### k6 metrics proxy listening on %v\n", serverPort)
	log.Printf("### API endpoint for k6: %s\n", k6ApiEndpoint)
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
