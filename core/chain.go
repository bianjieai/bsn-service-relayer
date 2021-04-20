package core

// ChainI defines the basic chain interface
type ChainI interface {
	GetChainID() string // chain ID getter
}

// HubChainI defines the interface to interact with the Hub chain
type HubChainI interface {
	ChainI

	// send the interchain request and handle the response with the given callback
	SendInterchainRequest(request InterchainRequest, cb ResponseCallback) error
}

// AppChainI defines the interface to interact with the application chain
type AppChainI interface {
	ChainI

	// start the application chain monitor
	Start(handler InterchainRequestHandler) error

	// stop the application chain monitor
	Stop() error

	// get the current height
	GetHeight() int64

	// send the response to the application chain
	SendResponse(requestID string, response ResponseI) error

	// iService market interface
	IServiceMarketI
}

// AppChainFactoryI abstracts the application chain operation interface
type AppChainFactoryI interface {
	// build an application chain according to the given app chain type and params
	BuildAppChain(chainType string, chainParams []byte) (AppChainI, error)

	// get the unique chain ID according to the given app chain type and params
	GetChainID(chainType string, chainParams []byte) (string, error)

	// store the base config by the given app chain type
	StoreBaseConfig(chainType string, baseConfig []byte) error

	// delete the chain config by the given chain id
	DeleteChainConfig(chainType string, chainID string) error
}

// InterchainRequest defines the interchain service request
type InterchainRequest struct {
	ID              string // request ID
	ChainID         string // chain ID
	ContractAddress string // contract address
	ServiceName     string // service name
	Provider        string // provider address
	Input           string // request input
	Timeout         uint64 // request timeout
	ServiceFeeCap   string // service fee cap
}

// ResponseI defines the response related interfaces
type ResponseI interface {
	GetErrMsg() string              // error msg getter
	GetOutput() string              // response output getter
	GetInterchainRequestID() string // interchain request ID getter
}

// KeyManager defines the key management interface
type KeyManager interface {
	Add(name, passphrase string) (addr string, mnemonic string, err error)
	Delete(name, passphrase string) error
	Show(name, passphrase string) (addr string, err error)
	Import(name, passphrase, keyArmor string) error
	Export(name, passphrase string) (keyArmor string, err error)
	Recover(name, passphrase, mnemonic string) (addr string, err error)
}

// IServiceMarketI defines the interface for the iService market on the application chain
type IServiceMarketI interface {
	// AddServiceBinding add a service binding to the iService market
	AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error

	// update the specified service binding
	UpdateServiceBinding(serviceName, provider, serviceFee string, qos uint64) error

	// get the service binding by the given service name from the iService market
	GetServiceBinding(serviceName string) (ServiceBindingI, error)
}

// ServiceBindingI defines the iService binding interface
type ServiceBindingI interface {
	GetServiceName() string // service name getter
	GetSchemas() string     // service schemas
	GetProvider() string    // service provider
	GetServiceFee() string  // service fee
	GetQoS() uint64         // quality of service
}

// InterchainRequestHandler defines the interchain request handler interface
type InterchainRequestHandler func(chainID string, request InterchainRequest, txHash string) error

// ResponseCallback defines the response callback interface
type ResponseCallback func(icRequestID string, response ResponseI)
