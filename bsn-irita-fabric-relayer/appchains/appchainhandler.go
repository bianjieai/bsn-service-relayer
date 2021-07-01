package appchains

type AppChainHandlerI interface {

	//注册
	RegisterChain([]byte) (uint64, error)

	//删除
	DeleteChain([]byte) error

	//获取信息
	GetChains() error

	//修改
	UpdateChain(data []byte) error

	AddServiceBinding(chainId string, data []byte) error

	UpdateServiceBinding(chainId string, data []byte) error

	GetServiceBinding(chainId string, serviceName string) (interface{}, error)
}
