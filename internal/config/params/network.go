package params

import (
	"time"
)

// Network represents a network configuration.
type Network string

const (
	// MainNet is the main Ethereum network.
	MainNet Network = "mainnet"

	// TestNet is the test network.
	TestNet Network = "testnet"

	// DevNet is the development network.
	DevNet Network = "devnet"
)

// NetworkConfig represents the configuration for a network.
type NetworkConfig struct {
	// Network is the network name.
	Network Network

	// GenesisTime is the genesis time of the network.
	GenesisTime time.Time

	// EpochLength is the number of slots in an epoch.
	EpochLength uint64

	// SlotTime is the duration of a slot.
	SlotTime time.Duration

	// CommitteeSize is the size of a committee.
	CommitteeSize uint64

	// ValidatorDepositAmount is the amount required to become a validator.
	ValidatorDepositAmount uint64

	// MinValidatorBalance is the minimum balance to be an active validator.
	MinValidatorBalance uint64

	// ForkEpochs contains the epoch numbers at which forks occur.
	ForkEpochs map[string]uint64

	// BootstrapNodes is a list of bootstrap nodes.
	BootstrapNodes []string
}

// MainNetConfig returns the configuration for the main network.
func MainNetConfig() *NetworkConfig {
	return &NetworkConfig{
		Network:                MainNet,
		GenesisTime:            time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), // Placeholder
		EpochLength:            32,
		SlotTime:               12 * time.Second,
		CommitteeSize:          128,
		ValidatorDepositAmount: 32_000_000_000, // 32 ETH in Gwei
		MinValidatorBalance:    16_000_000_000, // 16 ETH in Gwei
		ForkEpochs: map[string]uint64{
			"beam": 0, // From genesis
		},
		BootstrapNodes: []string{
			// These are placeholders - real bootstrap nodes will be added later
			"enr:-PLACEHOLDER1...",
			"enr:-PLACEHOLDER2...",
		},
	}
}

// TestNetConfig returns the configuration for the test network.
func TestNetConfig() *NetworkConfig {
	return &NetworkConfig{
		Network:                TestNet,
		GenesisTime:            time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC), // Placeholder
		EpochLength:            32,
		SlotTime:               6 * time.Second, // Faster for testing
		CommitteeSize:          64,              // Smaller for testing
		ValidatorDepositAmount: 32_000_000_000,  // 32 ETH in Gwei
		MinValidatorBalance:    16_000_000_000,  // 16 ETH in Gwei
		ForkEpochs: map[string]uint64{
			"beam": 0, // From genesis
		},
		BootstrapNodes: []string{
			// These are placeholders - real bootstrap nodes will be added later
			"enr:-TESTPLACEHOLDER1...",
			"enr:-TESTPLACEHOLDER2...",
		},
	}
}

// DevNetConfig returns the configuration for the development network.
func DevNetConfig() *NetworkConfig {
	return &NetworkConfig{
		Network:                DevNet,
		GenesisTime:            time.Now().Add(-time.Minute), // Start in the past to begin immediately
		EpochLength:            8,                            // Smaller for faster development cycles
		SlotTime:               2 * time.Second,              // Faster for development
		CommitteeSize:          4,                            // Minimal for development
		ValidatorDepositAmount: 1_000_000_000,                // 1 ETH in Gwei for development
		MinValidatorBalance:    1_000_000_000,                // 1 ETH in Gwei for development
		ForkEpochs: map[string]uint64{
			"beam": 0, // From genesis
		},
		BootstrapNodes: []string{
			// No bootstrap nodes for development network - typically run locally
		},
	}
}

// GetNetworkConfig returns the configuration for the specified network.
func GetNetworkConfig(network Network) *NetworkConfig {
	switch network {
	case MainNet:
		return MainNetConfig()
	case TestNet:
		return TestNetConfig()
	case DevNet:
		return DevNetConfig()
	default:
		return DevNetConfig() // Default to development network
	}
}
