package fisco

import (
	"encoding/json"
	"fmt"
)

const (
	ChainType              = "fisco"
	DefaultMonitorInterval = 1 // 1 second by default
)

// CompactBlock represents the compact block with tx hashes
type CompactBlock struct {
	Txs []string `json:"transactions"`
}

// ChainParams defines the params for the specific chain
type ChainParams struct {
	NodeURLs           map[string]string `json:"node_urls"`
	GroupID            int    `json:"group_id"`
	ChainID            int64  `json:"chain_id"`
	IServiceCoreAddr   string `json:"iservice_core_addr"`
}

type EndpointInfo struct {
	DestSubChainID string `json:"dest_sub_chain_id"`
	DestChainID string `json:"dest_chain_id"`
	DestChainType string `json:"dest_chain_type"`
	EndpointAddress string `json:"endpoint_address"`
	EndpointType string `json:"endpoint_type"`
}

// GetChainID returns the unique chain id from the specified chain params
func GetChainID(params ChainParams) string {
	return fmt.Sprintf("%s-%d-%d", ChainType, params.GroupID, params.ChainID)
}

// GetChainIDFromBytes returns the unique chain id from the given chain params bytes
func GetChainIDFromBytes(params []byte) (string, error) {
	var chainParams ChainParams
	err := json.Unmarshal(params, &chainParams)
	if err != nil {
		return "", err
	}

	return GetChainID(chainParams), nil
}