package entity

import (
	"strings"
	"time"
)

type FabricRelayer struct {
	ID uint64
	//ChainId        uint64
	ChainBase
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
