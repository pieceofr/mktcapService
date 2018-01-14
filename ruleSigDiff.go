package main

import (
	"fmt"
	"mktcapService/mktcap"
)

//SigDiffCond condition used in Sigdiff caculation
type SigDiffCond struct {
	Threadhold   float64
	Observations int
}

/*SigDiffRet Use For pLastChangeDiff return*/
type SigDiffRet struct {
	ID           string
	Name         string
	Observations int
	Threadhold   float64
	BTCPrice     float64
	USDPrice     float64
	DiffBtc      float64
	DiffUsd      float64
	DiffHour     float64
}

// IsSigDiff is any indicator has significant change
func (l SigDiffRet) IsSigDiff() bool {
	if l.DiffBtc != 0 || l.DiffUsd != 0 || l.DiffHour != 0 {
		return true
	}
	return false
}

func getAllSigDiffMessage(changes []SigDiffRet) string {
	if len(changes) <= 0 {
		return ""
	}
	message := fmt.Sprintf("SigDiff Report \n --- INC: thrhold:%.2f obvs:%d ---\n", changes[0].Threadhold, changes[0].Observations)
	negMessage := ""
	for _, val := range changes {
		neg, msg := val.getSigDiffMessage()
		if neg {
			negMessage = fmt.Sprintf("%s%s\n", negMessage, msg)
		} else {
			message = fmt.Sprintf("%s%s\n", message, msg)
		}
	}
	message = fmt.Sprintf("\n--- DEC: thrhold:%.2f obvs:%d ---\n%s", changes[0].Threadhold, changes[0].Observations, negMessage)
	return message
}

func (l SigDiffRet) getSigDiffMessage() (neg bool, message string) {
	message = fmt.Sprintf("%s BTC:%.8f USD:%.2f ", l.ID, l.BTCPrice, l.USDPrice)
	if l.DiffBtc != 0 {
		message = fmt.Sprintf("%s BTC(%.2f%s)", message, l.DiffBtc, "%")
	}
	if l.DiffUsd != 0 {
		message = fmt.Sprintf("%s USD(%.2f%s)", message, l.DiffUsd, "%")
	}
	if l.DiffHour != 0 {
		message = fmt.Sprintf("%s 1Hr(%.2f%s)", message, l.DiffHour, "%")
	}
	if l.DiffBtc != 0 && l.DiffBtc < 0 {
		neg = true
	} else if l.DiffBtc == 0 && l.DiffUsd != 0 && l.DiffUsd < 0 { //depend on DiffUsd
		neg = true
	} else if l.DiffBtc == 0 && l.DiffUsd == 0 && l.DiffHour != 0 && l.DiffHour < 0 {
		neg = true
	}
	return neg, message
}

func rLastDiffSig(observations int, threadhold float64, data []CoinMonitor) []SigDiffRet {
	ret := make([]SigDiffRet, 0, 100)
	for _, val := range data {
		newchange := SigDiffRet{}
		newchange.ID = val.ID
		newchange.Name = val.Name
		newchange.Observations = observations
		newchange.Threadhold = threadhold
		dnew := mktcap.MktCapInfo{}
		dold := mktcap.MktCapInfo{}
		if len(val.PeriodData) <= 1 {
			return ret
		} else if len(val.PeriodData) > 1 && len(val.PeriodData) >= observations {
			dnew = val.PeriodData[len(val.PeriodData)-1]
			dold = val.PeriodData[len(val.PeriodData)-observations]
		} else if len(val.PeriodData) > 1 {
			dnew = val.PeriodData[len(val.PeriodData)-1]
			dold = val.PeriodData[0]
		}
		newchange.BTCPrice = dnew.PriceBtc
		newchange.USDPrice = dnew.PriceUsd
		newPercent := ((dnew.PriceBtc - dold.PriceBtc) / dold.PriceBtc) * 100
		if newPercent >= threadhold || newPercent < (0-threadhold) {
			newchange.DiffBtc = newPercent
		}
		newPercent = ((dnew.PriceUsd - dold.PriceUsd) / dold.PriceUsd) * 100
		if newPercent >= threadhold || newPercent < (0-threadhold) {
			newchange.DiffUsd = newPercent
		}
		newPercent = dnew.PercentChange1h - dold.PercentChange1h
		if newPercent >= threadhold || newPercent < (0-threadhold) {
			newchange.DiffHour = newPercent
		}
		if newchange.IsSigDiff() {
			ret = append(ret, newchange)
		}
	}
	return ret
}
