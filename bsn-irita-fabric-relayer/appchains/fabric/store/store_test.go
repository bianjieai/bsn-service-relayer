package store

import (
	"fmt"
	"relayer/appchains/fabric/entity"
	"testing"
	"time"
)

func TestStoreRelayer(t *testing.T) {

	db := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"

	InitMysql(db)

	relayer := entity.FabricRelayer{
		ChainBase:      entity.ChainBase{ChainId: 1},
		AppCode:        "123",
		ChannelId:      "8899",
		CrossChainCode: "123",
		NodeName:       "beijing",
		CityNode:       "org12334",
		Status:         1,
		CreateTime:     time.Now(),
		LastUpdateTime: time.Now(),
	}

	err := StoreRelayerAppInfo(&relayer)
	if err != nil {
		t.Error(err)
	}

}

func TestGetRelayerAppInfos(t *testing.T) {

	db := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"

	InitMysql(db)

	queryResult, err := GetRelayerAppInfos()

	if err != nil {
		t.Errorf(err.Error())
	}

	for i := range queryResult {
		fmt.Println(queryResult[i])
	}

}

func TestUpdateRelayerAppInfo(t *testing.T) {

	db := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"

	InitMysql(db)

	data := entity.FabricRelayer{
		ChainBase:      entity.ChainBase{ChainId: 1},
		ChannelId:      "chan8899",
		CrossChainCode: "cross12345678",
		NodeName:       "orgbeijing123",
	}

	err := UpdateRelayerAppInfo(&data)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestDeleteRelayerAppInfo(t *testing.T) {

	db := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"

	InitMysql(db)

	err := DeleteRelayerAppInfo("2")

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestStoreRelayerTxReqInfo(t *testing.T) {
	db := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"

	InitMysql(db)

	relayerTx := entity.FabricRelayerTx{
		Request_id:    "123",
		From_chainid:  "456",
		From_tx:       "asd",
		Hub_req_tx:    "wqe",
		To_chainid:    "sadwq",
		To_tx:         "das",
		Hub_res_tx:    "sada",
		From_res_tx:   "sad",
		Tx_status:     1,
		Tx_time:       time.Now(),
		Tx_createtime: time.Now(),
	}

	err := StoreRelayerTxReqInfo(&relayerTx)
	if err != nil {
		t.Error(err)
	}
}

func TestStoreRelayerTxResInfo(t *testing.T) {
	db := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"

	InitMysql(db)

	relayerTx := entity.FabricRelayerTx{
		Request_id:    "123",
		From_chainid:  "aa",
		From_tx:       "wq",
		Hub_req_tx:    "es",
		To_chainid:    "s",
		To_tx:         "dasw",
		Hub_res_tx:    "sadda",
		From_res_tx:   "ad",
		Tx_status:     2,
		Tx_time:       time.Now(),
		Tx_createtime: time.Now(),
	}

	err := StoreRelayerTxResInfo(&relayerTx)

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestInsertInterchainRequestInfo(t *testing.T) {

	db := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"

	InitMysql(db)

	relayerTx := entity.FabricRelayerTx{
		Request_id:    "1995",
		From_chainid:  "11",
		From_tx:       "rong",
		Tx_createtime: time.Now(),
	}

	InsertInterchainRequestInfo(&relayerTx)
}

func TestSendHUBRequestInfo(t *testing.T) {

	db := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"

	InitMysql(db)

	relayerTx := entity.FabricRelayerTx{
		Request_id:    "1995",
		Hub_req_tx:    "11",
		Ic_request_id: "rong",
	}

	SendHUBRequestInfo(&relayerTx)

}

func TestCallBackSendResponse(t *testing.T) {

	db := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"

	InitMysql(db)

	relayerTx := entity.FabricRelayerTx{
		Request_id:    "1995",
		Ic_request_id: "rong",
		From_res_tx:   "222",
		Tx_time:       time.Now(),
	}

	CallBackSendResponse(&relayerTx)

}
