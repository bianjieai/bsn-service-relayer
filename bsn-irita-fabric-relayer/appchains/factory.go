package appchains

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"relayer/appchains/fabric"
	"strings"

	"relayer/core"
	"relayer/store"
)

// AppChainFactory defines an application chain factory
type AppChainFactory struct {
	Store *store.Store // store
}

// NewAppChainFactory constructs a new application chain factory
func NewAppChainFactory(store *store.Store) *AppChainFactory {
	return &AppChainFactory{
		Store: store,
	}
}

func NewAppChainHandler(chainType string, hub core.HubChainI, log *log.Logger, v *viper.Viper) AppChainHandlerI {

	return fabric.NewFabricHandler(hub, log, v)
}

func (f *AppChainFactory) GetAppChainHandler(chainType string, hub core.HubChainI, log *log.Logger, v *viper.Viper) AppChainHandlerI {

	return fabric.NewFabricHandler(hub, log, v)
}

// BuildAppChain implements AppChainFactoryI
func (f *AppChainFactory) BuildAppChain(chainType string, chainParams []byte) (core.AppChainI, error) {
	switch strings.ToLower(chainType) {
	case "eth":
		// return ethereum.MakeEthChain(ethereum.NewConfig(f.Config)), nil
		return nil, nil

	case "fabric":
		// return fabric.MakeFabricChain(fabric.NewConfig(f.Config)), nil
		return nil, nil

	case "fisco":
		return nil, nil //fisco.BuildFISCOChain(chainParams, f.Store)

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
		return "", nil //fisco.GetChainIDFromBytes(chainParams)

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
		return nil //fisco.StoreBaseConfig(f.Store, baseConfig)

	default:
		return fmt.Errorf("application chain %s not supported", chainType)
	}
}
