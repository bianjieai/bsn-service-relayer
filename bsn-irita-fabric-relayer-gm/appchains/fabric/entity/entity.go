package entity

import (
	"strconv"
	"time"
)

type ChainBase struct {
	ChainId uint64 `json:"chainId"`
}

func (r *ChainBase) SetChainId(cid string) error {

	ccid, err := strconv.ParseUint(cid, 10, 64)
	if err != nil {
		return err
	}
	r.ChainId = ccid
	return nil
}

func (r *ChainBase) GetChainId() string {
	return strconv.FormatUint(r.ChainId, 10)
}

func (r *ChainBase) GetChainIdKey() string {
	return "chainId_" + strconv.FormatUint(r.ChainId, 10)
}

type RegisterChain struct {
	ChainBase

	AppCode string `json:"appCode"`

	CrossChainCode string `json:"ccm"`

	ChannelId string `json:"channelId"`

	Nodes []string `json:"nodes"`
}

func (r *RegisterChain) ToStoreData(cityCode string) *FabricRelayer {

	relayer := &FabricRelayer{
		ChainBase: ChainBase{ChainId: r.ChainId},
		AppCode:   r.AppCode,
		ChannelId: r.ChannelId,

		CrossChainCode: r.CrossChainCode,
		CityNode:       cityCode,
		Status:         1,
		CreateTime:     time.Now(),
	}
	relayer.SetNodes(r.Nodes)
	return relayer
}

type DeleteChain struct {
	ChainId uint64 `json:"chainId"`
}

type UpdateChain struct {
	ChainBase
	CrossChainCode string `json:"ccm"`

	ChannelId string `json:"channelId"`

	Nodes []string `json:"nodes"`
}

func (u *UpdateChain) UpdateStoreData() *FabricRelayer {
	data := &FabricRelayer{}
	data.ChainId = u.ChainId
	data.CrossChainCode = u.CrossChainCode
	data.ChannelId = u.ChannelId
	data.SetNodes(u.Nodes)
	return data
}
