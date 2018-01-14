package main

import (
	"reflect"

	"github.com/golang/glog"
	"github.com/spf13/viper"
)

//ServiceConfig config struct
type ServiceConfig struct {
	SQLEndpoint          string
	SQLUser              string
	SQLPwd               string
	SQLDB                string
	SQLTickerTable       string
	QuickMonitor         bool
	SaveToDB             bool
	APIService           bool
	QuickMonitorInterval int
	QuickMonitorLimit    int
	SaveToDBInterval     int
	SaveToDBLimit        int
	SlackToken           string
	SlackChannel         string
	//	RuleSigDiffObserv     int
	//	RuleSigDiffThreadhold float64
	MonitorType       string
	MonitorCoinList   []string
	ApiPort           int
	SigDiffConditions []SigDiffCond
}

//InitConfig read from file and make a ServiceConfig
func InitConfig() ServiceConfig {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.SetDefault("enableService.quickMonitor", true)
	viper.SetDefault("enableService.saveToDB", false)
	viper.SetDefault("apiServer.port", 8080)

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		glog.Fatalln("Fatal error config file: %s \n", err)
	}
	config := ServiceConfig{SQLEndpoint: viper.GetString("database.sqlendpoint"),
		SQLUser:              viper.GetString("database.sqluser"),
		SQLPwd:               viper.GetString("database.sqlpwd"),
		SQLDB:                viper.GetString("database.mktcapdb"),
		SQLTickerTable:       viper.GetString("database.tickertable"),
		QuickMonitor:         viper.GetBool("enableService.quickMonitor"),
		SaveToDB:             viper.GetBool("enableService.saveToDB"),
		APIService:           viper.GetBool("enableService.apiService"),
		QuickMonitorInterval: viper.GetInt("quickMonitorService.monitorIntervalSec"),
		QuickMonitorLimit:    viper.GetInt("quickMonitorService.monitorLimitRecords"),
		SaveToDBInterval:     viper.GetInt("saveToDBService.saveToDBSec"),
		SaveToDBLimit:        viper.GetInt("saveToDBService.saveToDBLimitRecords"),
		SlackToken:           viper.GetString("slack.token"),
		SlackChannel:         viper.GetString("slack.slackChannel"),
		MonitorType:          viper.GetString("quickMonitorService.monitorType"),
		MonitorCoinList:      viper.GetStringSlice("quickMonitorService.monitorCoinIDs"),
		ApiPort:              viper.GetInt("apiServer.port")}
	sigCondRaw := viper.Get("ruleSigDiff")
	conds := make([]SigDiffCond, 0, 5)
	for _, val := range sigCondRaw.([]interface{}) {
		//map[interface {}]interface {}
		thrhold := val.(map[interface{}]interface{})["threadholdPercnt"]
		obvs := val.(map[interface{}]interface{})["numObervations"]
		var cond SigDiffCond
		if reflect.TypeOf(thrhold).String() == "float64" {
			cond = SigDiffCond{Observations: obvs.(int), Threadhold: thrhold.(float64)}
		} else if reflect.TypeOf(thrhold).String() == "int" {
			cond = SigDiffCond{Observations: obvs.(int), Threadhold: float64(thrhold.(int))}
		}
		conds = append(conds, cond)
	}
	config.SigDiffConditions = conds
	return config
}
