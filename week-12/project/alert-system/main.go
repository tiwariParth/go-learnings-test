package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	
	"github.com/parth/go-learnings/week-12/alert-system/api"
	"github.com/parth/go-learnings/week-12/alert-system/config"
	"github.com/parth/go-learnings/week-12/alert-system/db"
	"github.com/parth/go-learnings/week-12/alert-system/router"
	"github.com/parth/go-learnings/week-12/alert-system/webhook"
)

var (
	configFile  string
	port        int
	dbPath      string
	development bool
)

func init() {
	flag.StringVar(&configFile, "config", "config.yaml", "Path to configuration file")
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&dbPath, "db", "./alerts.db", "Path to database file")
	flag.BoolVar(&development, "dev", false, "Development mode")
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
	
	if dbPath != "./alerts.db" {
		conf.Database.Path = dbPath
	}
	
	// Initialize database
	database, err := db.New(conf.Database.Path)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()
	
	// Initialize webhook manager
	webhookManager := webhook.NewManager(conf.Webhooks)
	
	// Initialize alert router
	alertRouter := router.New(database, webhookManager)
	
	// Create API server
	apiServer := api.New(conf, database, alertRouter)
	
	// Create server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: apiServer.Router(),
	}
	
	// Setup graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		fmt.Printf("Alert system starting on port %d\n", conf.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	
	// Wait for interrupt signal
	<-stop
	fmt.Println("Shutting down server...")
	
	// Close resources
	if err := database.Close(); err != nil {
		log.Printf("Database close error: %v", err)
	}
	
	fmt.Println("Server gracefully stopped")
}
