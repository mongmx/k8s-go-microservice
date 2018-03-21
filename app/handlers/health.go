package handlers

import (
	"net/http"
	"sync"
)

var (
	healthzStatus   = http.StatusOK
	readinessStatus = http.StatusOK
	mu              sync.RWMutex
)

// HealthzStatus func
func HealthzStatus() int {
	mu.RLock()
	defer mu.RUnlock()
	return healthzStatus
}

// ReadinessStatus func
func ReadinessStatus() int {
	mu.RLock()
	defer mu.RUnlock()
	return readinessStatus
}

// SetHealthzStatus func
func SetHealthzStatus(status int) {
	mu.Lock()
	healthzStatus = status
	mu.Unlock()
}

// SetReadinessStatus func
func SetReadinessStatus(status int) {
	mu.Lock()
	readinessStatus = status
	mu.Unlock()
}

// HealthzHandler responds to health check requests.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(HealthzStatus())
}

// ReadinessHandler responds to readiness check requests.
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(ReadinessStatus())
}

// ReadinessStatusHandler func
func ReadinessStatusHandler(w http.ResponseWriter, r *http.Request) {
	switch ReadinessStatus() {
	case http.StatusOK:
		SetReadinessStatus(http.StatusServiceUnavailable)
	case http.StatusServiceUnavailable:
		SetReadinessStatus(http.StatusOK)
	}
	w.WriteHeader(http.StatusOK)
}

// HealthzStatusHandler func
func HealthzStatusHandler(w http.ResponseWriter, r *http.Request) {
	switch HealthzStatus() {
	case http.StatusOK:
		SetHealthzStatus(http.StatusServiceUnavailable)
	case http.StatusServiceUnavailable:
		SetHealthzStatus(http.StatusOK)
	}
	w.WriteHeader(http.StatusOK)
}
