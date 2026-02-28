package myxcoinapi

type PublicRestPublicSymbolsAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicSymbolsReq
}

type PublicRestPublicSymbolsReq struct {
	Symbol       *string `json:"symbol"`
	BusinessType *string `json:"businessType"`
	BaseCurrency *string `json:"baseCurrency"`
}

// string false Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
func (api *PublicRestPublicSymbolsAPI) Symbol(symbol string) *PublicRestPublicSymbolsAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string true Instrument type: spot for spot or spot margin, linear_perpetual for USDT-margined perpetual, linear_futures for delivery futures
func (api *PublicRestPublicSymbolsAPI) BusinessType(businessType string) *PublicRestPublicSymbolsAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string false baseCurrency, only applicable to futures
func (api *PublicRestPublicSymbolsAPI) BaseCurrency(baseCurrency string) *PublicRestPublicSymbolsAPI {
	api.req.BaseCurrency = GetPointer(baseCurrency)
	return api
}

type PublicRestPublicBaseRatesAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicBaseRatesReq
}

type PublicRestPublicBaseRatesReq struct {
	Currency *string `json:"currency"`
}

// string false Curreny, e.g., BTC. If omitted, all spot margin currencies will be returned.
func (api *PublicRestPublicBaseRatesAPI) Currency(currency string) *PublicRestPublicBaseRatesAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

type PublicRestPublicSpotMarginCollateralAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicSpotMarginCollateralReq
}

type PublicRestPublicSpotMarginCollateralReq struct {
	Currency *string `json:"currency"`
}

// string false Curreny, e.g., BTC. If omitted, all spot margin currencies will be returned.
func (api *PublicRestPublicSpotMarginCollateralAPI) Currency(currency string) *PublicRestPublicSpotMarginCollateralAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

type PublicRestPublicHaircutAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicHaircutReq
}

type PublicRestPublicHaircutReq struct {
	Currency *string `json:"currency"`
}

// string false Curreny, e.g., BTC. If omitted, all spot margin currencies will be returned.
func (api *PublicRestPublicHaircutAPI) Currency(currency string) *PublicRestPublicHaircutAPI {
	api.req.Currency = GetPointer(currency)
	return api
}