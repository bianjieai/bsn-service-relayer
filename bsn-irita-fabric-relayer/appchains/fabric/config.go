package fabric

import (
	"github.com/spf13/viper"

	cfg "relayer/config"
)

const (
	Prefix         = "fabric"
	ChannelID      = "channel_id"
	ChainCodeID    = "chaincode_id"
	PeerRPCAddrs   = "peer_rpc_addrs"
	OrdererRPCAddr = "orderer_rpc_addr"
	Key            = "key"
	Passphrase     = "passphrase"
)

// Config represents the Fabric chain config
type Config struct {
	ChannelID      string   `yaml:"channel_id"`
	ChainCodeID    string   `yaml:"chaincode_id"`
	PeerRPCAddrs   []string `yaml:"peer_rpc_addrs"`
	OrdererRPCAddr string   `yaml:"orderer_rpc_addr"`
	Key            string   `yaml:"key"`
	Passphrase     string   `yaml:"passphrase"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ChannelID:      v.GetString(cfg.GetConfigKey(Prefix, ChannelID)),
		ChainCodeID:    v.GetString(cfg.GetConfigKey(Prefix, ChainCodeID)),
		PeerRPCAddrs:   v.GetStringSlice(cfg.GetConfigKey(Prefix, PeerRPCAddrs)),
		OrdererRPCAddr: v.GetString(cfg.GetConfigKey(Prefix, OrdererRPCAddr)),
		Key:            v.GetString(cfg.GetConfigKey(Prefix, Key)),
		Passphrase:     v.GetString(cfg.GetConfigKey(Prefix, Passphrase)),
	}
}
