package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/devlongs/beamlight/internal/config"
	"github.com/devlongs/beamlight/internal/config/flags"
	"github.com/devlongs/beamlight/internal/node"
)

// Version information of the beamlight client.
var (
	Version   = "0.1.0"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

func main() {
	// Parse command-line flags
	cmdFlags := flags.Parse()

	// Check if help flag is set
	if cmdFlags.Help {
		cmdFlags.PrintHelp()
		os.Exit(0)
	}

	// Check if version flag is set
	if cmdFlags.Version {
		cmdFlags.PrintVersion(Version, GitCommit, BuildDate)
		os.Exit(0)
	}

	// Validate flags
	if err := cmdFlags.ValidateFlags(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.LoadConfig(cmdFlags.ConfigFile)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Apply flags to configuration
	cmdFlags.ApplyToConfig(cfg)

	fmt.Println("Beamlight Ethereum Beam Chain Client")
	fmt.Printf("Version: %s (commit: %s, built: %s)\n", Version, GitCommit, BuildDate)
	fmt.Printf("Data directory: %s\n", cfg.DataDir)

	// Create a new node instance
	beamlightNode := node.New()

	// TODO: Register services based on configuration
	// beamlightNode.RegisterService(...)

	// Start the node
	if err := beamlightNode.Start(); err != nil {
		fmt.Printf("Failed to start node: %v\n", err)
		os.Exit(1)
	}

	// Wait for termination signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("Received shutdown signal, gracefully shutting down...")

	// Stop the node
	if err := beamlightNode.Stop(); err != nil {
		fmt.Printf("Error during shutdown: %v\n", err)
	}

	fmt.Println("Shutdown complete")
}
