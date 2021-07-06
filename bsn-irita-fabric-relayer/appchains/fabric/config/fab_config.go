package config

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"relayer/appchains/fabric/config/redconfig"
	"relayer/appchains/fabric/config/redconfig/configbackend"
)

type FabricConfig struct {
	SdkConfig string

	OrgName     string
	MspUserName string

	OrgCode string
}

func (f *FabricConfig) GetSdkConfig(channelId string, nodes []string) core.ConfigProvider {

	ch := configbackend.ChannelConfig{ChannelId: channelId, PeerName: nodes[0]}

	var s []redconfig.SetOption
	s = append(s, redconfig.SetChannel(&ch))

	configProvider := redconfig.FromFile(f.SdkConfig, s)
	return configProvider

}
