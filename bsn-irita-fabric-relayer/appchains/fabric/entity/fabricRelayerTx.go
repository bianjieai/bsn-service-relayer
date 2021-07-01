package entity

import (
	"time"
)

type FabricRelayerTx struct {
	Funique_id     uint64
	Request_id     string
	From_chainid   string
	From_tx        string
	Hub_req_tx     string
	Ic_request_id  string
	To_chainid     string
	To_tx          string
	Hub_res_tx     string
	From_res_tx    string
	Tx_status      int
	Error          string
	Source_service int
	Tx_time        time.Time
	Tx_createtime  time.Time
}

//Empty method
func (f *FabricRelayerTx) ReqCrossTx() {

}

//Empty method
func (f *FabricRelayerTx) ResCrossTx() {

}
