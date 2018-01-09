package mktcap

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

//vm.dyn.pieceofr.click cryptomarket
func SqlConnect(name, passwd, address, dbname string) (*sql.DB, error) {
	url := name + ":" + passwd + "@tcp(" + address + ":3306)/" + dbname
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	glog.V(2).Infoln("Sql Service Connected")
	return db, nil
}

func SqlDisconnect(db *sql.DB) {
	if db != nil {
		db.Close()
	}
	glog.V(2).Infoln("Sql Service Disconnected")
}

func insertMultiple(db *sql.DB, table string, records []interface{}) (int, error) {
	if len(records) <= 0 {
		return 0, nil
	}
	keySeq := make([]string, 0, 14)
	//Make Key Array
	for k, _ := range records[0].(map[string]interface{}) {
		keySeq = append(keySeq, k)
	}
	istVal := ""
	for _, record := range records {
		pair := record.(map[string]interface{})
		valpair := ""
		for i := 0; i < len(keySeq); i++ {
			v := pair[keySeq[i]]
			switch keySeq[i] {
			case "id":
				fallthrough
			case "name":
				fallthrough
			case "symbol":
				if v != nil {
					valpair = fmt.Sprintf("%s, '%s'", valpair, v.(string))
				} else {
					valpair = fmt.Sprintf("%s, '%s'", valpair, "")
				}

			case "rank":
				fallthrough
			case "price_usd":
				fallthrough
			case "price_btc":
				fallthrough
			case "market_cap_usd":
				fallthrough
			case "available_supply":
				fallthrough
			case "total_supply":
				fallthrough
			case "max_supply":
				fallthrough
			case "last_updated":
				fallthrough
			case "percent_change_1h":
				fallthrough
			case "percent_change_24h":
				fallthrough
			case "percent_change_7d":
				fallthrough
			case "24h_volume_usd":
				if v != nil {
					valpair = fmt.Sprintf("%s, %s", valpair, v)
				} else { //make null = 0
					valpair = fmt.Sprintf("%s, %s", valpair, "0")
				}
			}

		}
		valpair = fmt.Sprintf("(%s%s", strings.TrimLeft(valpair, ","), ")")
		//Each Loop Needs to add a comma
		istVal = fmt.Sprintf("%s,%s", istVal, valpair)
	}
	//Trim end comma
	istVal = fmt.Sprintf("%s", strings.TrimLeft(istVal, ","))
	istKey := ""
	//Make Key Array
	for _, val := range keySeq {
		if val == "24h_volume_usd" {
			istKey = fmt.Sprintf("%s, %s", istKey, "volume_usd_24h")
		} else {
			istKey = fmt.Sprintf("%s, %s", istKey, val)
		}
	}
	istKey = fmt.Sprintf("(%s%s", strings.TrimLeft(istKey, ","), ")")

	sqlStatement := fmt.Sprintf("INSERT INTO %s %s VALUES %s", table, istKey, istVal)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return len(records), err
	}
	return len(records), nil
}

func insertIntoDB(db *sql.DB, table string, pair map[string]interface{}) error {
	istKey := ""
	istVal := ""
	for k, v := range pair {
		switch k {
		case "id":
			fallthrough
		case "name":
			fallthrough
		case "symbol":
			istKey = fmt.Sprintf("%s, %s", istKey, k)
			istVal = fmt.Sprintf("%s, '%s'", istVal, v.(string))
		case "rank":
			fallthrough
		case "price_usd":
			fallthrough
		case "price_btc":
			fallthrough
		case "market_cap_usd":
			fallthrough
		case "available_supply":
			fallthrough
		case "total_supply":
			fallthrough
		case "max_supply":
			fallthrough
		case "last_updated":
			fallthrough
		case "percent_change_1h":
			fallthrough
		case "percent_change_24h":
			fallthrough
		case "percent_change_7d":
			istKey = fmt.Sprintf("%s, %s", istKey, k)
			istVal = fmt.Sprintf("%s, %s", istVal, v.(string))
		case "24h_volume_usd":
			istKey = fmt.Sprintf("%s, %s", istKey, "volume_usd_24h")
			istVal = fmt.Sprintf("%s, %s", istVal, v.(string))
		}
	}
	istKey = fmt.Sprintf("(%s%s", strings.TrimLeft(istKey, ","), ")")
	istVal = fmt.Sprintf("(%s%s", strings.TrimLeft(istVal, ","), ")")
	sqlStatement := fmt.Sprintf("INSERT INTO %s %s VALUES %s", table, istKey, istVal)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return nil
}
