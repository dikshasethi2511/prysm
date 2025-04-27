package params

import (
	"log"
	"time"
)

// CalculateEphemeryGenesisTime determines the genesis timestamp for the current Ephemery week
// based on the EIP-6916 formula: `period * i + genesis_0.genesis_time` where
// `i = floor((current_timestamp - genesis_0.genesis_time) / period)`.
func CalculateEphemeryGenesisTime(currentTime time.Time) uint64 {
	currentTimestamp := uint64(currentTime.Unix())

	if currentTimestamp < EphemeryGenesis0Time {
		// If current time is before the absolute start, return the start time itself.
		// This might happen due to clock skew or if called before the official start.
		log.Warnf("Current time %d is before Ephemery Genesis 0 time %d. Using Genesis 0 time.", currentTimestamp, EphemeryGenesis0Time)
		return EphemeryGenesis0Time
	}

	timeElapsed := currentTimestamp - EphemeryGenesis0Time
	iteration := timeElapsed / EphemeryPeriod // Integer division performs the floor operation

	currentGenesisTime := (iteration * EphemeryPeriod) + EphemeryGenesis0Time
	return currentGenesisTime
} 