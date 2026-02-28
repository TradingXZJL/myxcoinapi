package myxcoinapi

type PublicRestPublicSymbolsResRow struct {
	BusinessType      string   `json:"businessType"`
	Symbol            string   `json:"symbol"`
	SymbolFamily      string   `json:"symbolFamily"`
	QuoteCurrency     string   `json:"quoteCurrency"`
	BaseCurrency      string   `json:"baseCurrency"`
	SettleCurrency    string   `json:"settleCurrency"`
	CtVal             string   `json:"ctVal"`
	TickSize          string   `json:"tickSize"`
	Group             []string `json:"group"`
	Status            string   `json:"status"`
	ContractType      string   `json:"contractType"`
	DeliveryTime      string   `json:"deliveryTime"`
	DeliveryFeeRate   string   `json:"deliveryFeeRate"`
	PricePrecision    string   `json:"pricePrecision"`
	QuantityPrecision string   `json:"quantityPrecision"`
	OnlineTime        string   `json:"onlineTime"`
	RiskEngineRate    string   `json:"riskEngineRate"`
	MaxLeverage       string   `json:"maxLeverage"`
	OrderParameters   struct {
		MinOrderQty        string `json:"minOrderQty"`
		MinOrderAmt        string `json:"minOrderAmt"`
		MaxOrderNum        string `json:"maxOrderNum"`
		MaxTriggerOrderNum string `json:"maxTriggerOrderNum"`
		MaxTpslOrderNum    string `json:"maxTpslOrderNum"`
		MaxLmtOrderAmt     string `json:"maxLmtOrderAmt"`
		MaxMktOrderAmt     string `json:"maxMktOrderAmt"`
		MaxLmtOrderQty     string `json:"maxLmtOrderQty"`
		MaxMktOrderQty     string `json:"maxMktOrderQty"`
	} `json:"orderParameters"`
	PositionParameters struct {
		PositionRatioThreshold string `json:"positionRatioThreshold"`
		PositionMaxRatio       string `json:"positionMaxRatio"`
		PositionCidMaxRatio    string `json:"positionCidMaxRatio"`
		DefaultLeverRatio      string `json:"defaultLeverRatio"`
	} `json:"positionParameters"`
	PriceParameters struct {
		MaxLmtPriceUp   string `json:"maxLmtPriceUp"`
		MinLmtPriceDown string `json:"minLmtPriceDown"`
		MaxMktPriceUp   string `json:"maxMktPriceUp"`
		MinMktPriceDown string `json:"minMktPriceDown"`
	} `json:"priceParameters"`
}

type PublicRestPublicSymbolsRes []PublicRestPublicSymbolsResRow

type PublicRestPublicBaseRatesResRow struct {
	Currency       string `json:"currency"`       // Currency symbol, e.g., BTC
	Borrowed       string `json:"borrowed"`       // Borrowed amount
	RemainingQuota string `json:"remainingQuota"` // Remaining borrowable amount
	Rate           string `json:"rate"`           // Annualized borrowing interest rate
}

type PublicRestPublicBaseRatesRes []PublicRestPublicBaseRatesResRow

type PublicRestPublicSpotMarginCollateralResRow struct {
	Currency       string `json:"currency"` // Currency, such as BTC
	LeverageLadder []struct {
		MaxQty   string `json:"maxQty"`   // Gradient upper limit, currency unit (such as BTC), "" indicates positive infinity
		MinQty   string `json:"minQty"`   // Gradient lower limit, currency unit (such as BTC), minimum value 0
		MaxLevel string `json:"maxLevel"` // According to the current quantity upper limit, support maximum leverage
		MmRate   string `json:"mmRate"`   // Maintenance margin rate
	} `json:"leverageLadder"` // Leverage ladder parameters
}

type PublicRestPublicSpotMarginCollateralRes []PublicRestPublicSpotMarginCollateralResRow

type PublicRestPublicHaircutResRow struct {
	Currency        string `json:"currency"` // Currency, such as BTC
	LeverageHaircut []struct {
		MaxQty  string `json:"maxQty"`  // Upper limit of the gradient, currency unit (such as BTC), "" indicates positive infinity
		MinQty  string `json:"minQty"`  // Lower limit of the gradient, currency unit (such as BTC), minimum value 0
		Haircut string `json:"haircut"` // According to the current upper limit quantity, supports maximum leverage
	} `json:"leverageHaircut"` // Discount step parameters
}

type PublicRestPublicHaircutRes []PublicRestPublicHaircutResRow