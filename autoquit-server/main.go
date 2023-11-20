package main

// test with curl -X POST http://localhost:8080/quit
import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	wg       sync.WaitGroup
	shutdown = make(chan struct{})
)

func main() {
	// Create a new server and handle routes
	server := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(handleRequest)}

	// Use a goroutine to run the server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	// Set up a signal channel to capture signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		fmt.Printf("Received signal: %v\n", sig)
	case <-shutdown:
		fmt.Println("Server shutdown initiated")
	}

	// Shut down the server gracefully
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Printf("Error during shutdown: %v\n", err)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Server gracefully stopped")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/quit" && r.Method == http.MethodPost {
		close(shutdown) // Signal to shut down the server
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Server will shut down soon")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
	}
}

