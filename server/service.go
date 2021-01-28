package server

import (
	"relayer/core"
)

// ChainManager defines a service for app chains management
type ChainManager struct {
	Relayer *core.Relayer
}

// NewChainManager constructs a new ChainManager instance
func NewChainManager(r *core.Relayer) *ChainManager {
	return &ChainManager{
		Relayer: r,
	}
}

// AddChain adds a new app chain for the relayer
func (cm *ChainManager) AddChain(params []byte) (chainID string, err error) {
	return cm.Relayer.AddChain(params)
}

// StartChain starts to relay an existent app chain
func (cm *ChainManager) StartChain(chainID string) error {
	return cm.Relayer.StartChain(chainID)
}

// StopChain stops to relay an app chain
func (cm *ChainManager) StopChain(chainID string) error {
	return cm.Relayer.StopChain(chainID)
}

// GetChains gets all active app chains
func (cm *ChainManager) GetChains() []string {
	return cm.Relayer.GetChains()
}

// GetChainStatus retrieves the status of the specified app chain
func (cm *ChainManager) GetChainStatus(chainID string) (state bool, height int64, err error) {
	return cm.Relayer.GetChainStatus(chainID)
}
