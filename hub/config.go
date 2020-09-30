package hub

import (
	"github.com/spf13/viper"

	"github.com/irisnet/service-sdk-go/types"

	cmn "relayer/common"
	cfg "relayer/config"
)

// default config variables
var (
	defaultChainID       = "irita-hub"
	defaultNodeRPCAddr   = "http://127.0.0.1:26657"
	defaultKeyPath       = cmn.MustGetHomeDir()
	defaultGas           = uint64(200000)
	defaultFee           = "4point"
	defaultBroadcastMode = types.Commit
	defaultKeyAlgorithm  = "sm2"
)

const (
	Prefix      = "hub"
	ChainID     = "chain_id"
	NodeRPCAddr = "node_rpc_addr"
	KeyPath     = "key_path"
	KeyName     = "key_name"
	Passphrase  = "passphrase"
)

// Config is a config struct for IRITA-HUB
type Config struct {
	ChainID     string `yaml:"chain_id"`
	NodeRPCAddr string `yaml:"node_rpc_addr"`
	KeyPath     string `yaml:"key_path"`
	KeyName     string `yaml:"key_name"`
	Passphrase  string `yaml:"passphrase"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ChainID:     v.GetString(cfg.GetConfigKey(Prefix, ChainID)),
		NodeRPCAddr: v.GetString(cfg.GetConfigKey(Prefix, NodeRPCAddr)),
		KeyPath:     v.GetString(cfg.GetConfigKey(Prefix, KeyPath)),
		KeyName:     v.GetString(cfg.GetConfigKey(Prefix, KeyName)),
		Passphrase:  v.GetString(cfg.GetConfigKey(Prefix, Passphrase)),
	}
}
