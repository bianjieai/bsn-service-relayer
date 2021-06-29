package store

import (
	"database/sql"
	"fmt"
	"relayer/appchains/fabric/entity"
	"relayer/common/mysql"
	"relayer/errors"
	"relayer/logging"
	"time"
)

func StoreRelayerAppInfo(data *entity.FabricRelayer) error {

	logging.Logger.Infoln("Into StoreRelayerAppInfo method：")

	logging.Logger.Infof("start into ChainId：%d\n", data.ChainId)

	insertSql := fmt.Sprintf("INSERT INTO %s (ChainId,AppCode,ChannelId,CrossChainCode,NodeName,CityCode,STATUS,CreateTime,LastUpdateTime) VALUES (?,?,?,?,?,?,?,?,?);", _TabName_Relayer)

	data.Status = 1
	creatTime := data.CreateTime.Format("2006-01-02 15:04:05")
	lastUpdateTime := data.LastUpdateTime.Format("2006-01-02 15:04:05")

	lastID, rows, err := mysql.Exec(
		insertSql,
		data.GetChainId(),
		data.AppCode,
		data.ChannelId,
		data.CrossChainCode,
		data.NodeName,
		data.CityNode,
		data.Status,
		creatTime,
		lastUpdateTime)

	if err != nil {
		logging.Logger.Errorf("save ChainId：%d err!：%s\n", data.ChainId, err.Error())
		return err
	}

	logging.Logger.Infof("save success ChainId：%d\n", data.ChainId)
	logging.Logger.Infof("save lastID：%d, Number of rows affected：%d\n", lastID, rows)

	return nil
}

func GetRelayerAppInfos() ([]*entity.FabricRelayer, error) {

	logging.Logger.Infoln("Into GetRelayerAppInfo method：")

	querySql := fmt.Sprintf(`
				SELECT
					id,
					ChainId,
					AppCode,
					ChannelId,
					CrossChainCode,
					NodeName,
					CityCode,
					STATUS,
					CreateTime,
					LastUpdateTime
				FROM
					%s
				WHERE
					STATUS = 1;`, _TabName_Relayer)

	queryData := func(rows *sql.Rows) (interface{}, error) {

		var createTime string
		var lastUpdateTime string
		var chainID string

		data := &entity.FabricRelayer{}

		rows.Scan(
			&data.ID,
			&chainID,
			&data.AppCode,
			&data.ChannelId,
			&data.CrossChainCode,
			&data.NodeName,
			&data.CityNode,
			&data.Status,
			&createTime,
			&lastUpdateTime)

		err := data.ChainBase.SetChainId(chainID)
		if err != nil {
			return nil, errors.New("ChainID type conversion failure")
		}

		// Time format
		DefaultTimeLoc := time.Local
		c, _ := time.ParseInLocation("2006-01-02 15:04:05", createTime, DefaultTimeLoc)
		l, _ := time.ParseInLocation("2006-01-02 15:04:05", lastUpdateTime, DefaultTimeLoc)
		data.CreateTime = c
		data.LastUpdateTime = l

		return data, nil
	}

	logging.Logger.Infoln("start query RelayerAppInfoList：")

	RelayerAppInfoList, err := mysql.Query(queryData, querySql)
	if err != nil {
		logging.Logger.Errorf("query err：%s\n", err.Error())
		return nil, err
	}

	dataList := []*entity.FabricRelayer{}
	for i, _ := range RelayerAppInfoList {
		dataList = append(dataList, RelayerAppInfoList[i].(*entity.FabricRelayer))
	}

	logging.Logger.Infof("query success, return count：%d\n", len(dataList))

	return dataList, nil
}

func UpdateRelayerAppInfo(data *entity.FabricRelayer) error {
	logging.Logger.Infoln("Into UpdateRelayerAppInfo method：")

	updateSql := fmt.Sprintf(`UPDATE %s SET ChannelId=?,CrossChainCode=?,NodeName =? WHERE ChainId = ?;`, _TabName_Relayer)
	lastId, rows, err := mysql.Exec(updateSql, data.ChannelId, data.CrossChainCode, data.NodeName, data.GetChainId())
	if err != nil {
		logging.Logger.Errorf("UPDATE err: %s", err.Error())
		return err
	}

	logging.Logger.Infof("UPDATE success ChainId：%d\n", data.ChainId)
	logging.Logger.Infof("UPDATE lastID：%d, Number of rows affected：%d\n", lastId, rows)

	return nil
}

func DeleteRelayerAppInfo(chainId string) error {

	logging.Logger.Infoln("Into DeleteRelayerAppInfo method：")

	deleteSql := fmt.Sprintf("DELETE FROM %s WHERE ChainId = %s;", _TabName_Relayer, chainId)

	lastId, rows, err := mysql.Exec(deleteSql)
	if err != nil {
		logging.Logger.Errorf("DELETE err: %s", err.Error())
		return err
	}

	logging.Logger.Infof("DELETE success ChainId：%s\n", chainId)
	logging.Logger.Infof("DELETE lastID：%d, Number of rows affected：%d\n", lastId, rows)

	return nil
}

func StoreRelayerTxReqInfo(date *entity.FabricRelayerTx) error {

	logging.Logger.Infoln("Into StoreRelayerAppInfo method：")

	insterSql := fmt.Sprintf(`INSERT INTO %s
		(Request_id,From_chainid,From_tx,Hub_req_tx,To_chainid,To_tx,Hub_res_tx,From_res_tx,Tx_status,Tx_time,Tx_createtime)
		VALUES (?,?,?,?,?,?,?,?,?,?,?);`, _TabName_cc_Tx)
	Tx_time := date.Tx_time.Format("2006-01-02 15:04:05")
	Tx_createtime := date.Tx_createtime.Format("2006-01-02 15:04:05")
	lastID, rows, err := mysql.Exec(
		insterSql,
		date.Request_id,
		date.From_chainid,
		date.From_tx,
		date.Hub_req_tx,
		date.To_chainid,
		date.To_tx,
		date.Hub_res_tx,
		date.From_res_tx,
		date.Tx_status,
		Tx_time,
		Tx_createtime,
	)
	if err != nil {
		fmt.Printf("StoreRelayerTxReqInfo failed, err:%v\n", err)
		return err
	}

	logging.Logger.Infof("StoreRelayerTxReqInfo success Request_id：%s\n", date.Request_id)
	logging.Logger.Infof("StoreRelayerTxReqInfo lastID：%d, Number of rows affected：%d\n", lastID, rows)
	return nil
}

func StoreRelayerTxResInfo(date *entity.FabricRelayerTx) error {
	updateStr := fmt.Sprintf(`UPDATE %s
		SET From_chainid=?,From_tx=?,Hub_req_tx=?,To_chainid=?,To_tx=?,Hub_res_tx=?,From_res_tx=?,Tx_status=?,Tx_time=?,Tx_createtime=?
		WHERE Request_id =?`, _TabName_cc_Tx)
	Tx_time := date.Tx_time.Format("2006-01-02 15:04:05")
	Tx_createtime := date.Tx_createtime.Format("2006-01-02 15:04:05")
	lastID, rows, err := mysql.Exec(
		updateStr,
		date.From_chainid,
		date.From_tx,
		date.Hub_req_tx,
		date.To_chainid,
		date.To_tx,
		date.Hub_res_tx,
		date.From_res_tx,
		date.Tx_status,
		Tx_time,
		Tx_createtime,
		date.Request_id,
	)
	if err != nil {
		logging.Logger.Errorf("StoreRelayerTxResInfo err: %s", err.Error())
		return err
	}

	logging.Logger.Infof("StoreRelayerTxResInfo success Request_id：%s\n", date.Request_id)
	logging.Logger.Infof("StoreRelayerTxResInfo lastID：%d, Number of rows affected：%d\n", lastID, rows)
	return nil
}
