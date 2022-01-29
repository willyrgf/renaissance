package coingecko

// import (
// 	"encoding/json"
// 	"time"

// 	"github.com/anonq/fund-history/pkg/httpclient"
// )

// type MarketTicker struct {
// 	Base   string `json:"base"`
// 	Target string `json:"target"`
// 	Market struct {
// 		Name                string `json:"name"`
// 		Identifier          string `json:"identifier"`
// 		HasTradingIncentive bool   `json:"has_trading_incentive"`
// 	} `json:"market"`
// 	Last          float64 `json:"last"`
// 	Volume        float64 `json:"volume"`
// 	ConvertedLast struct {
// 		Btc float64 `json:"btc"`
// 		Eth float64 `json:"eth"`
// 		Usd float64 `json:"usd"`
// 	} `json:"converted_last"`
// 	ConvertedVolume struct {
// 		Btc float64 `json:"btc"`
// 		Eth float64 `json:"eth"`
// 		Usd float64 `json:"usd"`
// 	} `json:"converted_volume"`
// 	TrustScore             string      `json:"trust_score"`
// 	BidAskSpreadPercentage float64     `json:"bid_ask_spread_percentage"`
// 	Timestamp              time.Time   `json:"timestamp"`
// 	LastTradedAt           time.Time   `json:"last_traded_at"`
// 	LastFetchAt            time.Time   `json:"last_fetch_at"`
// 	IsAnomaly              bool        `json:"is_anomaly"`
// 	IsStale                bool        `json:"is_stale"`
// 	TradeURL               string      `json:"trade_url"`
// 	TokenInfoURL           interface{} `json:"token_info_url"`
// 	CoinID                 string      `json:"coin_id"`
// 	TargetCoinID           string      `json:"target_coin_id"`
// }

// type CoinTickers struct {
// 	Name    string         `json:"name"`
// 	Tickers []MarketTicker `json:"tickers"`
// }

// func FetchCoinTickers(c httpclient.Do, id string) (*CoinTickers, error) {
// 	reqInfo := &reqInfo{
// 		endpoint: "/coins/" + id + "/tickers",
// 	}

// 	u, p, err := reqInfo.parse()
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp, err := c.GET(u, p)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var data *CoinTickers
// 	err = json.Unmarshal(resp, &data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }
