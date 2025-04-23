package params

import (
	"github.com/sirupsen/logrus"
)

// EphemeryConfig defines the config for the Ephemery beacon chain testnet.
// NOTE: This function needs to be kept up-to-date with the official Ephemery specs.
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