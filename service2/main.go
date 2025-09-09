package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"time"
)

var startTime = time.Now()

func getUptimeHours() float64 {
	elapsed := time.Since(startTime)
	return elapsed.Hours()
}

func getFreeDiskMB() uint64 {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	if err != nil {
		return 0
	}
	return stat.Bavail * uint64(stat.Bsize) / 1024 / 1024
}

func getTimestamp() string {
	// ISO 8601 UTC, no milliseconds
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func createStatusRecord() string {
	timestamp := getTimestamp()
	uptime := getUptimeHours()
	freeDisk := getFreeDiskMB()
	return fmt.Sprintf("%s: uptime %.2f hours, free disk in root: %d MBytes", timestamp, uptime, freeDisk)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	record := createStatusRecord()

	storageURL := os.Getenv("STORAGE_URL")
	if storageURL == "" {
		storageURL = "http://storage:5000"
	}
	_, err := http.Post(storageURL+"/log", "text/plain", bytes.NewBufferString(record))
	if err != nil {
		fmt.Printf("Error posting to storage: %v\n", err)
	}

	file, err := os.OpenFile("/vstorage/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening vStorage file: %v\n", err)
	} else {
		defer file.Close()
		file.WriteString(record + "\n")
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%s", record)
}

func main() {
	http.HandleFunc("/status", statusHandler)
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
