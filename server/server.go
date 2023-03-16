package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	DEFAULT_PORT             = 3001
	DEFAULT_SHUTDOWN_TIMEOUT = 5 * time.Second
)

var signals = []os.Signal{
	os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM,
}

func main() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", getPort()),
		Handler: makeHandler(),
	}
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, signals...)
	serverChannel := make(chan error)
	go runServer(server, serverChannel)
	fmt.Printf("Running: %s\n", server.Addr)

	select {
	case err := <-serverChannel:
		fmt.Printf("Error: %s\n", err)
	case <-signalChannel:
		fmt.Printf("Shutting down...\n")
		if err := shutdownServer(server); err != nil {
			fmt.Printf("Failed to shutdown: %s\n", err)
		}
	}
}

func makeHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/calculate", handleCalculate)
	return mux
}

func runServer(server *http.Server, ch chan<- error) {
	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		err = nil
	}
	ch <- err
}

func shutdownServer(server *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), DEFAULT_SHUTDOWN_TIMEOUT)
	defer cancel()
	return server.Shutdown(ctx)
}

func getPort() int {
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		return port
	}
	return DEFAULT_PORT
}
