package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/qwerty0981/gin-template/routes"
)

func main() {
	router := routes.GetRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Run the server on a separate 'thread'
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server Error: %s\n", err)
		}
	}()

	// Make a channel to listen for SIGINT and SIGTERM
	quit := make(chan os.Signal)

	// Register channel to interupts
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// wait for quit
	<-quit

	fmt.Println("Shutting down the server...")

	// Wait for 5 seconds to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Cancels the shutdown if 5 seconds have passed
	defer cancel()

	// Tells the server to shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// This line will only be reached if the server successfuly shuts down
	log.Println("Server gracefully shutdown")
}
