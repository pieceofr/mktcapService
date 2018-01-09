package main

import (
	"github.com/glog"
	"github.com/viper"
)

//ServiceConfig config struct
type ServiceConfig struct {
	SQLEndpoint           string
	SQLUser               string
	SQLPwd                string
	SQLDB                 string
	SQLTickerTable        string
	QuickMonitor          bool
	SaveToDB              bool
	QuickMonitorInterval  int
	QuickMonitorLimit     int
	SaveToDBInterval      int
	SaveToDBLimit         int
	SlackToken            string
	SlackChannel          string
	RuleSigDiffObserv     int
	RuleSigDiffThreadhold float64
}

//InitConfig read from file and make a ServiceConfig
func InitConfig() ServiceConfig {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.SetDefault("enableService.quickMonitor", true)
	viper.SetDefault("enableService.saveToDB", false)

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		glog.Fatalln("Fatal error config file: %s \n", err)
	}
	config := ServiceConfig{SQLEndpoint: viper.GetString("database.sqlendpoint"),
		SQLUser:               viper.GetString("database.sqluser"),
		SQLPwd:                viper.GetString("databse.sqlpwd"),
		SQLDB:                 viper.GetString("database.sqldb"),
		SQLTickerTable:        viper.GetString("database.tickertable"),
		QuickMonitor:          viper.GetBool("enableService.quickMonitor"),
		SaveToDB:              viper.GetBool("enableService.saveToDB"),
		QuickMonitorInterval:  viper.GetInt("quickMonitorService.monitorIntervalSec"),
		QuickMonitorLimit:     viper.GetInt("quickMonitorService.monitorLimitRecords"),
		SaveToDBInterval:      viper.GetInt("saveToDBService.saveToDBSec"),
		SaveToDBLimit:         viper.GetInt("saveToDBService.saveToDBLimitRecords"),
		SlackToken:            viper.GetString("slack.token"),
		SlackChannel:          viper.GetString("slack.slackChannel"),
		RuleSigDiffObserv:     viper.GetInt("ruleSigDiff.numObervations"),
		RuleSigDiffThreadhold: viper.GetFloat64("ruleSigDiff.threadholePercnt")}

	return config

}
