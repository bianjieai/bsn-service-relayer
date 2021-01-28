package fabric

import (
	"fmt"
	"time"

	"relayer/appchains/fabric/iservice"
	"relayer/core"
)

// FabricChain defines the Fabric chain
type FabricChain struct {
	ChannelID      string
	ChainCodeID    string
	PeerRPCAddrs   []string
	OrdererRPCAddr string
	Client         interface{}
	Key            string
	Passphrase     string
}

// NewFabricChain constructs a new FabricChain instance
func NewFabricChain(
	channelID string,
	chainCodeID string,
	peerRPCAddrs []string,
	ordererRPCAddr string,
	key string,
	passphrase string,
) FabricChain {
	fabric := FabricChain{
		ChannelID:      channelID,
		ChainCodeID:    chainCodeID,
		PeerRPCAddrs:   peerRPCAddrs,
		OrdererRPCAddr: ordererRPCAddr,
		Key:            key,
		Passphrase:     passphrase,
	}

	fabric.Client = nil

	return fabric
}

// MakeFabricChain builds a fabric chain from the given config
func MakeFabricChain(config Config) FabricChain {
	return NewFabricChain(
		config.ChannelID,
		config.ChainCodeID,
		config.PeerRPCAddrs,
		config.OrdererRPCAddr,
		config.Key,
		config.Passphrase,
	)
}

// GetChainID implements AppChainI
func (fc FabricChain) GetChainID() string {
	return fc.ChannelID
}

// InterchainEventListener implements AppChainI
func (fc FabricChain) InterchainEventListener(cb core.InterchainRequestHandler) error {
	i := 0

	for {
		request := core.InterchainRequest{
			ID:          fmt.Sprintf("request%d", i+1),
			ServiceName: "exchange_rate",
			Input:       fmt.Sprintf(`{"name":"CNY-USD"}`),
			Timeout:     uint64(50),
		}

		cb(fc.GetChainID(), request)

		time.Sleep(10 * time.Second)
		i++
	}
}

// SendResponse implements AppChainI
func (fc FabricChain) SendResponse(requestID string, response core.ResponseI) error {
	return nil
}

// AddServiceBinding implements AppChainI
func (fc FabricChain) AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error {
	return nil
}

// UpdateServiceBinding implements AppChainI
func (fc FabricChain) UpdateServiceBinding(serviceName, provider, serviceFee string, qos uint64) error {
	return nil
}

// GetServiceBinding implements AppChainI
func (fc FabricChain) GetServiceBinding(serviceName string) (core.ServiceBindingI, error) {
	return iservice.ServiceBinding{
		Provider:   "test-provider",
		ServiceFee: "100point",
	}, nil
}
