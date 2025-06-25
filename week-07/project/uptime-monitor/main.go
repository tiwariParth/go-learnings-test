package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Website represents a site to monitor
type Website struct {
	URL           string        `json:"url"`
	Name          string        `json:"name"`
	CheckInterval time.Duration `json:"check_interval_seconds"`
	Status        string        `json:"status"`
	LastChecked   time.Time     `json:"last_checked"`
	ResponseTime  time.Duration `json:"response_time_ms"`
	Uptime        float64       `json:"uptime_percentage"`
	TotalChecks   int           `json:"total_checks"`
	SuccessChecks int           `json:"success_checks"`
}

// Monitor handles the website monitoring logic
type Monitor struct {
	websites map[string]*Website
	mu       sync.RWMutex
	stop     chan struct{}
}

// NewMonitor creates a new website monitor
func NewMonitor() *Monitor {
	return &Monitor{
		websites: make(map[string]*Website),
		stop:     make(chan struct{}),
	}
}

// AddWebsite adds a new website to monitor
func (m *Monitor) AddWebsite(w Website) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	w.Status = "Pending"
	m.websites[w.URL] = &w
	
	// Start monitoring this website
	go m.monitorWebsite(&w)
}

// GetWebsites returns all monitored websites
func (m *Monitor) GetWebsites() []*Website {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	websites := make([]*Website, 0, len(m.websites))
	for _, w := range m.websites {
		websites = append(websites, w)
	}
	
	return websites
}

// GetWebsite returns a specific monitored website
func (m *Monitor) GetWebsite(url string) (*Website, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	w, exists := m.websites[url]
	return w, exists
}

// RemoveWebsite stops monitoring a website
func (m *Monitor) RemoveWebsite(url string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	_, exists := m.websites[url]
	if exists {
		delete(m.websites, url)
	}
	
	return exists
}

// monitorWebsite continuously checks a website at the specified interval
func (m *Monitor) monitorWebsite(w *Website) {
	ticker := time.NewTicker(w.CheckInterval * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			m.checkWebsite(w)
		case <-m.stop:
			return
		}
	}
}

// checkWebsite performs a single check of a website
func (m *Monitor) checkWebsite(w *Website) {
	start := time.Now()
	resp, err := http.Get(w.URL)
	elapsed := time.Since(start)
	
	m.mu.Lock()
	defer m.mu.Unlock()
	
	w.LastChecked = time.Now()
	w.ResponseTime = elapsed
	w.TotalChecks++
	
	if err != nil || resp.StatusCode >= 400 {
		w.Status = "Down"
	} else {
		w.Status = "Up"
		w.SuccessChecks++
		resp.Body.Close()
	}
	
	// Calculate uptime percentage
	w.Uptime = float64(w.SuccessChecks) / float64(w.TotalChecks) * 100.0
}

// TODO: Implement HTTP handlers for the REST API
// - GET /websites - List all websites
// - GET /websites/{url} - Get a specific website
// - POST /websites - Add a new website
// - DELETE /websites/{url} - Remove a website

func main() {
	// Create a new monitor
	monitor := NewMonitor()
	
	// Add a few websites to monitor
	monitor.AddWebsite(Website{
		URL:           "https://www.google.com",
		Name:          "Google",
		CheckInterval: 60,
	})
	
	monitor.AddWebsite(Website{
		URL:           "https://www.github.com",
		Name:          "GitHub",
		CheckInterval: 120,
	})
	
	// Set up API routes
	http.HandleFunc("/websites", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// Return list of all websites
			websites := monitor.GetWebsites()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(websites)
			
		case "POST":
			// Add a new website
			var website Website
			err := json.NewDecoder(r.Body).Decode(&website)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}
			
			monitor.AddWebsite(website)
			w.WriteHeader(http.StatusCreated)
			
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	
	// TODO: Add more routes and handlers
	
	// Start the server
	fmt.Println("Starting uptime monitor on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
