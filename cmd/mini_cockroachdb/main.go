// Here's the starting point.

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Abhisheklearn12/miniCockroachDB_in_Go/node"
)

func main() {
	log.Println("[mini-cockroachdb] Starting node...")

	// Node configuration (replace with flag/env parsing in production)
	cfg := node.Config{
		NodeID:      "node1",
		GRPCAddress: ":9000",
		DataDir:     "./data/node1",
		Peers:       []string{":9001", ":9002"}, // Example peer addresses
	}

	// Setup context for lifecycle management
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Signal handling for graceful shutdown
	go handleShutdown(cancel)

	// Initialize and start the node
	n, err := node.NewNode(cfg)
	if err != nil {
		log.Fatalf("Error initializing node: %v", err)
	}

	if err := n.Start(ctx); err != nil {
		log.Fatalf("Error starting node: %v", err)
	}

	// Wait until a termination signal is received
	<-ctx.Done()
	log.Println("[mini-cockroachdb] Shutting down node...")

	if err := n.Stop(); err != nil {
		log.Fatalf("Error during shutdown: %v", err)
	}

	log.Println("[mini-cockroachdb] Node shutdown complete.")
}

func handleShutdown(cancelFunc context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Printf("Termination signal received: %s", sig)
	cancelFunc()
}
