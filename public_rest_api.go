package myxcoinapi

type PublicRestAPI int

const (
	// Market
	PublicRestMarketTime PublicRestAPI = iota
	PublicRestMarketDepth
	PublicRestMarketTickerMini
	PublicRestMarketTrade
	PublicRestMarketTicker24hr
	PublicRestMarketKline
	PublicRestMarketMarkPriceKline
	PublicRestMarketIndexPriceKline
	PublicRestMarketDeliveryExcerciseHistory
	PublicRestMarketFundingRate
	PublicRestMarketFundingRateHistory

	// Public
	PublicRestPublicSymbols
	PublicRestPublicBaseRates
	PublicRestPublicSpotMarginCollateral
	PublicRestPublicHaircut
)

var PublicRestAPIMap = map[PublicRestAPI]string{
	// Market
	PublicRestMarketTime:                     "/v1/market/time",                     // GET 	Get Server Time. 								This endpoint retrieves the server time, precise to the millisecond.
	PublicRestMarketDepth:                    "/v1/market/depth",                    // GET 	Get Order Book (Depth). 				This endpoint retrieves the order book depth for a specified trading pair. By default, 100 bids and 100 asks are returned.
	PublicRestMarketTickerMini:               "/v1/market/ticker/mini",              // GET 	Get Latest Ticker Information. 	This endpoint retrieves real-time ticker data for a specified trading pair.
	PublicRestMarketTrade:                    "/v1/market/trade",                    // GET 	Get Recent Trades. 							Retrieve recent trade information for a specified trading pair.
	PublicRestMarketTicker24hr:               "/v1/market/ticker/24hr",              // GET 	Get 24-Hour Ticker Data. 				Retrieve rolling 24-hour price data for a specified trading pair.
	PublicRestMarketKline:                    "/v1/market/kline",                    // GET 	Get Kline Data. 								This endpoint retrieves kline data for a specified trading pair.
	PublicRestMarketMarkPriceKline:           "/v1/market/markPriceKline",           // GET 	Get market price kline. 				Obtain the K-line data of the specified symbol in market price
	PublicRestMarketIndexPriceKline:          "/v1/market/indexPriceKline",          // GET 	Get index price kline. 					Obtain the K-line data of the specified symbol in index price
	PublicRestMarketDeliveryExcerciseHistory: "/v1/market/deliveryExercise/history", // GET 	Get Delivery History. 					This endpoint retrieves delivery records of delivery futures within any period up to 3 months.
	PublicRestMarketFundingRate:              "/v1/market/fundingRate",              // GET		Get Current Funding Rate. 			Retrieve the current funding rates (only for perpetual futures).
	PublicRestMarketFundingRateHistory:       "/v1/market/fundingRate/history",      // GET		Get Funding Rate History. 			This endpoint retrieves historical funding rates for perpetual futures within any period up to 6 months.

	// Public
	PublicRestPublicSymbols:              "/v2/public/symbols",              // GET 	Get the Basic Information of Trading Products. 	This endpoint retrieves the basic configuration and rules for all tradable products on the platform.
	PublicRestPublicBaseRates:            "/v1/public/baseRates",            // GET 	Get Margin Interest Rates. 											This endpoint retrieves real-time information on margin interest rates.
	PublicRestPublicSpotMarginCollateral: "/v1/public/spotMarginCollateral", // GET 	Get Spot Leverage Parameters 										Get the leverage parameters for a currency, including leverage rate and maintenance margin.
	PublicRestPublicHaircut:              "/v1/public/haircut",              // GET 	Get Currency Asset Discount Ratio Information 	Get asset discount ratio.
}
