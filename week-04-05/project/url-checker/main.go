package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

// Result represents the result of checking a single URL
type Result struct {
	URL          string
	Status       string
	StatusCode   int
	ResponseTime time.Duration
	Error        error
}

// checkURL performs an HTTP GET request to the specified URL and returns the result
func checkURL(ctx context.Context, url string) Result {
	start := time.Now()
	result := Result{URL: url}

	// Create a new request with the provided context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		result.Error = err
		result.Status = "Error"
		return result
	}

	// Perform the request
	resp, err := http.DefaultClient.Do(req)
	responseTime := time.Since(start)
	result.ResponseTime = responseTime

	if err != nil {
		result.Error = err
		result.Status = "Error"
		return result
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Status = "Success"
	} else {
		result.Status = "Failed"
	}

	return result
}

// worker processes URLs from the jobs channel and sends results to the results channel
func worker(id int, jobs <-chan string, results chan<- Result, timeout time.Duration) {
	for url := range jobs {
		fmt.Printf("Worker %d checking: %s\n", id, url)
		
		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		result := checkURL(ctx, url)
		cancel()
		
		results <- result
	}
}

func main() {
	// Define and parse command line flags
	numWorkers := flag.Int("workers", 5, "Number of worker goroutines")
	timeoutSecs := flag.Int("timeout", 10, "Timeout in seconds for each request")
	flag.Parse()

	// Get URLs from command line arguments
	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("Please provide at least one URL to check")
		fmt.Println("Usage: url-checker [flags] url1 url2 ...")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Create channels for jobs and results
	jobs := make(chan string, len(urls))
	results := make(chan Result, len(urls))

	// Create worker pool
	var wg sync.WaitGroup
	for w := 1; w <= *numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results, time.Duration(*timeoutSecs)*time.Second)
		}(w)
	}

	// Send jobs to workers
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	// Start a goroutine to close the results channel once all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	fmt.Println("\nResults:")
	fmt.Println("==========================================")
	
	for result := range results {
		if result.Error != nil {
			fmt.Printf("%-30s | ERROR: %v\n", result.URL, result.Error)
		} else {
			fmt.Printf("%-30s | Status: %d | Time: %v\n", 
				result.URL, result.StatusCode, result.ResponseTime)
		}
	}
	
	fmt.Println("==========================================")
	fmt.Println("All URLs checked!")
}
