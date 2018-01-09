package mktcap

import (
	"strconv"
)

type MktCapInfo struct {
	ID               string
	Name             string
	Symbol           string
	Rank             int
	PriceUsd         float64
	PriceBtc         float64
	VolumeUsd24h     float64
	MarketCapUsd     float64
	AvailibleSupply  float64
	TotalSupply      float64
	MaxSupply        float64
	PercentChange1h  float64
	PercentChange24h float64
	PercentChange7d  float64
	LastUpdated      int64
}

func DataToMktCapInfo(f []interface{}) []MktCapInfo {
	infoArr := make([]MktCapInfo, 0, 100)
	for _, v := range f {
		info := MktCapInfo{}
		for key, val := range v.(map[string]interface{}) {
			switch key {
			case "id":
				if val != nil {
					info.ID = val.(string)
				}
			case "name":
				if val != nil {
					info.Name = val.(string)
				}
			case "symbol":
				if val != nil {
					info.Symbol = val.(string)
				}
			case "rank":
				if val != nil {
					iv, err := strconv.Atoi(val.(string))
					if err != nil {
						break
					}
					info.Rank = iv
				}
			case "price_usd":
				if val != nil {
					fv, err := strconv.ParseFloat(val.(string), 64)
					if err != nil {
						break
					}
					info.PriceUsd = fv
				}
			case "price_btc":
				if val != nil {
					fv, err := strconv.ParseFloat(val.(string), 64)
					if err != nil {
						break
					}
					info.PriceBtc = fv
				}
			case "market_cap_usd":
				if val != nil {
					fv, err := strconv.ParseFloat(val.(string), 64)
					if err != nil {
						break
					}
					info.MarketCapUsd = fv
				}
			case "available_supply":
				if val != nil {
					fv, err := strconv.ParseFloat(val.(string), 64)
					if err != nil {
						break
					}
					info.AvailibleSupply = fv
				}
			case "total_supply":
				if val != nil {
					fv, err := strconv.ParseFloat(val.(string), 64)
					if err != nil {
						break
					}
					info.TotalSupply = fv
				}
			case "max_supply":
				if val != nil {
					fv, err := strconv.ParseFloat(val.(string), 64)
					if err != nil {
						break
					}
					info.MaxSupply = fv
				}
			case "last_updated":
				if val != nil {
					iv, err := strconv.ParseInt(val.(string), 10, 64)
					if err != nil {
						break
					}
					info.LastUpdated = iv
				}
			case "percent_change_1h":
				if val != nil {
					fv, err := strconv.ParseFloat(val.(string), 64)
					if err != nil {
						break
					}
					info.PercentChange1h = fv
				}
			case "percent_change_24h":
				if val != nil {
					fv, err := strconv.ParseFloat(val.(string), 64)
					if err != nil {
						break
					}
					info.PercentChange24h = fv
				}
			case "percent_change_7d":
				if val != nil {
					fv, err := strconv.ParseFloat(val.(string), 64)
					if err != nil {
						break
					}
					info.PercentChange7d = fv
				}
			case "24h_volume_usd":
				if val != nil {
					fv, err := strconv.ParseFloat(val.(string), 64)
					if err != nil {
						break
					}
					info.VolumeUsd24h = fv
				}
			}
		}
		infoArr = append(infoArr, info)
	}
	return infoArr
}
