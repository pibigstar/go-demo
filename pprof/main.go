package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	"github.com/varstr/uaparser"
)

const hostPort = ":8080"

func main() {
	http.HandleFunc("/hello", Hello)
	fmt.Println("Starting server on", hostPort)
	if err := http.ListenAndServe(hostPort, nil); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	tags := getStatsTags(r)
	duration := time.Since(start)
	fmt.Println(tags, duration)
}

func getStatsTags(r *http.Request) map[string]string {
	userBrowser, userOS := parseUserAgent(r.UserAgent())
	stats := map[string]string{
		"browser":  userBrowser,
		"os":       userOS,
		"endpoint": filepath.Base(r.URL.Path),
	}

	hostName, _ := os.Hostname()

	if hostName != "" {
		stats["host"] = hostName
	}
	return stats
}

func parseUserAgent(uaString string) (browser, os string) {
	ua := uaparser.Parse(uaString)

	if ua.Browser != nil {
		browser = ua.Browser.Name
	}
	if ua.OS != nil {
		os = ua.OS.Name
	}

	return browser, os
}
