package fisco

import (
	"encoding/json"
	"strconv"
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
	NodeURLs           []string `json:"nodes"`
	GroupID            int      `json:"groupId"`
	ChainID            int64    `json:"chainId"`
	IServiceCoreAddr   string   `json:"iserviceCoreAddr"`
	IServiceMarketAddr string   `json:"iserviceMarketAddr"`
}

// GetChainID returns the unique chain id from the specified chain params
func GetChainID(params ChainParams) string {
	return strconv.FormatInt(params.ChainID, 10)

	//	return fmt.Sprintf("%s-%d-%d", ChainType, params.GroupID, params.ChainID)
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
