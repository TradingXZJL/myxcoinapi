package myxcoinapi

// GET 获取账户余额 获取账户余额
func (client *PrivateRestClient) NewPrivateRestAccountBalance() *PrivateRestAccountBalanceAPI {
	return &PrivateRestAccountBalanceAPI{
		client: client,
		req:    &PrivateRestAccountBalanceReq{},
	}
}

func (api *PrivateRestAccountBalanceAPI) Do() (*XcoinRestRes[PrivateRestAccountBalanceRes], error) {
	url := xcoinHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestAccountBalance])
	return xcoinCallApiWithSecret[PrivateRestAccountBalanceRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取交易账户最大可转余额 查询交易账户指定币种的可划转余额
func (client *PrivateRestClient) NewPrivateRestAccountTransferBalance() *PrivateRestAccountTransferBalanceAPI {
	return &PrivateRestAccountTransferBalanceAPI{
		client: client,
		req:    &PrivateRestAccountTransferBalanceReq{},
	}
}

func (api *PrivateRestAccountTransferBalanceAPI) Do() (*XcoinRestRes[PrivateRestAccountTransferBalanceRes], error) {
	url := xcoinHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestAccountTransferBalance])
	return xcoinCallApiWithSecret[PrivateRestAccountTransferBalanceRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取交易账户可用余额 查询交易账户指定币种的可用余额
func (client *PrivateRestClient) NewPrivateRestAccountAvailableBalance() *PrivateRestAccountAvailableBalanceAPI {
	return &PrivateRestAccountAvailableBalanceAPI{
		client: client,
		req:    &PrivateRestAccountAvailableBalanceReq{},
	}
}

func (api *PrivateRestAccountAvailableBalanceAPI) Do() (*XcoinRestRes[PrivateRestAccountAvailableBalanceRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAccountAvailableBalance])
	return xcoinCallApiWithSecret[PrivateRestAccountAvailableBalanceRes](api.client.c, url, NIL_REQBODY, GET)
}