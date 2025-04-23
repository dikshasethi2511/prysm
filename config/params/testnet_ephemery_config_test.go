package params_test

import (
	"testing"

	"github.com/OffchainLabs/prysm/v6/config/params"
	"github.com/stretchr/testify/require"
)

// TestEphemeryConfigValues validates the parameters set in the EphemeryConfig function.
// NOTE: This test needs to be updated with the correct expected values once
// EphemeryConfig() is fully implemented based on the official specification.
func TestEphemeryConfigValues(t *testing.T) {
	cfg := params.EphemeryConfig()

	// Basic Name Check.
	t.Run("ConfigName", func(t *testing.T) {
		require.Equal(t, params.EphemeryName, cfg.ConfigName, "ConfigName should match EphemeryName.")
	})
} 