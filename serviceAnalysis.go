package main

import (
	"github.com/golang/glog"
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
				go sigDiffRoutine([]SigDiffCond{}, newList, newData, slackMsg)
			}
		}
	}
}

func sigDiffRoutine(conditions []SigDiffCond, newList []string, newData map[string]CoinMonitor, toslack chan string) {
	//glog.V(3).Infoln("QuickAnalysisService Recieve idList", newList)
	//glog.V(3).Infoln("QuickAnalysisService Recieve newData", newData)
	newDataSet := GetCoinMonByIDList(newList, newData)
	for _, cond := range srvConfig.SigDiffConditions {
		r1DiffSig := rLastDiffSig(cond.Observations, cond.Threadhold, newDataSet)
		if len(r1DiffSig) > 0 {
			toslack <- getAllSigDiffMessage(r1DiffSig)
		}
	}

}
