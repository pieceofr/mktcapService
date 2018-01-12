package main

import (
	"flag"
	"log"
	"mktcapService/mktcap"
	"time"

	"github.com/golang/glog"
)

var srvConfig ServiceConfig
var mktcapConf mktcap.MktcapConfig

func main() {
	flag.Parse() // for glog
	srvConfig = InitConfig()
	mktcapConf = mktcap.InitConfig(srvConfig.SQLEndpoint, srvConfig.SQLUser, srvConfig.SQLPwd, srvConfig.SQLDB, srvConfig.SQLTickerTable)
	terminate := make(chan int)
	stopSave := make(chan int)
	stopMon := make(chan int)
	log.Println("CoinMarketCap Service has started!")
	if srvConfig.QuickMonitor {
		go MonitorCoinListService(stopMon, srvConfig.QuickMonitorInterval, srvConfig.QuickMonitorLimit)
	}
	if srvConfig.SaveToDB {
		go SaveDBRoutine(stopSave, srvConfig.SaveToDBInterval, srvConfig.SaveToDBLimit)
	}
	go apiService(srvConfig.ApiPort)
	<-terminate
}

/*MonitorCoinListService : monitoring coin list in every srvConfig.QuickMonitorInterval sec*/
func MonitorCoinListService(stop <-chan int, interval, numrecords int) {
	glog.Info("MonitorCoinListService has started!")
	firstMon := true
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	coinList := InitMonitorList()
	dataChan := make(chan map[string]CoinMonitor)
	idListChan := make(chan []string)
	go QuickAnalysisService(idListChan, dataChan) //Start QuickAnalysisService Only if MonitorService is Start

	for {
		idList := make([]string, 0, numrecords) // for store idList to analyzed
		if firstMon {                           // First time don't wait for srvConf.SaveToDBInterval Time
			firstMon = false
			monCoins, err := mktcap.TickerNow(0, numrecords, mktcapConf)
			if err != nil {
				log.Println(err)
			}
			for _, val := range monCoins {
				AddNewToList(val, coinList)
				if srvConfig.MonitorType != "assign" {
					idList = append(idList, val.ID)
				}

			}
			continue
		}
		select {
		case <-ticker.C:
			monCoins, err := mktcap.TickerNow(0, numrecords, mktcapConf)
			if err != nil {
				log.Println(err)
			}
			for _, val := range monCoins {
				if srvConfig.MonitorType != "assign" {
					idList = append(idList, val.ID)
				}
				AddNewToList(val, coinList)
			}
			if srvConfig.MonitorType == "assign" {
				idList = srvConfig.MonitorCoinList
			}
			dataChan <- coinList
			idListChan <- idList
		case <-stop:
			ticker.Stop()
		}
	}
}

/*
SaveDBRoutine : save coin data to db in every interval seconds
*/
func SaveDBRoutine(stop <-chan int, interval, numrecords int) {
	glog.Info("SaveDBRoutine has started!")
	firstsave := true
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	for {

		if firstsave { // First time don't wait for srvConf.SaveToDBInterval Time
			firstsave = false
			err := mktcap.TickerSave(0, numrecords, mktcapConf)
			if err != nil {
				log.Println(err)
			}
			continue
		}

		select {
		case <-ticker.C:
			err := mktcap.TickerSave(0, numrecords, mktcapConf)
			if err != nil {
				log.Println(err)
			}

		case <-stop:
			break
		}
	}
}
