package myxcoinapi

// GET Get the Basic Information of Trading Products.
func (client *PublicRestClient) NewPublicRestPublicSymbols() *PublicRestPublicSymbolsAPI {
	return &PublicRestPublicSymbolsAPI{
		client: client,
		req:    &PublicRestPublicSymbolsReq{},
	}
}

func (api *PublicRestPublicSymbolsAPI) Do() (*XcoinRestRes[PublicRestPublicSymbolsRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestPublicSymbols])
	return xcoinCallAPI[PublicRestPublicSymbolsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 	Get Margin Interest Rates.
func (client *PublicRestClient) NewPublicRestPublicBaseRates() *PublicRestPublicBaseRatesAPI {
	return &PublicRestPublicBaseRatesAPI{
		client: client,
		req:    &PublicRestPublicBaseRatesReq{},
	}
}

func (api *PublicRestPublicBaseRatesAPI) Do() (*XcoinRestRes[PublicRestPublicBaseRatesRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestPublicBaseRates])
	return xcoinCallAPI[PublicRestPublicBaseRatesRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 	Get Spot Leverage Parameters
func (client *PublicRestClient) NewPublicRestPublicSpotMarginCollateral() *PublicRestPublicSpotMarginCollateralAPI {
	return &PublicRestPublicSpotMarginCollateralAPI{
		client: client,
		req:    &PublicRestPublicSpotMarginCollateralReq{},
	}
}

func (api *PublicRestPublicSpotMarginCollateralAPI) Do() (*XcoinRestRes[PublicRestPublicSpotMarginCollateralRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestPublicSpotMarginCollateral])
	return xcoinCallAPI[PublicRestPublicSpotMarginCollateralRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 	Get Currency Asset Discount Ratio Information
func (client *PublicRestClient) NewPublicRestPublicHaircut() *PublicRestPublicHaircutAPI {
	return &PublicRestPublicHaircutAPI{
		client: client,
		req:    &PublicRestPublicHaircutReq{},
	}
}

func (api *PublicRestPublicHaircutAPI) Do() (*XcoinRestRes[PublicRestPublicHaircutRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestPublicHaircut])
	return xcoinCallAPI[PublicRestPublicHaircutRes](api.client.c, url, NIL_REQBODY, GET)
}