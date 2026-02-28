package myxcoinapi

type PrivateRestAccountBalanceAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountBalanceReq
}

type PrivateRestAccountBalanceReq struct {
	AccountName  *string `json:"accountName"`  // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	CurrencyList *string `json:"currencyList"` // false 币种列表，使用英文逗号分隔，如 BTC,USDT
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestAccountBalanceAPI) AccountName(accountName string) *PrivateRestAccountBalanceAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 币种列表，使用英文逗号分隔，如 BTC,USDT
func (api *PrivateRestAccountBalanceAPI) CurrencyList(currencyList string) *PrivateRestAccountBalanceAPI {
	api.req.CurrencyList = GetPointer(currencyList)
	return api
}

type PrivateRestAccountTransferBalanceAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountTransferBalanceReq
}

type PrivateRestAccountTransferBalanceReq struct {
	AccountName  *string `json:"accountName"`  // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	CurrencyList *string `json:"currencyList"` // false 币种列表，使用英文逗号分隔，如 BTC,USDT
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestAccountTransferBalanceAPI) AccountName(accountName string) *PrivateRestAccountTransferBalanceAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 币种列表，使用英文逗号分隔，如 BTC,USDT
func (api *PrivateRestAccountTransferBalanceAPI) CurrencyList(currencyList string) *PrivateRestAccountTransferBalanceAPI {
	api.req.CurrencyList = GetPointer(currencyList)
	return api
}

type PrivateRestAccountAvailableBalanceAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountAvailableBalanceReq
}

type PrivateRestAccountAvailableBalanceReq struct {
	AccountName *string `json:"accountName"` // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Symbol      *string `json:"symbol"`      // true 交易标的名称，仅支持现货
	Side        *string `json:"side"`        // true 交易方向，buy，sell
	Price       *string `json:"price"`       // false 委托价格 当为空时，按当前最新成交价计算最大可用
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestAccountAvailableBalanceAPI) AccountName(accountName string) *PrivateRestAccountAvailableBalanceAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string true 交易标的名称，仅支持现货
func (api *PrivateRestAccountAvailableBalanceAPI) Symbol(symbol string) *PrivateRestAccountAvailableBalanceAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string true 交易方向，buy，sell
func (api *PrivateRestAccountAvailableBalanceAPI) Side(side string) *PrivateRestAccountAvailableBalanceAPI {
	api.req.Side = GetPointer(side)
	return api
}

// string false 委托价格 当为空时，按当前最新成交价计算最大可用
func (api *PrivateRestAccountAvailableBalanceAPI) Price(price string) *PrivateRestAccountAvailableBalanceAPI {
	api.req.Price = GetPointer(price)
	return api
}
