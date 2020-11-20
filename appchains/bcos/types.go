package bcos

import (
	"fmt"
)

const (
	ChainIDPrefix          = "bcos-"
	DefaultMonitorInterval = 1 // 1 second by default

	StorePrefix = "bcos:"
	KeyHeight   = "height"
)

// CompactBlock represents the compact block with tx hashes
type CompactBlock struct {
	Txs []string `json:"transactions"`
}

// GetHeightKey returns the key for the BCOS height store
func GetHeightKey() []byte {
	return []byte(fmt.Sprintf("%s%s", StorePrefix, KeyHeight))
}
