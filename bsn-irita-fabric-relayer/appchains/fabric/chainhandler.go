package fabric

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"relayer/appchains/fabric/config"
	"relayer/appchains/fabric/entity"
	"relayer/appchains/fabric/store"
	"relayer/core"
	"relayer/errors"
	"relayer/logging"
	"time"
)

const (
	fabric_sdk_config    = "fabric.sdk_config"
	fabric_msp_user_name = "fabric.msp_user_name"
	fabric_org_name      = "fabric.org_name"
	base_mysql_conn      = "base.mysql_conn"
	base_city_code       = "base.city_code"
)

func NewFabricHandler(hub core.HubChainI, log *log.Logger, v *viper.Viper) *fabricHandler {

	store.InitMysql(v.GetString(base_mysql_conn))

	logging.Logger.Info("new fabric handler ")
	conf := &config.FabricConfig{
		SdkConfig:   v.GetString(fabric_sdk_config),
		MspUserName: v.GetString(fabric_msp_user_name),
		OrgName:     v.GetString(fabric_org_name),
		OrgCode:     v.GetString(base_city_code),
	}
	logging.Logger.Infof("fabric config is %v", conf)

	fabric := &fabricHandler{
		Logger:          log,
		HubChain:        hub,
		Config:          conf,
		AppChains:       make(map[string]core.AppChainI),
	}

	fabric.initTask()

	return fabric
}

type fabricHandler struct {
	Logger          *log.Logger
	HubChain        core.HubChainI
	AppChains       map[string]core.AppChainI

	Config *config.FabricConfig
}

func (f *fabricHandler) initTask() {

	logging.Logger.Info("InitTask")

	chains, err := store.GetRelayerAppInfos()
	if err != nil {
		logging.Logger.Errorf("get Relayer AppInfos failed %s", err.Error())
		return
	}

	for _, sc := range chains {
		logging.Logger.Infof("Chain is %v,Nodes is %v", sc, sc.GetNodes())
		_, ok := f.AppChains[sc.GetChainId()]
		if ok {
			continue
		}

		chain, err := NewFabricChain(sc, f.Config)
		if err != nil {
			logging.Logger.Errorf("the fabric chain init failed %s", err.Error())
			continue
		}
		err = chain.Start(f.HandleInterchainRequest)

		if err != nil {
			logging.Logger.Errorf("the fabric chain init failed %s", err.Error())
			continue
		}

		f.AppChains[sc.GetChainId()] = chain
	}

}

func (f *fabricHandler) RegisterChain(data []byte) (uint64, error) {

	rc := &entity.RegisterChain{}

	err := json.Unmarshal(data, rc)
	if err != nil {
		logging.Logger.Errorf("invalid JSON Params %s", err.Error())
		return rc.ChainId, errors.New("invalid JSON Params")
	}
	logging.Logger.Infof("this register data is %v", rc)

	// Check if it already exists
	chain, ok := f.AppChains[rc.GetChainId()]
	if ok {
		return rc.ChainId, nil //fmt.Errorf("chain ID %s already exists", rc.ChainId)
	}

	// ToStoreData
	sc := rc.ToStoreData(f.Config.OrgCode)

	// Start NewFabricChain
	chain, err = NewFabricChain(sc, f.Config)
	if err != nil {
		return rc.ChainId, errors.New("the fabric chain init failed")
	}
	err = chain.Start(f.HandleInterchainRequest)

	if err != nil {
		return rc.ChainId, errors.New("the fabric chain start failed")
	}

	f.AppChains[rc.GetChainId()] = chain

	//以上执行成功，向数据库存储，
	err = store.StoreRelayerAppInfo(sc)
	if err != nil {
		//存储失败，停止服务
		chain.Stop()
		delete(f.AppChains,rc.GetChainId())

		logging.Logger.Errorf("the appinfo store failed %s", err.Error())
		return rc.ChainId, errors.New("the appinfo store failed %s", err.Error())
	}

	return rc.ChainId, nil
}

func (f *fabricHandler) DeleteChain(data []byte) error {
	rc := &entity.ChainBase{}

	err := json.Unmarshal(data, rc)
	if err != nil {
		return errors.New("invalid JSON Params")
	}
	logging.Logger.Infof("this delete data is %v", rc)

	// Check if it already exists
	_, ok := f.AppChains[rc.GetChainId()]
	if !ok {
		return nil
	}

	// Delete DB
	logging.Logger.Infof("start delete ChainId: %d", rc.ChainId)
	err = store.DeleteRelayerAppInfo(rc.GetChainId())
	if err != nil {
		logging.Logger.Errorf("delete ChainId failed: %s", err.Error())
		return err
	}

	// Stop FabricChain
	err = f.AppChains[rc.GetChainId()].Stop()
	if err != nil {
		logging.Logger.Errorf("Stop ChainId failed: %s", err.Error())
		return errors.New("the fabric chain stop failed")
	}

	// Delete Map
	delete(f.AppChains, rc.GetChainId())

	return nil
}

func (f *fabricHandler) GetChains() error {
	return nil
}

func (f *fabricHandler) UpdateChain(data []byte) error {
	rc := &entity.UpdateChain{}

	err := json.Unmarshal(data, rc)
	if err != nil {
		logging.Logger.Errorf("invalid JSON Params %s", err.Error())
		return errors.New("invalid JSON Params")
	}
	logging.Logger.Infof("this UpdateChain register data is %v", rc)

	// Check if it already exists
	_, ok := f.AppChains[rc.GetChainId()]
	if !ok {
		return errors.New("the fabric chain not already exists")
	}

	// Update DB
	updateFabricRelayer := rc.UpdateStoreData()
	err = store.UpdateRelayerAppInfo(updateFabricRelayer)
	if err != nil {
		logging.Logger.Errorf("the appinfo update failed %s", err.Error())
		return errors.New("the appinfo update failed")
	}

	// Stop FabricChain
	err = f.AppChains[rc.GetChainId()].Stop()
	if err != nil {
		logging.Logger.Errorf("the fabric chain stop failed %s", err.Error())
		return errors.New("the fabric chain stop failed")
	}

	// Start new FabricChain
	chain, err := NewFabricChain(updateFabricRelayer, f.Config)
	if err != nil {
		return errors.New("the fabric chain init failed")
	}

	err = chain.Start(f.HandleInterchainRequest)
	if err != nil {
		return errors.New("the fabric chain start failed")
	}

	f.AppChains[updateFabricRelayer.GetChainId()] = chain

	return nil
}

func (r *fabricHandler) HandleInterchainRequest(chainID string, request core.InterchainRequest, txHash string) error {

	r.Logger.Infof("got the interchain request on %s: %+v", chainID, request)
	r.Logger.Infof("txHash : %s", txHash)

	// 交易记录  insert
	// request_id		[request.ID]
	// from_chanId		chainId
	// from_tx			txHash
	// tx_createtime
	// tx_status 0 未知
	// source_service 0 relayer

	interchainRequestInfo := entity.FabricRelayerTx{
		Request_id:     request.ID,
		From_chainid:   chainID,
		From_tx:        txHash,
		Tx_createtime:  time.Now(),
		Tx_status:      0,
		Source_service: 0,
	}

	store.InsertInterchainRequestInfo(&interchainRequestInfo)

	callback := func(icRequestID string, response core.ResponseI) {
		// 跨链交易回复

		r.Logger.Infof(
			"got the response of the interchain request on %s: %+v",
			r.HubChain.GetChainID(),
			response,
		)

		err := r.AppChains[chainID].SendResponse(request.ID, response)
		if err != nil {
			r.Logger.Errorf(
				"failed to send the response to %s: %s",
				chainID,
				err,
			)

			return
		}

		r.Logger.Infof(
			"response sent to %s successfully",
			chainID,
		)
	}

	err := r.HubChain.SendInterchainRequest(request, callback)
	if err != nil {
		r.Logger.Errorf(
			"failed to handle the interchain request %+v on %s: %s",
			request,
			r.HubChain.GetChainID(),
			err,
		)

		return err
	}

	r.Logger.Infof("HandleInterchainRequest is End !!!")
	return nil
}