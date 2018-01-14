package main

import (
	"fmt"
	"mktcapService/mktcap"
	"time"
)

//NumOfRecs :number of records to be stored in CoinMonitor Struct
const NumOfRecs = 2000

//EstimateListCoins : number of estimated coins in coinMonitor, this is used for pre-allocate coin
const EstimateListCoins = 1000

var coinMonitorList map[string]CoinMonitor

/*
CoinMonitor : is used for store time squence data of a specific coin
ID : coin ID
Name: coin Name
Rank: Coin Rank
LastedUpdated : LastedUpdated of a coin record
*/
type CoinMonitor struct {
	ID           string
	Name         string
	Rank         int
	LastedUpadte int64
	PeriodData   []mktcap.MktCapInfo
}

//InitMonitorList : Initialized a new coin list to be monitored
func InitMonitorList() map[string]CoinMonitor {
	coinMonitorList = make(map[string]CoinMonitor)
	return coinMonitorList
}

//AddNewToList : Add a New Coin to be monitored in a specific Monitored list
func AddNewToList(coin mktcap.MktCapInfo, m map[string]CoinMonitor) {
	oldCoinMonitor, exist := m[coin.ID]
	if !exist {
		newCoinMonitor := CoinMonitor{ID: coin.ID, Name: coin.Name, LastedUpadte: coin.LastUpdated}
		newCoinMonitor.NewData(coin)
		m[coin.ID] = newCoinMonitor
	} else {
		oldCoinMonitor.NewData(coin)
		m[coin.ID] = oldCoinMonitor // need to reassign value since oldCoin is a copy of value
	}
}

//GetCoinMonitorByID : Get a specific CoinMonitor data by coin ID
func GetCoinMonitorByID(id string, m map[string]CoinMonitor) *CoinMonitor {
	coinMon, exist := m[id]
	if !exist {
		return nil
	}
	return &coinMon
}

/*GetCoinMonByIDList : specify a list of coin ID and return a corresponding list of CoinMonitor*/
func GetCoinMonByIDList(idlist []string, m map[string]CoinMonitor) []CoinMonitor {
	allcoin := make([]CoinMonitor, 0, 3000)
	for _, id := range idlist {
		for key, val := range m {
			if key == id {
				allcoin = append(allcoin, val)
			}
		}
	}
	return allcoin
}

//PrintCoinMonitorList : Printout key and ID of a specific monitoring list
func PrintCoinMonitorList(list map[string]CoinMonitor) {
	for key, val := range list {
		fmt.Print("(", key, ":", val.LastedUpadte, ") ")
	}
	fmt.Println("\n", time.Now().Unix())
}

//NewData : add a new data to specific coin
func (c *CoinMonitor) NewData(data mktcap.MktCapInfo) {
	c.LastedUpadte = data.LastUpdated
	if len(c.PeriodData) < NumOfRecs {
		c.PeriodData = append(c.PeriodData, data)
	} else { // Greater than NumOfRecs keep 10% of data of data and add new
		keep := int(NumOfRecs * 0.1)
		if len(c.PeriodData)-keep > 0 {
			c.PeriodData = append(c.PeriodData[(len(c.PeriodData)-keep):len(c.PeriodData)], data)
		} else {
			c.PeriodData = append(c.PeriodData[0:len(c.PeriodData)], data)
		}
	}
}

//PrintMonitorData : Print out all data of a CoinMonitor
func (c *CoinMonitor) PrintMonitorData() {
	for _, val := range c.PeriodData {
		fmt.Print("(", val, ") ", "\n")
	}
	fmt.Println("PrintMonitorData End LastUpdated:", c.LastedUpadte)
}
