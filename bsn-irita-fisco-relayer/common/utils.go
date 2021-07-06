package common

import (
	"fmt"
	"math/big"
	"os"
)

// MustGetHomeDir gets the user home directory
// Panic if an error occurs
func MustGetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return homeDir
}

// Hex2Decimal converts the given hex string to a decimal number
func Hex2Decimal(hex string) (int64, error) {
	i := new(big.Int)

	i, ok := i.SetString(hex, 16)
	if !ok {
		return -1, fmt.Errorf("Cannot parse hex string to Int")
	}

	return i.Int64(), nil
}

// GetChainID returns the unique chain id from the specified chain params
func GetDestID(chainType string, groupID string, chainID string) string {
	return fmt.Sprintf("%s-%d-%d", chainType, groupID, chainID)
}
