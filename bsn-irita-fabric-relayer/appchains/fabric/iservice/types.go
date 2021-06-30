package iservice

// IServiceRequestEvent is the Fabric event which represents an interchain service invocation to Irita-Hub
type IServiceRequestEvent struct {
	RequestID   string
	ServiceName string
	Provider    string
	Input       string
	Timeout     uint64
}

// GetRequestID implements InterchainEventI
func (e IServiceRequestEvent) GetRequestID() string {
	return e.RequestID
}

// GetServiceName implements InterchainEventI
func (e IServiceRequestEvent) GetServiceName() string {
	return e.ServiceName
}

// GetInput implements InterchainEventI
func (e IServiceRequestEvent) GetInput() string {
	return e.Input
}

// GetTimeout implements InterchainEventI
func (e IServiceRequestEvent) GetTimeout() uint64 {
	return e.Timeout
}