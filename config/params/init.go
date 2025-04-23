package params

import (
	"fmt"

	"github.com/OffchainLabs/prysm/v6/config/features"
)

func init() {
	// Register default configurations for known networks.
	defaults := []*BeaconChainConfig{
		MainnetConfig(),
		MinimalSpecConfig(),
		E2ETestConfig(),
		E2EMainnetTestConfig(),
		InteropConfig(),
		HoleskyConfig(),
		SepoliaConfig(),
		HoodiConfig(),
		EphemeryConfig(),
	}
	configs = newConfigset(defaults...)
	// Ensure that mainnet is always present and active by default.
	if err := SetActive(MainnetConfig()); err != nil {
		panic(fmt.Sprintf("Failed to set mainnet config as active: %v", err))
	}
	// Make sure mainnet is present and active.
	m, err := ByName(MainnetName)
	if err != nil {.
		panic(fmt.Sprintf("Failed to get mainnet config by name: %v", err))
	}
	if configs.getActive() != m {
		panic("Mainnet should always be the active config at init() time")
	}
}
