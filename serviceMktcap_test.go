package main

import (
	"mktcapService/mktcap"
	"testing"

	"github.com/glog"
)

func TestAddNewCoin(t *testing.T) {
	srvConfig := InitConfig()
	mktcapConf := mktcap.InitConfig(srvConfig.SQLEndpoint, srvConfig.SQLUser, srvConfig.SQLPwd, srvConfig.SQLDB, srvConfig.SQLTickerTable)
	coinList := InitMonitorList()
	mkts, err := mktcap.TickerNow(0, 5, mktcapConf)
	if err != nil {
		glog.Fatalln(err)
	}
	for _, val := range mkts {
		AddNewToList(val, coinList)
	}
	PrintCoinMonitorList(coinList)
	mktdata := GetCoinMonitorByID(mkts[0].ID, coinList)
	mktdata.PrintMonitorData()

}
func TestGetListByIDList(t *testing.T) {
	srvConfig := InitConfig()
	mktcapConf := mktcap.InitConfig(srvConfig.SQLEndpoint, srvConfig.SQLUser, srvConfig.SQLPwd, srvConfig.SQLDB, srvConfig.SQLTickerTable)
	coinList := InitMonitorList()

	mkts, err := mktcap.TickerNow(0, 5, mktcapConf)
	if err != nil {
		glog.Fatalln("Can to get any coin records")
	}
	for _, val := range mkts {
		AddNewToList(val, coinList)
	}
	mktList := GetCoinMonByIDList([]string{"ethereum", "cardano"}, coinList)

	for _, val := range mktList {
		glog.V(3).Infoln("ID:", val.ID, "Name:", val.Name, "  len:", len(val.PeriodData), " data:", val.PeriodData)
	}

}
