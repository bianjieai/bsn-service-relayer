package server

import "fmt"

const (
	CODE_SUCCESS = 0
	CODE_ERROR   = 1
)

// AddChainRequest defines the request to add an app chain
type AddChainRequest struct {
	ChainParams string `json:"chain_params"`
}

// AddChainResult defines the result for adding an app chain
type AddChainResult struct {
	ChainID string `json:"chain_id"`
}

type AddChainAndBindServResult struct {
	ChainID string `json:"chain_id"`
	ServiceName string `json:"service_name"`
}
// ChainStatus defines the chain status
type ChainStatus struct {
	State  bool  `json:"state"`
	Height int64 `json:"height,omitempty"`
}

type AddChainAndBindServRequest struct {
	ChainParams string `json:"chain_params"`
	ServParamsPath string  `json:"serv_params_path,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	Schemas     string `json:"schemas,omitempty"`
	Provider    string `json:"provider,omitempty"`
	ServiceFee  string `json:"service_fee,omitempty"`
	QoS         uint64 `json:"qos,omitempty"`
}

// AddServiceBindingRequest defines the request to add a service binding
type AddServiceBindingRequest struct {
	ServiceName string `json:"service_name"`
	Schemas     string `json:"schemas"`
	Provider    string `json:"provider"`
	ServiceFee  string `json:"service_fee"`
	QoS         uint64 `json:"qos"`
}

// UpdateServiceBindingRequest defines the request to update a service binding
type UpdateServiceBindingRequest struct {
	Provider   string `json:"provider"`
	ServiceFee string `json:"service_fee"`
	QoS        uint64 `json:"qos"`
}

// SuccessResponse defines the response on success
type SuccessResponse struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result,omitempty"`
}

// ErrorResponse defines the response on error
type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// ValidateChainID validates the given chain ID
func ValidateChainID(chainID string) error {
	if len(chainID) == 0 {
		return fmt.Errorf("chain ID can not be empty")
	}

	return nil
}

// ValidateServiceName validates the given service name
func ValidateServiceName(serviceName string) error {
	if len(serviceName) == 0 {
		return fmt.Errorf("service name can not be empty")
	}

	return nil
}
