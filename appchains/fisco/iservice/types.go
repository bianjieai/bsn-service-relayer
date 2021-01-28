package iservice

import (
	"math/big"
)

// ServiceBinding defines an iService binding
type ServiceBinding struct {
	ServiceName string   // service name
	Schemas     string   // input and output schemas
	Provider    string   // service provider
	ServiceFee  string   // service fee
	QoS         *big.Int // quality of service, in terms of the minimum response time
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
	return b.QoS.Uint64()
}
