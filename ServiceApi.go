package main

import (
	"fmt"
	"mktcapService/mktcap"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func apiService(port int) {
	apiPort := ":" + strconv.Itoa(port)
	router := mux.NewRouter()
	router.HandleFunc("/coin/{id}", GetCoinByID).Methods("GET")
	glog.Infoln("apiService is going to started!")
	glog.Errorln(http.ListenAndServe(apiPort, router))
}

/*GetCoinByID Handle return a specfic coin information*/
func GetCoinByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	coinID := params["id"]
	ids := append([]string{}, coinID)
	coin, err := mktcap.TickerNowByIDList(ids, mktcapConf)
	if err != nil {
		glog.Errorln(err)
		fmt.Fprintf(w, err.Error())

	}
	if len(coin) == 1 {
		msg := fmt.Sprintf("%s (btc:%.8f)(usd:%.2f) (1hChg:%.2f)(24hrChg:%.2f)(7dayChg:%.2f)(lastupdate:%s)",
			coin[0].Name, coin[0].PriceBtc, coin[0].PriceUsd,
			coin[0].PercentChange1h, coin[0].PercentChange24h,
			coin[0].PercentChange7d, time.Unix(coin[0].LastUpdated, 0))
		fmt.Fprintf(w, msg)
	}
	fmt.Fprintf(w, "please check data retrieve")

}
