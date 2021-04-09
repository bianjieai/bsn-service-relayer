package mysql

import "relayer/logging"

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
}

// OnInterchainRequestSent is the hook which is called when the interchain request is sent to hub
func OnInterchainRequestSent(requestID, hubReqTx, toChainID string) {
	err := update("hub_req_tx", hubReqTx, requestID)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return
	}

	err = update("to_chainid", toChainID, requestID)
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
func OnInterchainRequestResponseSent() {
	// TODO
}

// OnInterchainRequestSucceeded is the hook which is called when the response to the interchain request is handled successfully
func OnInterchainRequestSucceeded() {
	// TODO
}
