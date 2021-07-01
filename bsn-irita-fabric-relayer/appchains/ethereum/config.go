package ethereum

import (
	"github.com/spf13/viper"

	cfg "relayer/config"
)

const (
	Prefix = "eth"

	ChainID     = "chain_id"
	NodeRPCAddr = "node_rpc_addr"
	GasLimit    = "gas_limit"
	GasPrice    = "gas_price"
	Key         = "key"
	Passphrase  = "passphrase"

	IServiceCoreAddr   = "iservice_core_addr"
	IServiceMarketAddr = "iservice_market_addr"
	IServiceEventName  = "iservice_event_name"
	IServiceEventSig   = "iservice_event_sig"
)

// Config represents the Ethereum chain config
type Config struct {
	ChainID     string `yaml:"chain_id"`
	NodeRPCAddr string `yaml:"node_rpc_addr"`
	GasLimit    uint64 `yaml:"gas_limit"`
	GasPrice    uint64 `yaml:"gas_price"`
	Key         string `yaml:"key"`
	Passphrase  string `yaml:"passphrase"`

	IServiceCoreAddr   string `yaml:"iservice_core_addr"`
	IServiceMarketAddr string `yaml:"iservice_market_addr"`
	IServiceEventName  string `yaml:"iservice_event_name"`
	IServiceEventSig   string `yaml:"iservice_event_sig"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ChainID:            v.GetString(cfg.GetConfigKey(Prefix, ChainID)),
		NodeRPCAddr:        v.GetString(cfg.GetConfigKey(Prefix, NodeRPCAddr)),
		GasLimit:           v.GetUint64(cfg.GetConfigKey(Prefix, GasLimit)),
		GasPrice:           v.GetUint64(cfg.GetConfigKey(Prefix, GasPrice)),
		Key:                v.GetString(cfg.GetConfigKey(Prefix, Key)),
		Passphrase:         v.GetString(cfg.GetConfigKey(Prefix, Passphrase)),
		IServiceCoreAddr:   v.GetString(cfg.GetConfigKey(Prefix, IServiceCoreAddr)),
		IServiceMarketAddr: v.GetString(cfg.GetConfigKey(Prefix, IServiceMarketAddr)),
		IServiceEventName:  v.GetString(cfg.GetConfigKey(Prefix, IServiceEventName)),
		IServiceEventSig:   v.GetString(cfg.GetConfigKey(Prefix, IServiceEventSig)),
	}
}
