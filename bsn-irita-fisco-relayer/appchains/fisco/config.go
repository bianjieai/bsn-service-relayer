package fisco

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/viper"

	"github.com/FISCO-BCOS/go-sdk/conf"

	cfg "relayer/config"
)

const (
	Prefix = "fisco"

	// base config
	ConnectionType  = "connection_type"
	CAFile          = "ca_file"
	CertFile        = "cert_file"
	KeyFile         = "key_file"
	SMCrypto        = "sm_crypto"
	PrivateKeyFile  = "priv_key_file"
	MonitorInterval = "monitor_interval"
)

// BaseConfig defines the base config
type BaseConfig struct {
	IsHTTP          bool
	CAFile          string
	KeyFile         string
	CertFile        string
	PrivateKey      []byte
	IsSMCrypto      bool
	MonitorInterval uint64
}

func (bc *BaseConfig) PrintConfig(){
}

// Config defines the specific chain config
type Config struct {
	BaseConfig
	ChainParams
}

// NewBaseConfig constructs a new BaseConfig instance from viper
func NewBaseConfig(v *viper.Viper) (*BaseConfig, error) {
	connType := v.GetString(cfg.GetConfigKey(Prefix, ConnectionType))
	caFile := v.GetString(cfg.GetConfigKey(Prefix, CAFile))
	certFile := v.GetString(cfg.GetConfigKey(Prefix, CertFile))
	keyFile := v.GetString(cfg.GetConfigKey(Prefix, KeyFile))
	smCrypto := v.GetBool(cfg.GetConfigKey(Prefix, SMCrypto))
	privKeyFile := v.GetString(cfg.GetConfigKey(Prefix, PrivateKeyFile))
	monitorInterval := v.GetUint64(cfg.GetConfigKey(Prefix, MonitorInterval))

	config := new(BaseConfig)

	if strings.EqualFold(connType, "rpc") {
		config.IsHTTP = true
	} else if strings.EqualFold(connType, "channel") {
		config.IsHTTP = false
	} else {
		return nil, fmt.Errorf("connection type %s is not supported", connType)
	}

	config.IsSMCrypto = smCrypto

	keyBytes, curve, err := conf.LoadECPrivateKeyFromPEM(privKeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key, err: %v", err)
	}

	if config.IsSMCrypto && curve != "sm2p256v1" {
		return nil, fmt.Errorf("smcrypto must use sm2p256v1 private key, but found %s", curve)
	}
	if !config.IsSMCrypto && curve != "secp256k1" {
		return nil, fmt.Errorf("must use secp256k1 private key, but found %s", curve)
	}

	config.PrivateKey = keyBytes
	config.CAFile = caFile
	config.CertFile = certFile
	config.KeyFile = keyFile
	config.MonitorInterval = monitorInterval

	return config, nil
}

// NewConfig constructs a new Config instance
func NewConfig(baseConfig BaseConfig, chainParams ChainParams) *Config {
	return &Config{
		BaseConfig:  baseConfig,
		ChainParams: chainParams,
	}
}

//randURL returns a rand URL
func randURL(m map[string]string) string {
	r := rand.Intn(len(m))
	for _,v := range m {
		if r == 0 {
			return v
		}
		r--
	}
	return ""
}

// BuildClientConfig builds the FISCO client config from the given Config
func BuildClientConfig(config Config) *conf.Config {
	return &conf.Config{
		IsHTTP:     config.IsHTTP,
		CAFile:     config.CAFile,
		Key:        config.KeyFile,
		Cert:       config.CertFile,
		PrivateKey: config.PrivateKey,
		IsSMCrypto: config.IsSMCrypto,
		GroupID:    GetFiscoGroupID(config.ChainID),
		ChainID:    GetFiscoChainID(config.ChainID),
		NodeURL:    randURL(config.NodeURLs),
	}
}

// ValidBaseConfig validates if the given bytes is valid BaseConfig
func ValidateBaseConfig(baseCfg []byte) error {
	var baseConfig BaseConfig
	return json.Unmarshal(baseCfg, &baseConfig)
}
