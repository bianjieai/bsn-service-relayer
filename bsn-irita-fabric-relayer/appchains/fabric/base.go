package fabric

import (
	"encoding/json"
	"time"
	"fmt"
)

const (
	ChainType              = "fabric"
	DefaultMonitorInterval = 1 // 1 second by default
)

// CompactBlock represents the compact block with tx hashes
type CompactBlock struct {
	Txs []string `json:"transactions"`
}

// ChainParams defines the params for the specific chain
type ChainParams struct {
	ChainId uint64 `json:"chainId"`

	AppCode string `json:"appCode"`

	CrossChainCode string `json:"ccm"`

	ChannelId string `json:"channelId"`

	Nodes []string `json:"nodes"`
}

func (r *ChainParams) ToStoreData(cityCode string) *FabricRelayer {

	relayer := &FabricRelayer{
		ID: r.ChainId,
		AppCode:   r.AppCode,
		ChannelId: r.ChannelId,

		CrossChainCode: r.CrossChainCode,
		CityNode:       cityCode,
		Status:         1,
		CreateTime:     time.Now(),
	}
	relayer.SetNodes(r.Nodes)
	return relayer
}

type EndpointInfo struct {
	DestChainID string `json:"dest_chain_id"`
	EndpointAddress string `json:"endpoint_address"`
	EndpointType string `json:"endpoint_type"`
}


// GetChainID returns the unique chain id from the specified chain params
func GetChainID(params ChainParams) string {
	return fmt.Sprintf("%s-%d-%d", ChainType, params.ChannelId, params.ChainId)
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
