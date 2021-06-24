package fisco

import (
	"encoding/json"
	"strconv"
	"strings"
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
	ChainID            string  `json:"chain_id"`
	IServiceCoreAddr   string `json:"iservice_core_addr"`
}

type EndpointInfo struct {
	DestChainID string `json:"dest_chain_id"`
	EndpointAddress string `json:"endpoint_address"`
	EndpointType string `json:"endpoint_type"`
}

// GetChainID returns the unique fisco chain id from the ChainID
func GetFiscoChainID(ChainID string) int64 {
	chainInfos := strings.Split(ChainID, "-")
	fiscoChainID, _ := strconv.ParseInt(chainInfos[2], 10, 64)
	return fiscoChainID
}

// GetGroupID returns the unique fisco group id from the ChainID
func GetFiscoGroupID(ChainID string) int  {
	chainInfos := strings.Split(ChainID, "-")
	fiscoGroupID, _ := strconv.Atoi(chainInfos[1])
	return fiscoGroupID
}

// GetChainIDFromBytes returns the unique chain id from the given chain params bytes
func GetChainIDFromBytes(params []byte) (string, error) {
	var chainParams ChainParams
	err := json.Unmarshal(params, &chainParams)
	if err != nil {
		return "", err
	}

	return chainParams.ChainID, nil
}
