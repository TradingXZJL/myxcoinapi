package myxcoinapi

type PublicRestMarketTimeAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketTimeReq
}

type PublicRestMarketTimeReq struct{}

type PublicRestMarketDepthAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketDepthReq
}

type PublicRestMarketDepthReq struct {
	BusinessType *string `json:"businessType"`
	Symbol       *string `json:"symbol"`
	Limit        *string `json:"limit"`
}

// string false Instrument type: spot for spot or spot margin, linear_perpetual for USDT-margined perpetual, and linear_futures for delivery futures
func (api *PublicRestMarketDepthAPI) BusinessType(businessType string) *PublicRestMarketDepthAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string true Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
func (api *PublicRestMarketDepthAPI) Symbol(symbol string) *PublicRestMarketDepthAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false Number of entries to return (default: 100; max: 100)
func (api *PublicRestMarketDepthAPI) Limit(limit string) *PublicRestMarketDepthAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestMarketTickerMiniAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketTickerMiniReq
}

type PublicRestMarketTickerMiniReq struct {
	BusinessType *string `json:"businessType"`
	Symbol       *string `json:"symbol"`
	BaseCurrency *string `json:"baseCurrency"`
}

// string true Instrument type: spot for spot or spot margin, linear_perpetual for USDT-margined perpetual, and linear_futures for delivery futures
func (api *PublicRestMarketTickerMiniAPI) BusinessType(businessType string) *PublicRestMarketTickerMiniAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string false Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
func (api *PublicRestMarketTickerMiniAPI) Symbol(symbol string) *PublicRestMarketTickerMiniAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false Base currency, only valid for futures
func (api *PublicRestMarketTickerMiniAPI) BaseCurrency(baseCurrency string) *PublicRestMarketTickerMiniAPI {
	api.req.BaseCurrency = GetPointer(baseCurrency)
	return api
}

type PublicRestMarketTradeAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketTradeReq
}

type PublicRestMarketTradeReq struct {
	BusinessType *string `json:"businessType"`
	Symbol       *string `json:"symbol"`
	BaseCurrency *string `json:"baseCurrency"`
	Limit        *string `json:"limit"`
}

// string false Instrument type: spot for spot or spot margin, linear_perpetual for USDT-margined perpetual, linear_futures for delivery futures
func (api *PublicRestMarketTradeAPI) BusinessType(businessType string) *PublicRestMarketTradeAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string false Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
func (api *PublicRestMarketTradeAPI) Symbol(symbol string) *PublicRestMarketTradeAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false Base currency, only valid for futures
// Just one valid symbol or baseCurrency. When symbol filled in, the baseCurrency parameter is invalidated. The default value of baseCurrency is BTC. Other currencies need to be manually filled in; for example, if the input parameter is ETH, only ETH and related underlying assets will be returned
func (api *PublicRestMarketTradeAPI) BaseCurrency(baseCurrency string) *PublicRestMarketTradeAPI {
	api.req.BaseCurrency = GetPointer(baseCurrency)
	return api
}

// string false Number of entries to return (default: 100; max: 100)
func (api *PublicRestMarketTradeAPI) Limit(limit string) *PublicRestMarketTradeAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestMarketTicker24hrAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketTicker24hrReq
}

type PublicRestMarketTicker24hrReq struct {
	BusinessType *string `json:"businessType"`
	Symbol       *string `json:"symbol"`
	BaseCurrency *string `json:"baseCurrency"`
}

// string true Instrument type: spot for spot or spot margin, linear_perpetual for USDT-margined perpetual, linear_futures for delivery futures
func (api *PublicRestMarketTicker24hrAPI) BusinessType(businessType string) *PublicRestMarketTicker24hrAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string false Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
func (api *PublicRestMarketTicker24hrAPI) Symbol(symbol string) *PublicRestMarketTicker24hrAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false Base currency, only valid for futures
// Just one valid symbol or baseCurrency. When symbol filled in, the baseCurrency parameter is invalidated. The default value of baseCurrency is BTC. Other currencies need to be manually filled in; for example, if the input parameter is ETH, only ETH and related underlying assets will be returned
func (api *PublicRestMarketTicker24hrAPI) BaseCurrency(baseCurrency string) *PublicRestMarketTicker24hrAPI {
	api.req.BaseCurrency = GetPointer(baseCurrency)
	return api
}

type PublicRestMarketKlineAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketKlineReq
}

type PublicRestMarketKlineReq struct {
	BusinessType *string `json:"businessType"` // Instrument type: spot for spot or spot margin, linear_perpetual for USDT-margined perpetual, linear_futures for delivery futures
	Period       *string `json:"period"`       // Interval period, e.g., 1s, 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1M
	Symbol       *string `json:"symbol"`       // Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
	StartTime    *string `json:"startTime"`    // Start timestamp in Unix milliseconds, e.g., 1732158178000
	EndTime      *string `json:"endTime"`      // End timestamp in Unix milliseconds, e.g., 1732158178000
	Limit        *string `json:"limit"`        // Number of entries to return (default: 1000; max: 1000)
}

// string false Instrument type: spot for spot or spot margin, linear_perpetual for USDT-margined perpetual, linear_futures for delivery futures
func (api *PublicRestMarketKlineAPI) BusinessType(businessType string) *PublicRestMarketKlineAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string true Interval period, e.g., 1s, 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1M
func (api *PublicRestMarketKlineAPI) Period(period string) *PublicRestMarketKlineAPI {
	api.req.Period = GetPointer(period)
	return api
}

// string false Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
func (api *PublicRestMarketKlineAPI) Symbol(symbol string) *PublicRestMarketKlineAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false Start timestamp in Unix milliseconds, e.g., 1732158178000
func (api *PublicRestMarketKlineAPI) StartTime(startTime string) *PublicRestMarketKlineAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

// string false End timestamp in Unix milliseconds, e.g., 1732158178000
func (api *PublicRestMarketKlineAPI) EndTime(endTime string) *PublicRestMarketKlineAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// string false Number of entries to return (default: 1000; max: 1000)
func (api *PublicRestMarketKlineAPI) Limit(limit string) *PublicRestMarketKlineAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestMarketMarkPriceKlineAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketMarkPriceKlineReq
}

type PublicRestMarketMarkPriceKlineReq struct {
	BusinessType *string `json:"businessType"`
	Symbol       *string `json:"symbol"`
	Period       *string `json:"period"`
	StartTime    *string `json:"startTime"`
	EndTime      *string `json:"endTime"`
	Limit        *string `json:"limit"`
}

// string false Instrument type: spot for spot or spot margin, linear_perpetual for USDT-margined perpetual, linear_futures for delivery futures
func (api *PublicRestMarketMarkPriceKlineAPI) BusinessType(businessType string) *PublicRestMarketMarkPriceKlineAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string false Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
func (api *PublicRestMarketMarkPriceKlineAPI) Symbol(symbol string) *PublicRestMarketMarkPriceKlineAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string true Interval period, e.g., 1s, 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1M
func (api *PublicRestMarketMarkPriceKlineAPI) Period(period string) *PublicRestMarketMarkPriceKlineAPI {
	api.req.Period = GetPointer(period)
	return api
}

// string false Start timestamp in Unix milliseconds, e.g., 1732158178000
func (api *PublicRestMarketMarkPriceKlineAPI) StartTime(startTime string) *PublicRestMarketMarkPriceKlineAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

// string false End timestamp in Unix milliseconds, e.g., 1732158178000
func (api *PublicRestMarketMarkPriceKlineAPI) EndTime(endTime string) *PublicRestMarketMarkPriceKlineAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// string false Number of entries to return (default: 1000; max: 1000)
func (api *PublicRestMarketMarkPriceKlineAPI) Limit(limit string) *PublicRestMarketMarkPriceKlineAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestMarketIndexPriceKlineAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketIndexPriceKlineReq
}

type PublicRestMarketIndexPriceKlineReq struct {
	SymbolFamily *string `json:"symbolFamily"`
	Period       *string `json:"period"`
	StartTime    *string `json:"startTime"`
	EndTime      *string `json:"endTime"`
	Limit        *string `json:"limit"`
}

// string true Symbol family name, spot underlying of the contract, such as BTC-USDT
func (api *PublicRestMarketIndexPriceKlineAPI) SymbolFamily(symbolFamily string) *PublicRestMarketIndexPriceKlineAPI {
	api.req.SymbolFamily = GetPointer(symbolFamily)
	return api
}

// string true Interval period, e.g., 1s, 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1M
func (api *PublicRestMarketIndexPriceKlineAPI) Period(period string) *PublicRestMarketIndexPriceKlineAPI {
	api.req.Period = GetPointer(period)
	return api
}

// string false Start timestamp in Unix milliseconds, e.g., 1732158178000
func (api *PublicRestMarketIndexPriceKlineAPI) StartTime(startTime string) *PublicRestMarketIndexPriceKlineAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

// string false End timestamp in Unix milliseconds, e.g., 1732158178000
func (api *PublicRestMarketIndexPriceKlineAPI) EndTime(endTime string) *PublicRestMarketIndexPriceKlineAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// string false Number of entries to return (default: 1000; max: 1000)
func (api *PublicRestMarketIndexPriceKlineAPI) Limit(limit string) *PublicRestMarketIndexPriceKlineAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestMarketDeliveryExcerciseHistoryAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketDeliveryExcerciseHistoryReq
}

type PublicRestMarketDeliveryExcerciseHistoryReq struct {
	BusinessType *string `json:"businessType"`
	SymbolFamily *string `json:"symbolFamily"`
	Symbol       *string `json:"symbol"`
	BeginTime    *string `json:"beginTime"`
	EndTime      *string `json:"endTime"`
	Limit        *string `json:"limit"`
}

// string true Instrument type (only for delivery futures) linear_futures: Delivery futures
func (api *PublicRestMarketDeliveryExcerciseHistoryAPI) BusinessType(businessType string) *PublicRestMarketDeliveryExcerciseHistoryAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string true Symbol family name, spot underlying of the contract, such as BTC-USDT
func (api *PublicRestMarketDeliveryExcerciseHistoryAPI) SymbolFamily(symbolFamily string) *PublicRestMarketDeliveryExcerciseHistoryAPI {
	api.req.SymbolFamily = GetPointer(symbolFamily)
	return api
}

// string false Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
func (api *PublicRestMarketDeliveryExcerciseHistoryAPI) Symbol(symbol string) *PublicRestMarketDeliveryExcerciseHistoryAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false Start timestamp in Unix milliseconds, e.g., 1732158178000
func (api *PublicRestMarketDeliveryExcerciseHistoryAPI) BeginTime(beginTime string) *PublicRestMarketDeliveryExcerciseHistoryAPI {
	api.req.BeginTime = GetPointer(beginTime)
	return api
}

// string false End timestamp in Unix milliseconds, e.g., 1732158178000
func (api *PublicRestMarketDeliveryExcerciseHistoryAPI) EndTime(endTime string) *PublicRestMarketDeliveryExcerciseHistoryAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// string false Number of entries to return (default: 100; max: 100)
func (api *PublicRestMarketDeliveryExcerciseHistoryAPI) Limit(limit string) *PublicRestMarketDeliveryExcerciseHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestMarketFundingRateAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketFundingRateReq
}

type PublicRestMarketFundingRateReq struct {
	Symbol *string `json:"symbol"`
}

// string false Trading pair name, only for perpetual futures, e.g., BTC-USDT-PERP
func (api *PublicRestMarketFundingRateAPI) Symbol(symbol string) *PublicRestMarketFundingRateAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

type PublicRestMarketFundingRateHistoryAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketFundingRateHistoryReq
}

type PublicRestMarketFundingRateHistoryReq struct {
	Symbol    *string `json:"symbol"`
	BeginTime *string `json:"beginTime"`
	EndTime   *string `json:"endTime"`
	Limit     *string `json:"limit"`
}

// string true Trading pair name, only for perpetual futures, e.g., BTC-USDT-PERP
func (api *PublicRestMarketFundingRateHistoryAPI) Symbol(symbol string) *PublicRestMarketFundingRateHistoryAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false Start timestamp in Unix milliseconds, e.g., 1732158178000
func (api *PublicRestMarketFundingRateHistoryAPI) BeginTime(beginTime string) *PublicRestMarketFundingRateHistoryAPI {
	api.req.BeginTime = GetPointer(beginTime)
	return api
}

// string false End timestamp in Unix milliseconds, e.g., 1732158178000
func (api *PublicRestMarketFundingRateHistoryAPI) EndTime(endTime string) *PublicRestMarketFundingRateHistoryAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// string false Number of entries to return (default: 100; max: 100)
func (api *PublicRestMarketFundingRateHistoryAPI) Limit(limit string) *PublicRestMarketFundingRateHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
