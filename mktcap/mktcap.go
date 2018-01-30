package mktcap

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/golang/glog"
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
	resp, err := http.Get(getTickerEndPoint(start, limit, config))
	if err != nil {
		glog.Errorln(getTickerEndPoint(start, limit, config))
		return []MktCapInfo{}, err
	}
	defer func() {
		if resp != nil && resp.Body != nil { //prevent nil make crush
			resp.Body.Close()
		}
	}()
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

//TickerNowByIDList Only Get the ids list records
func TickerNowByIDList(ids []string, config MktcapConfig) ([]MktCapInfo, error) {
	var ret []MktCapInfo
	for _, val := range ids {
		resp, err := http.Get(getTickerCoinEndPoint(val, config))
		if err != nil {
			glog.Errorln(getTickerCoinEndPoint(val, config))
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			return []MktCapInfo{}, err
		}
		jsondata, err := getRawData(body)
		if err != nil {
			resp.Body.Close()
			return []MktCapInfo{}, err
		}
		list := DataToMktCapInfo(jsondata)
		if len(list) == 1 {
			ret = append(ret, list[0])
		}
		if resp.Body != nil {
			resp.Body.Close()
		}

	}
	return ret, nil
}

//TickerSave save ticker datat to sql database
func TickerSave(start, limit int, config MktcapConfig) error {
	resp, err := http.Get(getTickerEndPoint(start, limit, config))
	if err != nil {
		glog.Errorln(getTickerEndPoint(start, limit, config))
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
func getTickerEndPoint(start, limit int, config MktcapConfig) string {
	endpoint := "https://" + config.MktCapEndpoint + "/" + config.MktAPIVer + "/" + config.MktAPITicker

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
func getTickerCoinEndPoint(coin string, config MktcapConfig) string {
	endpoint := "https://" + config.MktCapEndpoint + "/" + config.MktAPIVer + "/" + config.MktAPITicker
	endpoint = endpoint + "/" + coin
	return endpoint
}
