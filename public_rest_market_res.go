package myxcoinapi

type PublicRestMarketTimeRes struct {
	Time string `json:"time"`
}

type Books struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type PublicRestMarketDepthResMiddle struct {
	Asks         [][]string `json:"asks"`
	Bids         [][]string `json:"bids"`
	LastUpdateId string     `json:"lastUpdateId"`
}

type PublicRestMarketDepthRes struct {
	Asks         []Books `json:"asks"`
	Bids         []Books `json:"bids"`
	LastUpdateId string  `json:"lastUpdateId"`
}

func (m *PublicRestMarketDepthResMiddle) ConvertToRes() *PublicRestMarketDepthRes {
	asks := make([]Books, len(m.Asks))
	for i, ask := range m.Asks {
		asks[i] = Books{
			Price:    ask[0],
			Quantity: ask[1],
		}
	}
	bids := make([]Books, len(m.Bids))
	for i, bid := range m.Bids {
		bids[i] = Books{
			Price:    bid[0],
			Quantity: bid[1],
		}
	}
	return &PublicRestMarketDepthRes{
		Asks:         asks,
		Bids:         bids,
		LastUpdateId: m.LastUpdateId,
	}
}

type PublicRestMarketTickerMiniResRow struct {
	BusinessType       string `json:"businessType"`
	Symbol             string `json:"symbol"`
	LastPrice          string `json:"lastPrice"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	FillQty            string `json:"fillQty"`
	FillAmount         string `json:"fillAmount"`
	BaseCurrency       string `json:"baseCurrency"`
}
type PublicRestMarketTickerMiniRes []PublicRestMarketTickerMiniResRow

type PublicRestMarketTradeResRow struct {
	Id         string `json:"id"`         // Unique trade ID
	Symbol     string `json:"symbol"`     // Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
	Price      string `json:"price"`      // Fill price
	Qty        string `json:"qty"`        // Filled quantity For spot margin, in units of the token; For USDT-margined perpetual futures, in units of the token
	Time       string `json:"time"`       // Trade timestamp in Unix milliseconds, e.g., 1732158178000
	Side       string `json:"side"`       // Order direction buy: Buy, sell: Sell
	IndexPrice string `json:"indexPrice"` // Index price when order filled
	MarkPrice  string `json:"markPrice"`  // Market price when order filled
}
type PublicRestMarketTradeRes []PublicRestMarketTradeResRow

type PublicRestMarketTicker24hrResRow struct {
	BusinessType       string `json:"businessType"`       // Instrument type: spot for spot or spot margin, linear_perpetual for USDT-margined perpetual, linear_futures for delivery futures
	Symbol             string `json:"symbol"`             // Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
	LastPrice          string `json:"lastPrice"`          // Last fill price
	PriceChange        string `json:"priceChange"`        // Price change amount in USDT
	PriceChangePercent string `json:"priceChangePercent"` // Price change percentage
	HighPrice          string `json:"highPrice"`          // Highest price
	LowPrice           string `json:"lowPrice"`           // Lowest price
	OpenPrice          string `json:"openPrice"`          // Opening price
	FillQty            string `json:"fillQty"`            // Filled quantity For spot margin, in units of the token; For USDT-margined perpetual futures, in units of the token
	FillAmount         string `json:"fillAmount"`         // Filled amount in USDT
	Count              string `json:"count"`              // Number of trades
	BaseCurrency       string `json:"baseCurrency"`       // Base currency, e.g., BTC in BTC-USDT, only for spot or spot margin
	IndexPrice         string `json:"indexPrice"`         // Index price
	MarkPrice          string `json:"markPrice"`          // Market price
	FundingRate        string `json:"fundingRate"`        // Current funding rate Only for futures; 0 in other cases
	ToNextFundRateTime string `json:"toNextFundRateTime"` // Timestamp of next funding rate, in Unix milliseconds, e.g., 1732158178000
}
type PublicRestMarketTicker24hrRes []PublicRestMarketTicker24hrResRow

type Kline struct {
	Period             string `json:"period"`
	OpenTime           string `json:"openTime"`
	CloseTime          string `json:"closeTime"`
	OpenPrice          string `json:"openPrice"`
	ClosePrice         string `json:"closePrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	FillQty            string `json:"fillQty,omitempty"`
	FillAmount         string `json:"fillAmount,omitempty"`
	Count              string `json:"count,omitempty"`
	PriceChange        string `json:"priceChange,omitempty"`
	PriceChangePercent string `json:"priceChangePercent,omitempty"`
}

type PublicRestMarketKlineResMiddle [][]string
type PublicRestMarketKlineRes []Kline

func (m *PublicRestMarketKlineResMiddle) ConvertToRes() *PublicRestMarketKlineRes {
	klines := make([]Kline, len(*m))
	for i, kline := range *m {
		klines[i] = Kline{
			Period:             kline[0],
			OpenTime:           kline[1],
			CloseTime:          kline[2],
			OpenPrice:          kline[3],
			ClosePrice:         kline[4],
			HighPrice:          kline[5],
			LowPrice:           kline[6],
			FillQty:            kline[7],
			FillAmount:         kline[8],
			Count:              kline[9],
			PriceChange:        kline[10],
			PriceChangePercent: kline[11],
		}
	}
	res := PublicRestMarketKlineRes(klines)
	return &res
}

type PublicRestMarketMarkPriceKlineResMiddle [][]string
type PublicRestMarketMarkPriceKlineRes []Kline

func (m *PublicRestMarketMarkPriceKlineResMiddle) ConvertToRes() *PublicRestMarketMarkPriceKlineRes {
	klines := make([]Kline, len(*m))
	for i, kline := range *m {
		klines[i] = Kline{
			Period:     kline[0],
			OpenTime:   kline[1],
			CloseTime:  kline[2],
			OpenPrice:  kline[3],
			ClosePrice: kline[4],
			HighPrice:  kline[5],
			LowPrice:   kline[6],
		}
	}
	res := PublicRestMarketMarkPriceKlineRes(klines)
	return &res
}

type PublicRestMarketIndexPriceKlineResMiddle [][]string
type PublicRestMarketIndexPriceKlineRes []Kline

func (m *PublicRestMarketIndexPriceKlineResMiddle) ConvertToRes() *PublicRestMarketIndexPriceKlineRes {
	klines := make([]Kline, len(*m))
	for i, kline := range *m {
		klines[i] = Kline{
			Period:     kline[0],
			OpenTime:   kline[1],
			CloseTime:  kline[2],
			OpenPrice:  kline[3],
			ClosePrice: kline[4],
			HighPrice:  kline[5],
			LowPrice:   kline[6],
		}
	}
	res := PublicRestMarketIndexPriceKlineRes(klines)
	return &res
}

type PublicRestMarketDeliveryExcerciseHistoryResRow struct {
	Symbol string `json:"symbol"` // Trading pair name, e.g., BTC-USDT-3JAN25
	Type   string `json:"type"`   // Type: delivery: Delivery
	Price  string `json:"price"`  // Delivery price
	Time   string `json:"time"`   // Delivery timestamp in Unix milliseconds, e.g., 1732158178000
}
type PublicRestMarketDeliveryExcerciseHistoryRes []PublicRestMarketDeliveryExcerciseHistoryResRow

type PublicRestMarketFundingRateResRow struct {
	Symbol           string `json:"symbol"`           // Trading pair name, e.g., BTC-USDT-PERP
	FundingRate      string `json:"fundingRate"`      // Expected funding rate
	FundingTime      string `json:"fundingTime"`      // Timestamp when the upcoming funding fee will be collected, in Unix milliseconds, e.g., 1732158178000
	FundingInterval  string `json:"fundingInterval"`  // Current effective funding rate interval, measured in hours
	UpperFundingRate string `json:"upperFundingRate"` // Upper limit of expected funding rate
	LowerFundingRate string `json:"lowerFundingRate"` // Lower limit of expected funding rate
}
type PublicRestMarketFundingRateRes []PublicRestMarketFundingRateResRow

type PublicRestMarketFundingRateHistoryResRow struct {
	Symbol      string `json:"symbol"`      // Trading pair name, e.g., BTC-USDT-PERP
	FundingRate string `json:"fundingRate"` // Funding rate
	FundingTime string `json:"fundingTime"` // Timestamp when the current funding fee was collected, in Unix milliseconds, e.g., 1732158178000
	MarkPrice   string `json:"markPrice"`   // Mark price
}
type PublicRestMarketFundingRateHistoryRes []PublicRestMarketFundingRateHistoryResRow
