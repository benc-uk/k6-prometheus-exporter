package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
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

// Group from k6 API
type Group struct {
	ID         string
	Attributes struct {
		Path   string
		Name   string
		Checks []Check
	}
}

// Check from k6 API
type Check struct {
	Name   string
	Passes int
	Fails  int
}

var k6EndpointMetrics = ""
var k6EndpointGroups = ""

func metricsHandler(resp http.ResponseWriter, req *http.Request) {
	// Call the k6 metrics API
	apiRes, err := http.Get(k6EndpointMetrics)
	if err != nil || apiRes.StatusCode != 200 {
		var errorMsg string
		if err != nil {
			errorMsg = err.Error()
		} else {
			errorMsg = apiRes.Status
		}
		log.Printf("Error with downstream k6 API %s\n", errorMsg)

		// We return 200 and empty page on error, why?
		// It stops Prometheus from thinking we're dead and stopping scraping
		_, _ = resp.Write([]byte(""))
		return
	}

	// Hold API result for metrics
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

	sort.Slice(metricData.Data, func(i, j int) bool {
		return metricData.Data[i].ID < metricData.Data[j].ID
	})

	// Call the k6 groups API
	apiRes, err = http.Get(k6EndpointGroups)
	if err != nil || apiRes.StatusCode != 200 {
		var errorMsg string
		if err != nil {
			errorMsg = err.Error()
		} else {
			errorMsg = apiRes.Status
		}
		log.Printf("Error with downstream k6 API %s\n", errorMsg)

		// We return 200 and empty page on error, why?
		// It stops Prometheus from thinking we're dead and stopping scraping
		_, _ = resp.Write([]byte(""))
		return
	}

	// Hold API result for groups
	groupData := struct {
		Data []Group
	}{}

	// JSON unmarshall
	err = json.NewDecoder(apiRes.Body).Decode(&groupData)
	if err != nil {
		resp.WriteHeader(500)
		fmt.Fprintln(resp, err.Error())
		return
	}

	// Now build Prometheus exposition format
	expoResult := ""

	// Push metrics and sample values into results
	for _, metric := range metricData.Data {
		metricName := metric.ID

		for sampleName, sampleValue := range metric.Attributes.Sample {
			sampleName = strings.ReplaceAll(sampleName, "(", "")
			sampleName = strings.ReplaceAll(sampleName, ")", "")
			expoResult += fmt.Sprintf("k6_%s_%s %f\n", metricName, sampleName, sampleValue)
		}

		expoResult += "\n"
	}

	// Push group checks passes and fails
	for _, group := range groupData.Data {
		groupName := group.Attributes.Name
		if strings.TrimSpace(groupName) == "" {
			groupName = "default"
		}

		for _, check := range group.Attributes.Checks {
			expoResult += fmt.Sprintf("k6_group_check_fail{group_name=\"%s\", check_name=\"%s\"} %d\n", groupName, check.Name, check.Fails)
			expoResult += fmt.Sprintf("k6_group_check_passes{group_name=\"%s\", check_name=\"%s\"} %d\n", groupName, check.Name, check.Passes)
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
		k6ApiEndpoint = endpoint
	}
	k6EndpointMetrics = k6ApiEndpoint + "/metrics"
	k6EndpointGroups = k6ApiEndpoint + "/groups"

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
