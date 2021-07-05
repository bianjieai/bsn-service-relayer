package fisco

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/viper"
	"relayer/logging"

	"github.com/FISCO-BCOS/go-sdk/conf"

	cfg "relayer/config"
)

const (
	Prefix = "fisco"

	// base config
	ChainId         = "chainId"
	ConnectionType  = "connection_type"
	CAFile          = "ca_file"
	CertFile        = "cert_file"
	KeyFile         = "key_file"
	SMCrypto        = "sm_crypto"
	PrivateKeyFile  = "priv_key_file"
	MonitorInterval = "monitor_interval"
	Nodes           = "nodes"
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
	NodesMap        map[string]string
	ChainId         int64
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

	chainId := v.GetInt64(cfg.GetConfigKey(Prefix, ChainId))
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
	if chainId == 0 {
		chainId = 1
	}
	config.ChainId = chainId
	config.PrivateKey = keyBytes
	config.CAFile = caFile
	config.CertFile = certFile
	config.KeyFile = keyFile
	config.MonitorInterval = monitorInterval

	config.NodesMap = v.GetStringMapString(cfg.GetConfigKey(Prefix, Nodes))
	logging.Logger.Infof("config fisco nods : %v", config.NodesMap)

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

// BuildClientConfig builds the FISCO client config from the given Config
func BuildClientConfig(config Config) *conf.Config {
	//将接口传递的节点名称通过配置转换为 节点地址，如果不在配置中，不转换
	//随机取一个传入的node
	nodeName := randURL(config.NodeURLs)
	//获取配置的nodeURL
	nodeUrl, ok := config.NodesMap[nodeName]
	if ok {
		nodeName = nodeUrl
	}

	return &conf.Config{
		IsHTTP:     config.IsHTTP,
		CAFile:     config.CAFile,
		Key:        config.KeyFile,
		Cert:       config.CertFile,
		PrivateKey: config.PrivateKey,
		IsSMCrypto: config.IsSMCrypto,
		GroupID:    config.GroupID,
		ChainID:    config.BaseConfig.ChainId,
		NodeURL:    nodeName,
	}
}

// ValidBaseConfig validates if the given bytes is valid BaseConfig
func ValidateBaseConfig(baseCfg []byte) error {
	var baseConfig BaseConfig
	return json.Unmarshal(baseCfg, &baseConfig)
}
