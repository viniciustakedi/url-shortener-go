package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"urlshortener/config"
	"urlshortener/cron"
	"urlshortener/infra/db"
	"urlshortener/server"
)

func main() {
	// TO-DO:
	// 1. Create sign-up and sign-in system.
	// 2. System to generate QRCode and track link access.

	environment := flag.String("e", "development", "Environment to run the application in (development, staging, production)")

	flag.Usage = func() {
		log.Fatalf(
			"Usage: %s [options]\nOptions:\n  -env string\n\tEnvironment to run the application in (development, staging, production)",
			flag.CommandLine.Name(),
		)

		os.Exit(1)
	}

	flag.Parse()

	// Initialize the API config & Envs with flag
	config.Init(*environment)

	// Initialize Database Connection - MongoDB
	db.InitMongoDB()

	// Initialize the HTTP server to serve the API
	httpServer := server.Init(*environment)

	// Initialize the cron jobs
	// This will run the cron jobs in a separate goroutine
	// and will not block the main thread.
	cron.Init()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\nShutting down server...")
	server.Shutdown(httpServer)
}
