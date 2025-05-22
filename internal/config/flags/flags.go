package flags

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/devlongs/beamlight/internal/config"
)

// Flags represents the command-line flags.
type Flags struct {
	// ConfigFile is the path to the config file.
	ConfigFile string

	// DataDir is the directory to store data.
	DataDir string

	// NetworkFlag is the network to connect to (e.g. "mainnet", "testnet").
	NetworkFlag string

	// LogLevel is the log level.
	LogLevel string

	// LogFile is the file to write logs to.
	LogFile string

	// MetricsFlag enables or disables metrics collection.
	MetricsFlag bool

	// MetricsAddr is the address to expose metrics.
	MetricsAddr string

	// Help shows the help message.
	Help bool

	// Version shows the version.
	Version bool
}

// Parse parses the command-line flags.
func Parse() *Flags {
	flags := &Flags{}

	defaultConfig := config.DefaultConfig()

	flag.StringVar(&flags.ConfigFile, "config", "", "Path to the config file")
	flag.StringVar(&flags.DataDir, "datadir", defaultConfig.DataDir, "Data directory for the database and keystore")
	flag.StringVar(&flags.NetworkFlag, "network", "mainnet", "Network to connect to (mainnet, testnet, devnet)")
	flag.StringVar(&flags.LogLevel, "log.level", defaultConfig.LogConfig.Level, "Log level (debug, info, warn, error)")
	flag.StringVar(&flags.LogFile, "log.file", defaultConfig.LogConfig.File, "Log file path (default: stderr)")
	flag.BoolVar(&flags.MetricsFlag, "metrics", defaultConfig.EnableMetrics, "Enable metrics collection")
	flag.StringVar(&flags.MetricsAddr, "metrics.addr", defaultConfig.MetricsAddress, "Metrics HTTP server address")
	flag.BoolVar(&flags.Help, "help", false, "Show help")
	flag.BoolVar(&flags.Version, "version", false, "Show version")

	// Add short flags
	flag.BoolVar(&flags.Help, "h", false, "Show help (shorthand)")
	flag.BoolVar(&flags.Version, "v", false, "Show version (shorthand)")

	flag.Parse()

	return flags
}

// ApplyToConfig applies the flags to the config.
func (f *Flags) ApplyToConfig(cfg *config.Config) {
	if f.DataDir != "" {
		cfg.DataDir = f.DataDir
	}

	if f.LogLevel != "" {
		cfg.LogConfig.Level = f.LogLevel
	}

	if f.LogFile != "" {
		cfg.LogConfig.File = f.LogFile
	}

	if f.MetricsAddr != "" {
		cfg.MetricsAddress = f.MetricsAddr
	}

	cfg.EnableMetrics = f.MetricsFlag
}

// PrintHelp prints the help message.
func (f *Flags) PrintHelp() {
	fmt.Println("Usage of beamlight:")
	fmt.Println()
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  beamlight --datadir ~/.beamlight --log.level debug")
}

// PrintVersion prints the version.
func (f *Flags) PrintVersion(version, commit, date string) {
	fmt.Printf("Beamlight v%s\n", version)
	fmt.Printf("Git Commit: %s\n", commit)
	fmt.Printf("Build Date: %s\n", date)
}

// ValidateFlags validates the flags.
func (f *Flags) ValidateFlags() error {
	// Check if the data directory exists and is writable
	if f.DataDir != "" {
		if err := os.MkdirAll(f.DataDir, 0755); err != nil {
			return fmt.Errorf("failed to create data directory: %w", err)
		}

		// Try to create a temporary file to check if the directory is writable
		testFile := filepath.Join(f.DataDir, ".write_test")
		if err := os.WriteFile(testFile, []byte{}, 0644); err != nil {
			return fmt.Errorf("data directory is not writable: %w", err)
		}
		os.Remove(testFile)
	}

	// Validate log level
	validLevels := []string{"debug", "info", "warn", "error"}
	if f.LogLevel != "" {
		valid := false
		for _, level := range validLevels {
			if strings.EqualFold(f.LogLevel, level) {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("invalid log level: %s (valid levels: %s)", f.LogLevel, strings.Join(validLevels, ", "))
		}
	}

	// Validate network
	validNetworks := []string{"mainnet", "testnet", "devnet"}
	if f.NetworkFlag != "" {
		valid := false
		for _, network := range validNetworks {
			if strings.EqualFold(f.NetworkFlag, network) {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("invalid network: %s (valid networks: %s)", f.NetworkFlag, strings.Join(validNetworks, ", "))
		}
	}

	return nil
}
