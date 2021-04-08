package appchains

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"relayer/appchains/fisco"
	// "relayer/appchains/ethereum"
	"relayer/core"
	"relayer/config"
	"relayer/store"
)

// AppChainFactory defines an application chain factory
type AppChainFactory struct {
	Store *store.Store // store
}

// AppChainFactory defines an application chain factory
type BaseConfigFactory struct {
	config *viper.Viper// store
}

// NewAppChainFactory constructs a new application chain factory
func NewAppChainFactory(store *store.Store) *AppChainFactory {
	return &AppChainFactory{
		Store: store,
	}
}

// NewAppChainFactory constructs a new application chain factory
func NewBaseConfigFactory(v *viper.Viper) *BaseConfigFactory {
	return &BaseConfigFactory{
		config: v,
	}
}

// BuildAppChain implements AppChainFactoryI
func (f *AppChainFactory) BuildAppChain(chainType string, chainParams []byte) (core.AppChainI, error) {
	switch strings.ToLower(chainType) {
	case "eth":
		return nil, nil

	case "fabric":
		return nil, nil

	case "fisco":
		return fisco.BuildFISCOChain(chainParams, f.Store)

	default:
		return nil, fmt.Errorf("application chain %s not supported", chainType)
	}
}

// GetChainID implements AppChainFactoryI
func (f *AppChainFactory) GetChainID(chainType string, chainParams []byte) (chainID string, err error) {
	switch strings.ToLower(chainType) {
	case "eth":
		return "", nil

	case "fabric":
		return "", nil

	case "fisco":
		return fisco.GetChainIDFromBytes(chainParams)

	default:
		return "", fmt.Errorf("application chain %s not supported", chainType)
	}
}

// StoreBaseConfig implements AppChainFactoryI
func (f *AppChainFactory) StoreBaseConfig(chainType string, baseConfig []byte) error {
	switch strings.ToLower(chainType) {
	case "eth":
		return nil

	case "fabric":
		return nil

	case "fisco":
		return fisco.StoreBaseConfig(f.Store, baseConfig)

	default:
		return fmt.Errorf("application chain %s not supported", chainType)
	}
}

// StoreBaseConfig implements AppChainFactoryI
func (bc *BaseConfigFactory) NewBaseConfig(chainType string) (config.BaseConfigI, error) {
	switch strings.ToLower(chainType) {
	case "eth":
		return nil, nil

	case "fabric":
		return nil, nil

	case "fisco":
		return fisco.NewBaseConfig(bc.config)

	default:
		return nil, fmt.Errorf("application chain %s not supported", chainType)
	}
}
