package fabric

type serviceCallInfo struct {
	Request  *serviceRequest  `json:"request,omitempty"`
	Response *serviceResponse `json:"response,omitempty"`
	Service  *serviceBindInfo `json:"service,omitempty"`
	Status   string           `json:"status"`
	Type     string           `json:"type"`
}

type serviceRequest struct {
	RequestId   string        `json:"requestID,omitempty"`   //服务请求 ID  本合约中使用 合约交易ID
	ServiceName string        `json:"serviceName,omitempty"` //服务定义名称
	Input       string        `json:"input,omitempty"`       //服务请求输入；需符合服务的输入规范
	Timeout     uint64        `json:"timeout,omitempty"`     //请求超时时间；在目标链上等待的最大区块数
	CallBack    *CallBackInfo `json:"callback"`              //回调的合约以及方法

	//todo new p
	CrossChainCodeId string
	ChainId          string
	ServiceFeeCap    string
}

type CallBackInfo struct {
	ChainCode string `json:"chainCode"`
	FuncName  string `json:"funcName"`
}

type serviceResponse struct {
	RequestId   string `json:"requestID,omitempty"` //服务请求 ID  本合约中使用 合约交易ID
	ErrMsg      string `json:"errMsg,omitempty"`
	Output      string `json:"output,omitempty"`
	IcRequestId string `json:"icRequestID,omitempty"`
}

type updateServiceInfo struct {
	ServiceName string `json:"serviceName"`
	ServiceFee  string `json:"serviceFee"`
	Provider    string `json:"provider"`
	Qos         uint64 `json:"qos"`
}

type serviceBindInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Schemas     string `json:"schemas"`
	Provider    string `json:"provider"`
	ServiceFee  string `json:"serviceFee"`
	Qos         uint64 `json:"qos"`
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
	ServiceName string `json:"service_name"`
	Provider    string `json:"provider"`
	ServiceFee  string `json:"service_fee"`
	QoS         uint64 `json:"qos"`
}
