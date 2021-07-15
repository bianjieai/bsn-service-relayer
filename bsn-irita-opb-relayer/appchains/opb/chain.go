package opb

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	sdk "github.com/bianjieai/irita-sdk-go"
	"github.com/bianjieai/irita-sdk-go/modules/wasm"
	"github.com/bianjieai/irita-sdk-go/types"
	sdktypes "github.com/bianjieai/irita-sdk-go/types"
	sdkstore "github.com/bianjieai/irita-sdk-go/types/store"
	abci "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"strings"
	"time"

	"relayer/core"
	"relayer/logging"
	"relayer/mysql"
	"relayer/store"
)

// opbChain defines the opb chain
type OpbChain struct {
	Config    Config
	OpbClient sdk.IRITAClient
	ChainID   string // unique chain ID

	store      *store.Store // store backend instance
	lastHeight int64        // last height

	done    bool                          // indicates if the chain monitor is done
	handler core.InterchainRequestHandler // handler for the interchain request
}

// NewFISCOChain constructs a new FISCOChain instance
func NewOpbChain(
	config Config,
	store *store.Store,
) (*OpbChain, error) {
	//将接口传递的节点名称通过配置转换为 节点地址，如果不在配置中，不转换
	//随机取一个传入的node
	nodeName := randURL(config.ChainParams.NodeURLs)
	var rpcAddr string
	var grpcAddr string
	//获取配置的nodeURL
	rpcAddrstr, ok := config.RpcAddrsMap[nodeName]
	if ok {
		rpcAddr = rpcAddrstr
	}
	grpcAddrstr, ok := config.GrpcAddrsMap[nodeName]
	if ok {
		grpcAddr = grpcAddrstr
	}

	options := []sdktypes.Option{
		sdktypes.CachedOption(true),
		sdktypes.KeyDAOOption(sdkstore.NewFileDAO(config.KeyPath)),
	}

	clientConfig, err := sdktypes.NewClientConfig(rpcAddr, grpcAddr, config.BaseConfig.ChainId, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to init clientConfig: %s", err)
	}
	opbClient := sdk.NewIRITAClient(clientConfig)
	chainID := GetChainID(config.ChainParams)

	opb := &OpbChain{
		Config:    config,
		OpbClient: opbClient,
		ChainID:   chainID,
		store:     store,
		done:      true,
	}

	err = opb.storeChainParams()
	if err != nil {
		return nil, err
	}

	err = opb.storeChainID()
	if err != nil {
		return nil, err
	}

	store.GetInt64(HeightKey(chainID))

	return opb, nil
}

// BuildFISCOChain builds a FISCOChain instance from the given chain params and store
func BuildOpbChain(
	chainParams []byte,
	store *store.Store,
) (*OpbChain, error) {
	var params ChainParams
	err := json.Unmarshal(chainParams, &params)
	if err != nil {
		return nil, err
	}

	baseCfgBz, err := store.Get(BaseConfigKey())
	if err != nil {
		return nil, err
	}

	var baseConfig BaseConfig
	err = json.Unmarshal(baseCfgBz, &baseConfig)
	if err != nil {
		return nil, err
	}

	config := Config{
		BaseConfig:  baseConfig,
		ChainParams: params,
	}

	return NewOpbChain(config, store)
}

// GetChainID implements AppChainI
func (opb *OpbChain) GetChainID() string {
	return opb.ChainID
}

// Start implements AppChainI
func (opb *OpbChain) Start(handler core.InterchainRequestHandler) error {
	if !opb.done {
		return fmt.Errorf("chain %s has been started", opb.ChainID)
	}

	opb.done = false
	opb.handler = handler

	go opb.monitor()

	logging.Logger.Infof("chain %s started", opb.ChainID)

	return nil
}

// Stop implements AppChainI
func (f *OpbChain) Stop() error {
	logging.Logger.Infof("stopping chain %s", f.ChainID)
	f.done = true

	return nil
}

// GetHeight implements AppChainI
func (f *OpbChain) GetHeight() int64 {
	return f.lastHeight
}

// SendResponse implements AppChainI
func (opb *OpbChain) SendResponse(requestID string, response core.ResponseI) error {
	execAbi := wasm.NewContractABI().
		WithMethod("set_response").
		WithArgs("request_id", requestID).
		WithArgs("err_msg", response.GetErrMsg()).
		WithArgs("output", response.GetOutput())
	resultTx, err := opb.OpbClient.WASM.Execute(opb.Config.ChainParams.IServiceCoreAddr, execAbi, nil, opb.BuildBaseTx())
	if err != nil {
		mysql.TxErrCollection(requestID, err.Error())
		return err
	}

	// TODO
	mysql.OnInterchainRequestResponseSent(requestID, resultTx.Hash)

	err = opb.waitForSuccess(resultTx.Hash, "SetResponse")
	if err != nil {
		mysql.TxErrCollection(requestID, err.Error())
		return err
	}

	// TODO
	mysql.OnInterchainRequestSucceeded(requestID)

	return nil
}

// waitForSuccess waits for the receipt of the given tx
func (opb *OpbChain) waitForSuccess(txHash string,name string) error {
	logging.Logger.Infof("%s: transaction sent to %s, hash: %s", name, opb.GetChainID(), txHash)

	tx, _:= opb.OpbClient.QueryTx(txHash)
	if tx.Result.Code != 0{
		return fmt.Errorf("transaction %s execution failed: %s", txHash, tx.Result.Log)
	}

	logging.Logger.Infof("%s: transaction %s execution succeeded", name, txHash)

	return nil
}

// BuildBaseTx builds a base tx
func (opb *OpbChain) BuildBaseTx() types.BaseTx {
	return types.BaseTx{
		From:     opb.Config.KeyName,
		Password: opb.Config.Passphrase,
		Mode:     sdktypes.Commit,
	}
}

// buildInterchainRequest builds an interchain request from the interchain event
func (opb *OpbChain) buildInterchainRequest(e abci.Event) core.InterchainRequest {
	requestID, err := opb.getAttributeValue(e, "request_id")
	if err != nil {
		logging.Logger.Errorf("failed to get requestID: %s", err)
	}
	endpointInfoStr, err := opb.getAttributeValue(e, "endpoint_info")
	if err != nil {
		logging.Logger.Errorf("failed to get endpointInfo: %s", err)
	}
	method, err := opb.getAttributeValue(e, "method")
	if err != nil {
		logging.Logger.Errorf("failed to get method: %s", err)
	}
	callData, err := opb.getAttributeValue(e, "callData")
	if err != nil {
		logging.Logger.Errorf("failed to get callData: %s", err)
	}
	var endpointInfo EndpointInfo
	err = json.Unmarshal([]byte(endpointInfoStr), &endpointInfo)
	if err != nil {
		logging.Logger.Errorf("failed to decode endpointInfo: %s", err)
	}
	return core.InterchainRequest{
		ID:              requestID,
		SourceChainID:   opb.ChainID,
		DestChainID:     endpointInfo.DestChainID,
		DestSubChainID:  endpointInfo.DestSubChainID,
		DestChainType:   endpointInfo.DestChainType,
		EndpointAddress: endpointInfo.EndpointAddress,
		EndpointType:    endpointInfo.EndpointType,
		Method:          method,
		CallData:        []byte(callData),
	}
}

// monitor is responsible for monitoring the chain
func (opb *OpbChain) monitor() {
	for {
		opb.scan()

		if opb.done {
			return
		}

		time.Sleep(time.Duration(opb.Config.MonitorInterval) * time.Second)
	}
}

// scan performs chain scanning
func (opb *OpbChain) scan() {
	currentHeight, err := opb.getBlockNumber()
	if err != nil {
		logging.Logger.Errorf("failed to get the current block height: %s", err)
		return
	}

	if opb.lastHeight == 0 {
		opb.lastHeight = currentHeight - 1
	}

	if currentHeight <= opb.lastHeight {
		return
	}

	opb.scanBlocks(opb.lastHeight+1, currentHeight)
}

// scanBlocks scans the blocks of the specified range
func (opb *OpbChain) scanBlocks(startHeight int64, endHeight int64) {
	for h := startHeight; h <= endHeight; {
		blockResult, err := opb.OpbClient.BlockResults(context.Background(), &h)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			time.Sleep(time.Duration(5) * time.Second)
			continue
		}
		block, err := opb.OpbClient.Block(context.Background(), &h)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			time.Sleep(time.Duration(5) * time.Second)
			continue
		}
		opb.parseCrossChainRequest(blockResult.TxsResults, block)

		err = opb.updateHeight(h)
		if err != nil {
			logging.Logger.Errorf("failed to update height: %s", err)
		}

		h++
	}
}

// getBlockNumber retrieves the current block number
func (opb *OpbChain) getBlockNumber() (int64, error) {
	resultState, err := opb.OpbClient.Status(context.Background())
	if err != nil {
		return -1, err
	}

	return resultState.SyncInfo.LatestBlockHeight,nil
}

// parseServiceInvokedEvents parses the ServiceInvoked events from the receipt
func (opb *OpbChain) parseCrossChainRequest(txResults []*abci.ResponseDeliverTx, block *ctypes.ResultBlock) {
	for i, txResult := range txResults {
		for _, e := range txResult.Events {
			if e.Type == "wasm"{
				contractAddr, _ := opb.getAttributeValue(e, "contract_address")
				if contractAddr == opb.Config.ChainParams.IServiceCoreAddr{
					request := opb.buildInterchainRequest(e)
					opb.handler(opb.ChainID, request, strings.ToUpper(hex.EncodeToString(block.Block.Txs[i].Hash())))
				}
			}
		}
	}
}

func (opb *OpbChain) getAttributeValue(event abci.Event, attributeKey string) (string, error) {
	for _, attr := range event.Attributes {
		if string(attr.Key) == attributeKey {
			return string(attr.Value), nil
		}
	}

	return "", fmt.Errorf("attribute key %s does not exist", attributeKey)
}

// storeChainParams stores the chain params
func (f *OpbChain) storeChainParams() error {
	bz, err := json.Marshal(f.Config.ChainParams)
	if err != nil {
		return err
	}

	return f.store.Set(ChainParamsKey(f.ChainID), bz)
}

func (f *OpbChain) storeChainID() error {
	chainIDsbz, err := f.store.Get([]byte("chainIDs"))
	if err != nil {
		return err
	}
	chainIDs := map[string]string{}
	err = json.Unmarshal(chainIDsbz, &chainIDs)
	if err != nil {
		return err
	}
	chainIDs[f.ChainID] = "opb"
	bz, err := json.Marshal(chainIDs)
	return f.store.Set([]byte("chainIDs"), bz)
}

// updateHeight updates the height
func (f *OpbChain) updateHeight(height int64) error {
	f.lastHeight = height
	return f.store.SetInt64(HeightKey(f.ChainID), height)
}
