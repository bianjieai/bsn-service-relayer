package server

import (
	"relayer/core"
)

// ChainManager defines a service for app chains management
type ChainManager struct {
	relayer *core.Relayer
}

// NewChainManager constructs a new ChainManager instance
func NewChainManager(r *core.Relayer) *ChainManager {
	return &ChainManager{
		relayer: r,
	}
}

// AddChain adds a new app chain for the relayer
func (cm *ChainManager) AddChain(params []byte) (chainID string, err error) {
	return cm.relayer.AddChain(params)
}

// DeleteChain delete a app chain for the relayer
func (cm *ChainManager) DeleteChain(chainID string) error {
	return cm.relayer.DeleteChain(chainID)
}
// StartChain starts to relay an existent app chain
func (cm *ChainManager) StartChain(chainID string) error {
	return cm.relayer.StartChain(chainID)
}

// StopChain stops to relay an app chain
func (cm *ChainManager) StopChain(chainID string) error {
	return cm.relayer.StopChain(chainID)
}

// GetChains gets all active app chains
func (cm *ChainManager) GetChains() []string {
	return cm.relayer.GetChains()
}

// GetChainStatus retrieves the status of the specified app chain
func (cm *ChainManager) GetChainStatus(chainID string) (state bool, height int64, err error) {
	return cm.relayer.GetChainStatus(chainID)
}

// MarketManager defines a service for the iService market management on app chains
type MarketManager struct {
	relayer *core.Relayer
}

// NewMarketManager constructs a new MarketManager instance
func NewMarketManager(r *core.Relayer) *MarketManager {
	return &MarketManager{
		relayer: r,
	}
}

// AddServiceBinding adds the given service binding on the specified app chain
func (mm *MarketManager) AddServiceBinding(
	chainID string,
	req AddServiceBindingRequest,
) error {
	appChain, err := mm.relayer.GetChain(chainID)
	if err != nil {
		return err
	}

	return appChain.AddServiceBinding(
		req.ServiceName,
		req.Schemas,
		req.Provider,
		req.ServiceFee,
		req.QoS,
	)
}

// UpdateServiceBinding updates the given service binding on the specified app chain
func (mm *MarketManager) UpdateServiceBinding(
	chainID string,
	serviceName string,
	req UpdateServiceBindingRequest,
) error {
	appChain, err := mm.relayer.GetChain(chainID)
	if err != nil {
		return err
	}

	return appChain.UpdateServiceBinding(
		serviceName,
		req.Provider,
		req.ServiceFee,
		req.QoS,
	)
}

// GetServiceBinding gets the given service binding on the specified app chain
func (mm *MarketManager) GetServiceBinding(
	chainID string,
	serviceName string,
) (interface{}, error) {
	appChain, err := mm.relayer.GetChain(chainID)
	if err != nil {
		return nil, err
	}

	return appChain.GetServiceBinding(serviceName)
}
