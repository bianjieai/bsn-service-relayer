package appchains

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"relayer/appchains/bcos"
	"relayer/appchains/ethereum"
	"relayer/appchains/fabric"
	"relayer/core"
)

// AppChainFactory defines an application chain factory
type AppChainFactory struct {
	Config *viper.Viper // config context in which the application chain is built
}

// NewAppChainFactory constructs a new application chain factory
func NewAppChainFactory(config *viper.Viper) AppChainFactory {
	return AppChainFactory{Config: config}
}

// Make builds an application chain according to the given chain name
func (f AppChainFactory) Make(chainName string) (core.AppChainI, error) {
	switch strings.ToLower(chainName) {
	case "eth":
		return ethereum.MakeEthChain(ethereum.NewConfig(f.Config)), nil

	case "fabric":
		return fabric.MakeFabricChain(fabric.NewConfig(f.Config)), nil

	case "bcos":
		return bcos.MakeBCOSChain(bcos.NewConfig(f.Config)), nil

	default:
		return nil, fmt.Errorf("application chain %s not supported", chainName)
	}
}
