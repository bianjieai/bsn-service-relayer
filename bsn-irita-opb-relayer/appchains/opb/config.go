package opb

import (
	"encoding/json"
	"github.com/spf13/viper"
	"math/rand"

	cfg "relayer/config"
)

const (
	Prefix = "opb"

	// base config
	ChainId         = "chain_id"
	KeyPath         = "key_path"
	KeyName         = "key_name"
	Passphrase      = "passphrase"
	RpcAddrsMap     = "rpc_addrs"
	GrpcAddrsMap    = "grpc_addrs"
	MonitorInterval = "monitor_interval"
	Timeout         = "timeout"
)

// BaseConfig defines the base config
type BaseConfig struct {
	KeyPath         string
	KeyName         string
	Passphrase      string
	RpcAddrsMap     map[string]string
	GrpcAddrsMap    map[string]string
	ChainId         string
	Timeout         uint
	MonitorInterval uint64
}

func (bc *BaseConfig) PrintConfig() {
}

// Config defines the specific chain config
type Config struct {
	BaseConfig
	ChainParams
}

// NewBaseConfig constructs a new BaseConfig instance from viper
func NewBaseConfig(v *viper.Viper) (*BaseConfig, error) {
	config := new(BaseConfig)
	config.KeyPath = v.GetString(cfg.GetConfigKey(Prefix, KeyPath))
	config.ChainId = v.GetString(cfg.GetConfigKey(Prefix, ChainId))
	config.KeyName = v.GetString(cfg.GetConfigKey(Prefix, KeyName))
	config.Passphrase = v.GetString(cfg.GetConfigKey(Prefix, Passphrase))
	config.RpcAddrsMap = v.GetStringMapString(cfg.GetConfigKey(Prefix, RpcAddrsMap))
	config.GrpcAddrsMap = v.GetStringMapString(cfg.GetConfigKey(Prefix, GrpcAddrsMap))
	config.Timeout = v.GetUint(cfg.GetConfigKey(Prefix, Timeout))
	config.MonitorInterval = v.GetUint64(cfg.GetConfigKey(Prefix, MonitorInterval))

	return config, nil
}
func randURL(m []string) string {
	if len(m) == 0 {
		return ""
	}
	for _, index := range rand.Perm(len(m)) {
		return m[index]
	}
	return ""
}

// ValidBaseConfig validates if the given bytes is valid BaseConfig
func ValidateBaseConfig(baseCfg []byte) error {
	var baseConfig BaseConfig
	return json.Unmarshal(baseCfg, &baseConfig)
}
