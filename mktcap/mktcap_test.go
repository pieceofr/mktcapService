package mktcap

import (
	//"fmt"
	"testing"

	"github.com/golang/glog"
)

const sqlurl = ""
const sqluser = ""
const sqlpwd = ""
const sqldb = ""
const tickertable = ""

func TestConfig(t *testing.T) {
	conf := InitConfig(sqlurl, sqluser, sqlpwd, sqldb, tickertable)
	if conf.Sqlurl != sqlurl || conf.Sqluser != sqluser || conf.Sqlpwd != sqlpwd || conf.sqlTickerTable != tickertable {
		glog.Errorln(conf)
		t.Error()
	}
}

/* run test by go test -args -logtostderr=true -v=2 to enable glog*/

func TestTickerGet(t *testing.T) {
	conf := InitConfig(sqlurl, sqluser, sqlpwd, sqldb, tickertable)
	list, err := TickerNow(0, 10, conf)
	if err != nil {
		glog.Infoln(err)
		t.FailNow()
	}
	for _, val := range list {
		glog.V(2).Infoln(val)
	}
}

func TestTickerSave(t *testing.T) {
	conf := InitConfig(sqlurl, sqluser, sqlpwd, sqldb, tickertable)
	err := TickerSave(0, 5, conf)
	if err != nil {
		glog.Infoln(err)
		t.FailNow()
	}

}
