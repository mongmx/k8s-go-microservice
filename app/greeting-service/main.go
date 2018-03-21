package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"github.com/mongmx/mymicro/app/handlers"
)

func main() {
	var (
		httpAddr   = flag.String("http", "0.0.0.0:8080", "HTTP service address.")
		healthAddr = flag.String("health", "0.0.0.0:8081", "Health service address.")
	)
	flag.Parse()

	log.Println("Starting server...")
	log.Printf("Health service listening on %s", *healthAddr)
	// log.Printf("HTTP service listening on %s", *httpAddr)

	errChan := make(chan error, 10)

	hmux := http.NewServeMux()
	hmux.HandleFunc("/healthz", handlers.HealthzHandler)
	hmux.HandleFunc("/readiness", handlers.ReadinessHandler)
	hmux.HandleFunc("/healthz/status", handlers.HealthzStatusHandler)
	hmux.HandleFunc("/readiness/status", handlers.ReadinessStatusHandler)
	healthServer := manners.NewServer()
	healthServer.Addr = *healthAddr
	healthServer.Handler = handlers.LoggingHandler(hmux)

	go func() {
		errChan <- healthServer.ListenAndServe()
	}()

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/greeting", greetingEndpoint)
	}
	router.Run(*httpAddr)
}

func greetingEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}
