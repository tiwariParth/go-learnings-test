package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/parth/go-learnings/week-11/job-runner/api"
	"github.com/parth/go-learnings/week-11/job-runner/config"
	"github.com/parth/go-learnings/week-11/job-runner/queue"
	"github.com/parth/go-learnings/week-11/job-runner/worker"
)

var (
	configFile string
	port       int
	workers    int
)

func init() {
	flag.StringVar(&configFile, "config", "config.yaml", "Path to configuration file")
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.IntVar(&workers, "workers", 5, "Number of worker goroutines")
}

func main() {
	// Parse command line flags
	flag.Parse()
	
	// Load configuration
	conf, err := config.Load(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	
	// Override config with command line flags
	if port != 8080 {
		conf.Port = port
	}
	
	if workers != 5 {
		conf.Workers = workers
	}
	
	// Initialize job queue
	jobQueue := queue.New(conf.QueueSize)
	
	// Start worker pool
	workerPool := worker.NewPool(conf.Workers, jobQueue)
	workerPool.Start()
	
	// Create API server
	apiServer := api.New(conf, jobQueue)
	
	// Start the server
	addr := fmt.Sprintf(":%d", conf.Port)
	fmt.Printf("Starting job runner on %s\n", addr)
	
	err = http.ListenAndServe(addr, apiServer.Router())
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
