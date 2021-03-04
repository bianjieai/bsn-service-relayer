package core

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
)

// Relayer represents a relayer transmitting msgs
// from app chains with the same architecture
// to the Hub chain
type Relayer struct {
	AppChainType    string
	HubChain        HubChainI
	AppChains       map[string]AppChainI
	AppChainStates  map[string]bool
	AppChainFactory AppChainFactoryI
	Logger          *log.Logger
	mtx             sync.Mutex
}

// NewRelayer constructs a new Relayer instance
func NewRelayer(appChainType string, hub HubChainI, appChainFactory AppChainFactoryI, logger *log.Logger) *Relayer {
	return &Relayer{
		AppChainType:    appChainType,
		HubChain:        hub,
		AppChainFactory: appChainFactory,
		Logger:          logger,
	}
}

// AddChain adds an app chain with the specified app chain params
func (r *Relayer) AddChain(appChainParams []byte) (chainID string, err error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	chainID, err = r.AppChainFactory.GetChainID(r.AppChainType, appChainParams)
	if err != nil {
		return "", err
	}

	_, ok := r.AppChains[chainID]
	if ok {
		return "", fmt.Errorf("chain ID %s already exists", chainID)
	}

	chain, err := r.AppChainFactory.BuildAppChain(r.AppChainType, appChainParams)
	if err != nil {
		return "", err
	}

	if err := chain.Start(r.HandleInterchainRequest); err != nil {
		return "", err
	}

	r.AppChains[chainID] = chain
	r.AppChainStates[chainID] = true

	return chainID, nil
}

// StartChain starts the specified app chain
func (r *Relayer) StartChain(chainID string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	state, ok := r.AppChainStates[chainID]
	if !ok {
		return fmt.Errorf("chain ID %s does not exist", chainID)
	}

	if state {
		return fmt.Errorf("chain ID %s is running", chainID)
	}

	chain := r.AppChains[chainID]
	if err := chain.Start(r.HandleInterchainRequest); err != nil {
		return err
	}

	r.AppChainStates[chainID] = true

	return nil
}

// StopChain stops the specified app chain
func (r *Relayer) StopChain(chainID string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	state, ok := r.AppChainStates[chainID]
	if !ok {
		return fmt.Errorf("chain ID %s does not exist", chainID)
	}

	if !state {
		return fmt.Errorf("chain ID %s is not running", chainID)
	}

	chain := r.AppChains[chainID]
	if err := chain.Stop(); err != nil {
		return err
	}

	r.AppChainStates[chainID] = false

	return nil
}

// GetChain gets the specified app chain
func (r *Relayer) GetChain(chainID string) (appChain AppChainI, err error) {
	appChain, ok := r.AppChains[chainID]
	if !ok {
		return nil, fmt.Errorf("chain ID %s does not exist", chainID)
	}

	return appChain, nil
}

// GetChains retrieves the current active app chains
func (r *Relayer) GetChains() []string {
	chains := make([]string, 0)

	r.mtx.Lock()
	defer r.mtx.Unlock()

	for c, s := range r.AppChainStates {
		if s {
			chains = append(chains, c)
		}
	}

	return chains
}

// GetChainStatus gets the status of the specified app chain
func (r *Relayer) GetChainStatus(chainID string) (state bool, height int64, err error) {
	state, ok := r.AppChainStates[chainID]
	if !ok {
		return state, height, fmt.Errorf("chain ID %s does not exist", chainID)
	}

	height = r.AppChains[chainID].GetHeight()

	return state, height, nil
}
