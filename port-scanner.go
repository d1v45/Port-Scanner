package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	Port   int
	Status string
}

func scanPort(host string, port int, wg *sync.WaitGroup, results chan<- ScanResult, timeout time.Duration) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		results <- ScanResult{Port: port, Status: "closed"}
		return
	}
	defer conn.Close()

	results <- ScanResult{Port: port, Status: "open"}
}

func scanPorts(host string, startPort, endPort, maxConcurrency int) {
	var wg sync.WaitGroup
	results := make(chan ScanResult, (endPort - startPort + 1))
	sem := make(chan struct{}, maxConcurrency)
	timeout := 2 * time.Second
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(port int) {
			defer func() { <-sem }()
			scanPort(host, port, &wg, results, timeout)
		}(port)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Status == "open" {
			fmt.Printf("Port %d is open\n", result.Port)
		}
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <IP> <start_port> <end_port> <max_concurrency>")
		os.Exit(1)
	}

	host := os.Args[1]
	startPort, err := strconv.Atoi(os.Args[2])
	if err != nil || startPort < 1 || startPort > 65535 {
		fmt.Println("Invalid start port")
		os.Exit(1)
	}

	endPort, err := strconv.Atoi(os.Args[3])
	if err != nil || endPort < 1 || endPort > 65535 || endPort < startPort {
		fmt.Println("Invalid end port")
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(os.Args[4])
	if err != nil || maxConcurrency < 1 {
		fmt.Println("Invalid concurrency value")
		os.Exit(1)
	}

	fmt.Printf("Scanning ports %d-%d on host %s with max concurrency of %d...\n", startPort, endPort, host, maxConcurrency)
	scanPorts(host, startPort, endPort, maxConcurrency)
	fmt.Println("Port scan completed.")
}
