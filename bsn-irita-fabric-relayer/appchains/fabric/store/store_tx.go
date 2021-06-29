package store

import (
	"fmt"
	"relayer/appchains/fabric/entity"
	"relayer/common/mysql"
	"relayer/logging"
	"time"
)

func InsertInterchainRequestInfo(date *entity.FabricRelayerTx) {

	requestId := date.Request_id
	fromChainId := date.From_chainid
	fromTx := date.From_tx
	txCreatetime := date.Tx_createtime.Format("2006-01-02 15:04:05")
	tx_status := date.Tx_status
	source_service := date.Source_service

	insertsql := fmt.Sprintf("INSERT INTO %s "+
		"( request_id, from_chainId, from_tx, tx_createtime, tx_status, source_service ) "+
		"VALUES ( ?, ?, ?, ?, ?, ?);", _TabName_cc_Tx)

	lastId, rows, err := mysql.Exec(insertsql,
		requestId,
		fromChainId,
		fromTx,
		txCreatetime,
		tx_status,
		source_service)
	if err != nil {
		logging.Logger.Errorf("Store InsertInterchainRequestInfo Failed :%s", err.Error())
	} else {
		logging.Logger.Infof("lastId:%d ;rows:%d ", lastId, rows)
	}
}

//todo 根据requestId 设置 错误

func SetErrorTransRecord(requestId, errMsg string) {
	logging.Logger.Infof("SetErrorTransRecord requestId:%s,errMsg:%s", requestId, errMsg)
	sql := `update tb_irita_crosschain_tx set error = ? ,tx_time =? ,tx_status = 2 where request_id = ? and source_service = 0`
	txtime := time.Now().Format("2006-01-02 15:04:05")
	lastId, rows, err := mysql.Exec(sql,
		errMsg,
		txtime,
		requestId)

	if err != nil {
		logging.Logger.Errorf("Store SetErrorTransRecord Failed :%s", err.Error())
	} else {
		logging.Logger.Infof("lastId:%d ;rows:%d ", lastId, rows)
	}

}

func SendHUBRequestInfo(date *entity.FabricRelayerTx) {

	request_id := date.Request_id
	hub_req_tx := date.Hub_req_tx
	ic_request_id := date.Ic_request_id
	source_service := date.Source_service

	updatesql := fmt.Sprintf("UPDATE %s SET hub_req_tx=?,ic_request_id=? WHERE request_id = ? And source_service= ?;", _TabName_cc_Tx)

	lastId, rows, err := mysql.Exec(updatesql,
		hub_req_tx,
		ic_request_id,
		request_id,
		source_service)
	if err != nil {
		logging.Logger.Errorf("update SendHUBRequestInfo Failed :%s", err.Error())
	}
	fmt.Printf("lastId:%d \n", lastId)
	fmt.Printf("rows:%d \n", rows)
}

func CallBackSendResponse(date *entity.FabricRelayerTx) {

	request_id := date.Request_id
	ic_request_id := date.Ic_request_id
	from_res_tx := date.From_res_tx
	errmsg := date.Error
	tx_time := date.Tx_time.Format("2006-01-02 15:04:05")
	tx_status := date.Tx_status
	source_service := date.Source_service

	updatesql := fmt.Sprintf("UPDATE %s SET ic_request_id=?,from_res_tx=?,tx_time = ?, tx_status = ?,error = ? WHERE request_id = ? and source_service = ?;", _TabName_cc_Tx)

	lastId, rows, err := mysql.Exec(updatesql,
		ic_request_id,
		from_res_tx,
		tx_time,
		tx_status,
		errmsg,
		request_id,
		source_service)

	if err != nil {
		logging.Logger.Errorf("update CallBackSendResponse Failed :%s", err.Error())
	}
	fmt.Printf("lastId:%d \n", lastId)
	fmt.Printf("rows:%d \n", rows)

}
