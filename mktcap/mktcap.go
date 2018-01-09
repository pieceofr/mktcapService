package mktcap

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/glog"
)

//ChangeEndPoint  endpoint format: "api.coinmarketcap.com"
func ChangeEndPoint(config MktcapConfig, endpoint string) {
	config.MktCapEndpoint = endpoint
}

//ChangeVer format: "v1"
func ChangeVer(config MktcapConfig, ver string) {
	config.MktAPIVer = ver
}

//TickerNow get Ticker data
func TickerNow(start, limit int, config MktcapConfig) ([]MktCapInfo, error) {
	resp, err := http.Get(getEndPoint(start, limit, config.MktAPITicker, config))
	if err != nil {
		glog.Errorln(getEndPoint(start, limit, config.MktAPITicker, config))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []MktCapInfo{}, err
	}
	jsondata, err := getRawData(body)

	if err != nil {
		return []MktCapInfo{}, err
	}
	list := DataToMktCapInfo(jsondata)
	return list, nil
}

//TickerSave save ticker datat to sql database
func TickerSave(start, limit int, config MktcapConfig) error {
	resp, err := http.Get(getEndPoint(start, limit, config.MktAPITicker, config))
	if err != nil {
		glog.Errorln(getEndPoint(start, limit, config.MktAPITicker, config))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	jsondata, err := getRawData(body)

	if err != nil {
		return err
	}
	db, err := SqlConnect(config.Sqluser, config.Sqlpwd, config.Sqlurl, config.Sqldb)
	defer SqlDisconnect(db)
	if err != nil {
		glog.Errorln(err)
	}
	_, err = insertMultiple(db, config.sqlTickerTable, jsondata)
	if err != nil {
		glog.Errorln(err)
		return err
	}
	return nil
}

// Convert Recieve Raw Data to json format
func getRawData(data []byte) ([]interface{}, error) {
	var f []interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {
		return f, err
	}
	return f, nil
}

//consist endpoint for http
func getEndPoint(start, limit int, apigroup string, config MktcapConfig) string {
	endpoint := "https://" + config.MktCapEndpoint + "/" + config.MktAPIVer + "/" + apigroup

	if start > 0 || limit > 0 {
		if start > 0 && limit > 0 {
			endpoint = endpoint + "/?start=" + strconv.Itoa(start) + "&limit=" + strconv.Itoa(limit)
		} else if start > 0 {
			endpoint = endpoint + "/?start=" + strconv.Itoa(start)
		} else {
			endpoint = endpoint + "/?limit=" + strconv.Itoa(limit)
		}
	}
	return endpoint
}
