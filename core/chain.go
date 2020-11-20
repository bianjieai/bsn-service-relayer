package core

// ChainI defines the basic chain interface
type ChainI interface {
	GetChainID() string // chain ID getter
}

// IritaHubChainI defines the interface to interact with Irita-Hub
type IritaHubChainI interface {
	ChainI

	// send the interchain request and handle the response with the given callback
	SendInterchainRequest(request InterchainRequestI, cb ResponseCallback) error
}

// AppChainI defines the interface to interact with the application chain
type AppChainI interface {
	ChainI

	// listen to the interchain events with an event handler
	InterchainEventListener(cb InterchainEventHandler) error

	// send the response to the application chain
	SendResponse(requestID string, response ResponseI) error

	// iService market interface
	IServiceMarketI
}

// InterchainEventI abstracts the event which is triggered for an interchain service invocation
type InterchainEventI interface {
	GetRequestID() string   // request ID getter
	GetServiceName() string // service name getter
	GetInput() string       // request input getter
	GetTimeout() uint64     // request timeout getter
}

// InterchainRequestI is an interface for the interchain request
type InterchainRequestI interface {
	GetServiceName() string   // service name getter
	GetProvider() string      // provider getter
	GetInput() string         // request input getter
	GetTimeout() uint64       // request timeout getter
	GetServiceFeeCap() string // service fee cap getter
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
	// add a service binding to the iService market
	AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error

	// update the specified service binding
	UpdateServiceBinding(serviceName, provider, serviceFee string, qos uint64) error

	// get the service binding by the given service name from the iService market
	GetServiceBinding(serviceName string) (IServiceBinding, error)
}

// IServiceBinding defines an iService binding interface
type IServiceBinding interface {
	GetServiceName() string // service name getter
	GetSchemas() string     // service schemas
	GetProvider() string    // service provider
	GetServiceFee() string  // service fee
	GetQoS() uint64         // quality of service
}

// InterchainEventHandler defines the interchain event callback interface
type InterchainEventHandler func(icEvent InterchainEventI) error

// ResponseCallback defines the response callback interface
type ResponseCallback func(icRequestID string, response ResponseI)
