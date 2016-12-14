package accountclient

import (
	"errors"
	"golang-project/dpsg/common"
	"golang-project/dpsg/logger"
	"golang-project/dpsg/proto"
	"golang-project/dpsg/rpcplusclientpool"
)

var pPoll *rpcplusclientpool.ClientPool

var FailedError error = errors.New("operate failed, please try again !")

func Init() {
	var dbcfg common.DBConfig
	if err := common.ReadDbConfig("accountserver.json", &dbcfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
		return
	}

	aHosts := make([]string, 0)
	aHosts = append(aHosts, dbcfg.DBHost)
	pPoll = rpcplusclientpool.CreateClientPool(aHosts)
	if pPoll == nil {
		logger.Fatal("create failed")
	}
}

//关联
func SetPartnerIdToPlayerId(partnername, partnerid, playerid string) error {
	try := &proto.AccountDbWrite{Table: partnername, Key: partnerid, Value: playerid}
	rst := &proto.AccountDbWriteResult{}

	err, conn := pPoll.RandomGetConn()
	if err != nil {
		return err
	}

	if err = conn.Call("AccountServer.Write", try, rst); err != nil {
		return err
	}

	if rst.Code != proto.Ok {
		return FailedError
	}

	return nil
}

//删除关联
func DelPartnerIdToPlayerId(partnername, partnerid string) error {
	try := &proto.AccountDbQuery{Table: partnername, Key: partnerid}
	rst := &proto.AccountDbQueryResult{}

	err, conn := pPoll.RandomGetConn()
	if err != nil {
		return err
	}

	if err = conn.Call("AccountServer.Delete", try, rst); err != nil {
		return nil
	}

	if rst.Code != proto.Ok {
		return FailedError
	}

	return nil
}

//查询关联的玩家id
func QueryPlayerIdByPartnerId(partnername, partnerid string) (playerid string, err error) {
	try := &proto.AccountDbQuery{Table: partnername, Key: partnerid}
	rst := &proto.AccountDbQueryResult{}

	err, conn := pPoll.RandomGetConn()
	if err != nil {
		return
	}

	if err = conn.Call("AccountServer.Query", try, rst); err != nil {
		return
	}

	//查询要关心存不存在
	if rst.Code == proto.NoExist {
		playerid = ""
		return
	}

	if rst.Code != proto.Ok {
		err = FailedError
		return
	}

	playerid = rst.Value
	err = nil

	return
}

//反向查询关联的玩家id
func QueryPartnerIdByPlayerId(partnername, playerid string) (partnerid string, err error) {
	try := &proto.AccountDbQuery{Table: partnername, Key: playerid}
	rst := &proto.AccountDbQueryResult{}

	err, conn := pPoll.RandomGetConn()
	if err != nil {
		return
	}

	if err = conn.Call("AccountServer.ReQuery", try, rst); err != nil {
		return
	}

	//查询要关心存不存在
	if rst.Code == proto.NoExist {
		playerid = ""
		return
	}

	if rst.Code != proto.Ok {
		err = FailedError
		return
	}

	partnerid = rst.Value
	err = nil

	return
}
