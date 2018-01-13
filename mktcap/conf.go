package mktcap

import (
	"github.com/golang/glog"

	"github.com/spf13/viper"
)

//APIEndPoint coinmarketcap.com api endpoint
const APIEndPoint = "api.coinmarketcap.com"

//APIVer api version number
const APIVer = "v1"

//APITicker tick apis
const APITicker = "ticker"

//MktcapConfig Config file for mktcap package
type MktcapConfig struct {
	MktCapEndpoint string
	MktAPIVer      string
	MktAPITicker   string
	Sqlurl         string
	Sqluser        string
	Sqlpwd         string
	Sqldb          string
	sqlTickerTable string
}

/*InitConfig initialized mktcap config*/
func InitConfig(sqlurl, sqluser, sqlpwd, sqldb, sqlticktable string) MktcapConfig {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		glog.Fatalln("Fatal error config file: %s \n", err)
	}
	viper.SetDefault("endpoint", APIEndPoint)
	viper.SetDefault("ver", APIVer)
	viper.SetDefault("tickerapi", APITicker)
	config := MktcapConfig{MktCapEndpoint: viper.GetString("endpoint"), MktAPIVer: viper.GetString("ver"),
		MktAPITicker: viper.GetString("tickerapi"), Sqlurl: sqlurl, Sqluser: sqluser, Sqlpwd: sqlpwd, Sqldb: sqldb, sqlTickerTable: sqlticktable}
	return config
}
