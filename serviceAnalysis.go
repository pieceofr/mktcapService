package main

import (
	"github.com/glog"
)

/*
QuickAnalysisService : is a data analysis service
*/
func QuickAnalysisService(idList <-chan []string, dataIn <-chan map[string]CoinMonitor) {
	glog.Info("QuickAnalysisService has started")
	slackMsg := make(chan string)
	slackStop := make(chan bool)
	go slackSendService(slackMsg, slackStop)
	for {
		var newData map[string]CoinMonitor
		for {
			select {
			case newData = <-dataIn:
			case newList := <-idList:
				//glog.V(3).Infoln("QuickAnalysisService Recieve idList", newList)
				//glog.V(3).Infoln("QuickAnalysisService Recieve newData", newData)
				newDataSet := GetCoinMonByIDList(newList, newData)
				r1DiffSig := rLastDiffSig(srvConfig.RuleSigDiffObserv, srvConfig.RuleSigDiffThreadhold, newDataSet)
				if len(r1DiffSig) > 0 {
					slackMsg <- getAllSigDiffMessage(r1DiffSig)
				}
			}
		}
	}

}
