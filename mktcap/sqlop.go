package mktcap

import (
	"database/sql"
	/*SQL Driver*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

//SQLConnect connect to sQL
func SQLConnect(name, passwd, address, dbname string) (*sql.DB, error) {
	url := name + ":" + passwd + "@tcp(" + address + ":3306)/" + dbname
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	glog.V(2).Infoln("Sql Service Connected")
	return db, nil
}

/*SQLDisconnect disconnect specific db*/
func SQLDisconnect(db *sql.DB) {
	if db != nil {
		db.Close()
	}
	glog.V(2).Infoln("Sql Service Disconnected")
}

func insertMultRows(db *sql.DB, table string, records []interface{}) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		glog.Errorln(err)
		return len(records), err
	}
	defer tx.Rollback()
	// 15 items
	insertStr := "INSERT INTO " + table + " (id, name, symbol, rank, price_usd," +
		" price_btc, market_cap_usd, available_supply, total_supply, max_supply, last_updated, percent_change_1h," +
		" percent_change_24h, percent_change_7d, volume_usd_24h)" +
		" VALUES (? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,?, ?, ?, ?)"

	stm, err := tx.Prepare(insertStr)
	if err != nil {
		glog.Errorln(err)
		return len(records), err
	}
	defer stm.Close()
	for _, record := range records {
		pair := record.(map[string]interface{})
		_, err = stm.Exec(pair["id"], pair["name"], pair["symbol"], pair["rank"], pair["price_usd"],
			pair["price_btc"], pair["market_cap_usd"], pair["available_supply"], pair["total_supply"], pair["max_supply"],
			pair["last_updated"], pair["percent_change_1h"], pair["percent_change_24h"], pair["percent_change_7d"], pair["24h_volume_usd"])
	}
	err = tx.Commit()
	if err != nil {
		glog.Errorln(err)
		return len(records), err
	}
	return len(records), nil
}
