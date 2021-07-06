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
}

// AppChainFactoryI abstracts the application chain operation interface
type AppChainFactoryI interface {
	// build an application chain according to the given app chain type and params
	BuildAppChain(chainType string, chainParams []byte) (AppChainI, error)

	// get the unique chain ID according to the given app chain type and params
	GetChainID(chainType string, chainParams []byte) (string, error)

	// store the base config by the given app chain type
	StoreBaseConfig(chainType string, baseConfig []byte) error
}

// InterchainRequest defines the interchain service request
type InterchainRequest struct {
	ID              string // request ID
	SourceChainID   string // source chain ID
	DestChainID     string // target chain ID
	DestSubChainID  string // target sub chain ID
	DestChainType  string // target chain type
	EndpointAddress string // end point address
	EndpointType    string // end point type
	Method          string // method name
	CallData        []byte // target method name and json string of arguments
	TxHash          string // source transaction hash
	Sender          string // message sender
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
// InterchainRequestHandler defines the interchain request handler interface
type InterchainRequestHandler func(chainID string, request InterchainRequest, txHash string) error

// ResponseCallback defines the response callback interface
type ResponseCallback func(icRequestID string, response ResponseI)
