package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents the configuration for the beamlight client.
type Config struct {
	// DataDir is the directory to store data.
	DataDir string `json:"data_dir"`

	// Network is the network configuration.
	Network NetworkConfig `json:"network"`

	// Database is the database configuration.
	Database DatabaseConfig `json:"database"`

	// LogConfig is the logging configuration.
	LogConfig LogConfig `json:"log"`

	// EnableMetrics enables the metrics collection.
	EnableMetrics bool `json:"enable_metrics"`

	// MetricsAddress is the address to expose metrics.
	MetricsAddress string `json:"metrics_address"`
}

// NetworkConfig is the configuration for the network.
type NetworkConfig struct {
	// ListenAddress is the address to listen for incoming connections.
	ListenAddress string `json:"listen_address"`

	// BootstrapNodes is a list of bootstrap nodes to connect to.
	BootstrapNodes []string `json:"bootstrap_nodes"`

	// MaxPeers is the maximum number of peers to connect to.
	MaxPeers int `json:"max_peers"`

	// EnableDiscovery enables the peer discovery.
	EnableDiscovery bool `json:"enable_discovery"`
}

// DatabaseConfig is the configuration for the database.
type DatabaseConfig struct {
	// Engine is the database engine to use (e.g. "leveldb", "badger").
	Engine string `json:"engine"`

	// CacheSize is the cache size for the database.
	CacheSize int `json:"cache_size"`
}

// LogConfig is the configuration for logging.
type LogConfig struct {
	// Level is the log level (e.g. "debug", "info", "warn", "error").
	Level string `json:"level"`

	// File is the file to write logs to. If empty, logs are written to stderr.
	File string `json:"file"`

	// Format is the log format (e.g. "text", "json").
	Format string `json:"format"`
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	return &Config{
		DataDir:        filepath.Join(homeDir(), ".beamlight"),
		EnableMetrics:  true,
		MetricsAddress: "127.0.0.1:8080",
		Network: NetworkConfig{
			ListenAddress:   "0.0.0.0:30303",
			BootstrapNodes:  []string{},
			MaxPeers:        50,
			EnableDiscovery: true,
		},
		Database: DatabaseConfig{
			Engine:    "leveldb",
			CacheSize: 256,
		},
		LogConfig: LogConfig{
			Level:  "info",
			File:   "",
			Format: "text",
		},
	}
}

// LoadConfig loads the configuration from a file.
func LoadConfig(file string) (*Config, error) {
	config := DefaultConfig()

	if file == "" {
		return config, nil
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return config, nil
}

// SaveConfig saves the configuration to a file.
func SaveConfig(config *Config, file string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize config: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(file, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// homeDir returns the user's home directory.
func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if home := os.Getenv("USERPROFILE"); home != "" {
		return home
	}
	return os.TempDir()
}
