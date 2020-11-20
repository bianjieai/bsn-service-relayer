package appchains

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"relayer/appchains/bcos"
	"relayer/appchains/ethereum"
	"relayer/appchains/fabric"
	"relayer/core"
	"relayer/store"
)

// AppChainFactory defines an application chain factory
type AppChainFactory struct {
	Config *viper.Viper // configuration context in which the application chain is built
	Store  *store.Store // store for the application chain
}

// NewAppChainFactory constructs a new application chain factory
func NewAppChainFactory(config *viper.Viper, store *store.Store) AppChainFactory {
	return AppChainFactory{
		Config: config,
		Store:  store,
	}
}

// Make builds an application chain according to the given chain name
func (f AppChainFactory) Make(chainName string) (core.AppChainI, error) {
	switch strings.ToLower(chainName) {
	case "eth":
		return ethereum.MakeEthChain(ethereum.NewConfig(f.Config)), nil

	case "fabric":
		return fabric.MakeFabricChain(fabric.NewConfig(f.Config)), nil

	case "bcos":
		return bcos.MakeBCOSChain(bcos.NewConfig(f.Config), f.Store), nil

	default:
		return nil, fmt.Errorf("application chain %s not supported", chainName)
	}
}
