package fabric

import (
	"encoding/json"
	"github.com/BSNDA/fabric-sdk-go-gm/pkg/common/providers/core"
	"math/rand"
	"relayer/appchains/fabric/config/redconfig"
	"relayer/appchains/fabric/config/redconfig/configbackend"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	fabric_sdk_config       = "fabric.sdk_config"
	fabric_msp_user_name    = "fabric.msp_user_name"
	fabric_org_name         = "fabric.org_name"
	fabric_monitor_interval = "fabric.monitor_interval"
	base_mysql_conn         = "base.mysql_conn"
	base_city_code          = "base.city_code"
)

// BaseConfig defines the base config
type BaseConfig struct {
	SdkConfig       string
	OrgName         string
	MspUserName     string
	OrgCode         string
	MonitorInterval uint64
}

func (bc *BaseConfig) PrintConfig() {
}

// Config defines the specific chain config
type Config struct {
	BaseConfig
	*FabricRelayer
}

func (f *Config) GetSdkConfig(channelId string, nodes []string) core.ConfigProvider {

	ch := configbackend.ChannelConfig{ChannelId: channelId, PeerName: nodes[0]}

	var s []redconfig.SetOption
	s = append(s, redconfig.SetChannel(&ch))

	configProvider := redconfig.FromFile(f.SdkConfig, s)
	return configProvider

}

// NewBaseConfig constructs a new BaseConfig instance from viper
func NewBaseConfig(v *viper.Viper) (*BaseConfig, error) {
	return &BaseConfig{
		SdkConfig:       v.GetString(fabric_sdk_config),
		OrgName:         v.GetString(fabric_org_name),
		MspUserName:     v.GetString(fabric_msp_user_name),
		OrgCode:         v.GetString(base_city_code),
		MonitorInterval: v.GetUint64(fabric_monitor_interval),
	}, nil
}

//randURL returns a rand URL
func randURL(m map[string]string) string {
	r := rand.Intn(len(m))
	for _, v := range m {
		if r == 0 {
			return v
		}
		r--
	}
	return ""
}

// ValidBaseConfig validates if the given bytes is valid BaseConfig
func ValidateBaseConfig(baseCfg []byte) error {
	var baseConfig BaseConfig
	return json.Unmarshal(baseCfg, &baseConfig)
}

type FabricRelayer struct {
	ID             uint64
	ChainId        uint64
	AppCode        string
	ChannelId      string
	CrossChainCode string
	NodeName       string
	CityNode       string
	Status         int
	CreateTime     time.Time
	LastUpdateTime time.Time
}

func (f *FabricRelayer) SetNodes(nodes []string) {
	f.NodeName = strings.Join(nodes, ";")
}

func (f *FabricRelayer) GetNodes() []string {
	return strings.Split(f.NodeName, ";")
}
