package params

import (
	"time"

	beaconnumeric "github.com/OffchainLabs/prysm/v6/consensus-types/primitives"
	"github.com/OffchainLabs/prysm/v6/runtime/version"
	"github.com/holiman/uint256"
	"github.com/sirupsen/logrus"
)

// Ephemery constants based on EIP-6916
const (
	// EphemeryPeriod defines the duration of each Ephemery network iteration (e.g., 7 days).
	EphemeryPeriod = 7 * 24 * 60 * 60 // 604800 seconds
	// EphemeryGenesis0Time is the timestamp of the very first Ephemery genesis (iteration i=0).
	// Example: April 9th, 2023, 00:00:00 UTC
	EphemeryGenesis0Time = 1680998400
	// EphemeryGenesis0ChainID is the base chain ID for the first Ephemery iteration (i=0).
	// Example: 6916
	EphemeryGenesis0ChainID = 6916
)

// EphemeryConfig defines the config for the Ephemery beacon chain testnet.
// NOTE: This function needs to be kept up-to-date with the official Ephemery specs.
// It defines the STATIC parameters. Dynamic parameters (genesis time, validator root)
// will be calculated and overridden during initialization.
func EphemeryConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()

	cfg.ConfigName = EphemeryName

	// TODO(dikshasethi2511): Add the parameters from the Ephemery spec.

	// Initialize the fork schedule based on the overridden parameters.
	cfg.InitializeForkSchedule()
	return cfg
}

// UseEphemeryNetworkConfig applies Ephemery-specific network parameters,
// primarily loading bootnodes from the configuration preset.
// This function should be called AFTER the main config (EphemeryConfig) is set active.
func UseEphemeryNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy() 

	// TODO(dikshasethi2511): Implement the logic to load the bootnodes from the Ephemery spec.

	// Override the global network config with the Ephemery specifics (mainly bootnodes).
	OverrideBeaconNetworkConfig(cfg)
}