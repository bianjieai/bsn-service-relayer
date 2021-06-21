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

