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
	cfg.PresetBase = "mainnet" // Assuming Ephemery follows the mainnet preset for base parameters such as epoch lengths, slot times, etc.

	// Forks (Enable all known forks at genesis).
	cfg.AltairForkEpoch = 0
	cfg.BellatrixForkEpoch = 0
	cfg.CapellaForkEpoch = 0
	cfg.DenebForkEpoch = 0
	cfg.ElectraForkEpoch = 0
	cfg.FuluForkEpoch = beaconnumeric.FarFutureEpoch // Disable Fulu unless needed.

	// Merge Transition (Effectively occurs at genesis).
	cfg.TerminalTotalDifficulty = uint256.NewInt(0)
	cfg.TerminalBlockHash = [32]byte{}
	cfg.TerminalBlockHashActivationEpoch = 0

	// Eth1 Configuration..
	// Use the common testnet deposit contract address.
	cfg.DepositContractAddress = "0x4242424242424242424242424242424242424242"

	// Genesis.
	cfg.GenesisDelay = 604800 // 1 week
	// MIN_GENESIS_TIME will be set dynamically based on calculatedEphemeryGenesisTime
	// GENESIS_TIME will be set dynamically
	// GENESIS_VALIDATORS_ROOT will be set dynamically
	cfg.MinGenesisActiveValidatorCount = 128 // A lower count suitable for testnets

	// Fork Versions (Use distinct versions for Ephemery, e.g., 0x6...69)
	cfg.GenesisForkVersion = [4]byte{0x60, 0x00, 0x00, 0x69}
	cfg.AltairForkVersion = [4]byte{0x61, 0x00, 0x00, 0x69}
	cfg.BellatrixForkVersion = [4]byte{0x62, 0x00, 0x00, 0x69}
	cfg.CapellaForkVersion = [4]byte{0x63, 0x00, 0x00, 0x69}
	cfg.DenebForkVersion = [4]byte{0x64, 0x00, 0x00, 0x69}
	cfg.ElectraForkVersion = [4]byte{0x65, 0x00, 0x00, 0x69} // Adjust if needed
	cfg.FuluForkVersion = [4]byte{0x66, 0x00, 0x00, 0x69}     // Adjust if needed

	// Network ID / Chain ID (These might be set dynamically based on genesis time in EIP-6916)
	// Let's leave them as Mainnet defaults for now, and adjust later when implementing
	// the dynamic loading logic if EIP-6916 requires overriding them.
	// cfg.DepositChainID = 1 (Mainnet default) - This will be set dynamically by features/config.go
	// cfg.DepositNetworkID = 1 (Mainnet default)

	// Ensure Prysm knows this is NOT mainnet for certain checks
	if version.IsRelease() {
		cfg.IsPreRelease = true
	}

	// Initialize the fork schedule based on the overridden parameters.
	cfg.InitializeForkSchedule()
	return cfg
}

// UseEphemeryNetworkConfig applies Ephemery-specific network parameters,
// primarily loading bootnodes from the configuration preset.
// This function should be called AFTER the main config (EphemeryConfig) is set active.
func UseEphemeryNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()

	cfg.BootstrapNodes = []string{
		"enr:-LK4QL7PXJSGyl4T_iKkerq07T_BdsxsZ2LlTBohq2s23w0JINqG5nsMDQZY1RAAKnFmD4n4l4tGdxl-z9A2amRBz38Eh2F0dG5ldHOIAAAAAABgAACEZXRoMpA8nXi9YAAQG___________gmlkgnY0gmlwhIlKy_CJc2VjcDI1NmsxoQObUJVAhjz_1bIDScKyC3rUbNDbX88jW3B-419NQ5jgmYN0Y3CCI4yDdWRwgiOM",
		"enr:-Iq4QNMYHuJGbnXyBj6FPS2UkOQ-hnxT-mIdNMMr7evR9UYtLemaluorL6J10RoUG1V4iTPTEbl3huijSNs5_ssBWFiGAYhBNHOzgmlkgnY0gmlwhIlKy_CJc2VjcDI1NmsxoQNULnJBzD8Sakd9EufSXhM4rQTIkhKBBTmWVJUtLCp8KoN1ZHCCIyk",
		"enr:-Jq4QLBK_62odVx0uL6eki9fBBt3rhzJUwS9SbKdLOOtOpzLCQo3rvSMVRF0yB2C-a0tZbl5hA6qnymbJvrHPZVfMx4BhGV0aDKQPJ14vWAAEBv__________4JpZIJ2NIJpcIRBbZouiXNlY3AyNTZrMaECEuTafsVntm3n88tfJrviveX1xravS9wHTMOEHEaUItKDdWRwgiMp",
		"enr:-Iq4QIc297-de1P6hznMX2cIdVsQkve9BD9NUsJ7vVQa7eh5UpekA9rLid5A-yLiS3gZwOGugYZPi58x76zNs2cEQFCGAYhBJlTYgmlkgnY0gmlwhEFtmi6Jc2VjcDI1NmsxoQJDyix-IHa_mVwLBEN9NeG8I-RUjNQK_MGxk9OqRQUAtIN1ZHCCIyg",
		"enr:-OS4QETSHRuGhmGM7w9wMTpu1--rAmI6T4dQS-95A-r677RuWqAeznXxuuW9IeGFh0xLzHQZuj5xoKmUy93yM26Z53UDh2F0dG5ldHOIAAAAGAAAAACGY2xpZW502IpMaWdodGhvdXNljDcuMC4wLWJldGEuN4RldGgykDydeL1gABAb__________-CaWSCdjSCaXCEp-sBuYRxdWljgiNRiXNlY3AyNTZrMaEC8Ho52PHI4omvO5wit5mI-lcQ68muWWGZ_6yibAODRk2Ic3luY25ldHMAg3RjcIIjUIN1ZHCCI1A",
	}

	// Override the global network config with the Ephemery specifics (mainly bootnodes).
	OverrideBeaconNetworkConfig(cfg)
}