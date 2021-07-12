package mysql

import (
	"relayer/logging"
)

// OnInterchainRequestReceived is the hook which is called when the interchain request is received
func OnInterchainRequestReceived(requestID, fromChainID, fromTx string) {
	err := insert("request_id", requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}

	err = update("from_chainid", fromChainID, requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}

	err = update("from_tx", fromTx, requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}

	err = updateTime("tx_createtime", requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}
}

// OnInterchainRequestSent is the hook which is called when the interchain request is sent to hub
func OnInterchainRequestSent(requestID, icResID, hubReqTx string) {
	err := update("hub_req_tx", hubReqTx, requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}

	err = update("ic_request_id", icResID, requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}
}

// OnInterchainRequestHandled is the hook which is called when the interchain request is handled by hub
func OnInterchainRequestHandled() {
	// TODO
}

// OnInterchainRequestResponseSent is the hook which is called when the response to the interchain request is sent to the app chain
func OnInterchainRequestResponseSent(requestID, fromResTx string) {
	err := update("from_res_tx", fromResTx, requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}
}

// OnInterchainRequestSucceeded is the hook which is called when the response to the interchain request is handled successfully
func OnInterchainRequestSucceeded(requestID string) {
	err := update("tx_status", "1", requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}

	err = updateTime("tx_time", requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}
}

func TxErrCollection(requestID, errStr string) {
	err := update("error", errStr, requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}
	err = update("tx_status", "2", requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}
}
