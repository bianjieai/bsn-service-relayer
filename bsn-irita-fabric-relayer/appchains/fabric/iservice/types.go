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

// ServiceBinding defines an iService binding
type ServiceBinding struct {
	ServiceName string // service name
	Schemas     string // input and output schemas
	Provider    string // service provider
	ServiceFee  string // service fee
	QoS         uint64 // quality of service, in terms of the minimum response time
}

// GetServiceName implements IServiceBinding
func (b ServiceBinding) GetServiceName() string {
	return b.ServiceName
}

// GetSchemas implements IServiceBinding
func (b ServiceBinding) GetSchemas() string {
	return b.Schemas
}

// GetProvider implements IServiceBinding
func (b ServiceBinding) GetProvider() string {
	return b.Provider
}

// GetServiceFee implements IServiceBinding
func (b ServiceBinding) GetServiceFee() string {
	return b.ServiceFee
}

// GetQoS implements IServiceBinding
func (b ServiceBinding) GetQoS() uint64 {
	return b.QoS
}
