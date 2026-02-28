package myxcoinapi

// POST 下单 支持普通订单和止盈止损订单
func (client *PrivateRestClient) NewPrivateRestTradeOrder() *PrivateRestTradeOrderAPI {
	return &PrivateRestTradeOrderAPI{
		client: client,
		req:    &PrivateRestTradeOrderReq{},
	}
}

func (api *PrivateRestTradeOrderAPI) Do() (*XcoinRestRes[PrivateRestTradeOrderRes], error) {
	url := xcoinHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeOrder])
	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return xcoinCallApiWithSecret[PrivateRestTradeOrderRes](api.client.c, url, body, POST)
}

// POST 批量下单 支持普通订单和止盈止损订单
func (client *PrivateRestClient) NewPrivateRestTradeBatchOrder() *PrivateRestTradeBatchOrderAPI {
	return &PrivateRestTradeBatchOrderAPI{
		client: client,
		req:    &PrivateRestTradeBatchOrderReq{},
	}
}

func (api *PrivateRestTradeBatchOrderAPI) Do() (*XcoinRestRes[PrivateRestTradeBatchOrderRes], error) {
	url := xcoinHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeBatchOrder])
	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return xcoinCallApiWithSecret[PrivateRestTradeBatchOrderRes](api.client.c, url, body, POST)
}

// POST 取消订单 这里进行撤单下单操作
func (client *PrivateRestClient) NewPrivateRestTradeCancelOrder() *PrivateRestTradeCancelOrderAPI {
	return &PrivateRestTradeCancelOrderAPI{
		client: client,
		req:    &PrivateRestTradeCancelOrderReq{},
	}
}

func (api *PrivateRestTradeCancelOrderAPI) Do() (*XcoinRestRes[PrivateRestTradeCancelOrderRes], error) {
	url := xcoinHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeCancelOrder])
	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return xcoinCallApiWithSecret[PrivateRestTradeCancelOrderRes](api.client.c, url, body, POST)
}

// POST 批量撤单 撤单请求参数为List对象，每个对象定义如下，List对象最大个数为20
func (client *PrivateRestClient) NewPrivateRestTradeBatchCancelOrder() *PrivateRestTradeBatchCancelOrderAPI {
	return &PrivateRestTradeBatchCancelOrderAPI{
		client: client,
		req:    &PrivateRestTradeBatchCancelOrderReq{},
	}
}

func (api *PrivateRestTradeBatchCancelOrderAPI) Do() (*XcoinRestRes[PrivateRestTradeBatchCancelOrderRes], error) {
	url := xcoinHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeBatchCancelOrder])
	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return xcoinCallApiWithSecret[PrivateRestTradeBatchCancelOrderRes](api.client.c, url, body, POST)
}

// POST 撤销所有订单 批量撤所有订单
func (client *PrivateRestClient) NewPrivateRestTradeCancelAllOrder() *PrivateRestTradeCancelAllOrderAPI {
	return &PrivateRestTradeCancelAllOrderAPI{
		client: client,
		req:    &PrivateRestTradeCancelAllOrderReq{},
	}
}

func (api *PrivateRestTradeCancelAllOrderAPI) Do() (*XcoinRestRes[PrivateRestTradeCancelAllOrderRes], error) {
	url := xcoinHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeCancelAllOrder])
	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return xcoinCallApiWithSecret[PrivateRestTradeCancelAllOrderRes](api.client.c, url, body, POST)
}

// GET 获取当前挂单 该接口仅查询未成交订单列表，已成交或已撤单订单通过历史订单接口查询
func (client *PrivateRestClient) NewPrivateRestTradeOpenOrders() *PrivateRestTradeOpenOrdersAPI {
	return &PrivateRestTradeOpenOrdersAPI{
		client: client,
		req:    &PrivateRestTradeOpenOrdersReq{},
	}
}

func (api *PrivateRestTradeOpenOrdersAPI) Do() (*XcoinRestRes[PrivateRestTradeOpenOrdersRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeOpenOrders])
	return xcoinCallApiWithSecret[PrivateRestTradeOpenOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取订单信息 基于订单id获取完整订单信息
func (client *PrivateRestClient) NewPrivateRestTradeOrderInfo() *PrivateRestTradeOrderInfoAPI {
	return &PrivateRestTradeOrderInfoAPI{
		client: client,
		req:    &PrivateRestTradeOrderInfoReq{},
	}
}

func (api *PrivateRestTradeOrderInfoAPI) Do() (*XcoinRestRes[PrivateRestTradeOrderInfoRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeOrderInfo])
	return xcoinCallApiWithSecret[PrivateRestTradeOrderInfoRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取历史订单 该接口仅查询完结态订单，未完结态通过当前订单列表来查询，仅获取最近90天已完成的订单数据，返回数据按照订单创建时间倒序排序。
func (client *PrivateRestClient) NewPrivateRestTradeHistoryOrders() *PrivateRestTradeHistoryOrdersAPI {
	return &PrivateRestTradeHistoryOrdersAPI{
		client: client,
		req:    &PrivateRestTradeHistoryOrdersReq{},
	}
}

func (api *PrivateRestTradeHistoryOrdersAPI) Do() (*XcoinRestRes[PrivateRestTradeHistoryOrdersRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeHistoryOrders])
	return xcoinCallApiWithSecret[PrivateRestTradeHistoryOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取订单操作明细记录 获取最近90天具体订单操作的记录，返回数据按照订单操作时间倒序排序。
func (client *PrivateRestClient) NewPrivateRestTradeOrderOperations() *PrivateRestTradeOrderOperationsAPI {
	return &PrivateRestTradeOrderOperationsAPI{
		client: client,
		req:    &PrivateRestTradeOrderOperationsReq{},
	}
}

func (api *PrivateRestTradeOrderOperationsAPI) Do() (*XcoinRestRes[PrivateRestTradeOrderOperationsRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeOrderOperations])
	return xcoinCallApiWithSecret[PrivateRestTradeOrderOperationsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 账户成交历史 获取最近90天已成交订单明细数据，返回数据按照成交时间倒序排序。
func (client *PrivateRestClient) NewPrivateRestTradeHistoryTrades() *PrivateRestTradeHistoryTradesAPI {
	return &PrivateRestTradeHistoryTradesAPI{
		client: client,
		req:    &PrivateRestTradeHistoryTradesReq{},
	}
}

func (api *PrivateRestTradeHistoryTradesAPI) Do() (*XcoinRestRes[PrivateRestTradeHistoryTradesRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeHistoryTrades])
	return xcoinCallApiWithSecret[PrivateRestTradeHistoryTradesRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取交易账户仓位 仅适用于合约业务线
func (client *PrivateRestClient) NewPrivateRestTradePosition() *PrivateRestTradePositionAPI {
	return &PrivateRestTradePositionAPI{
		client: client,
		req:    &PrivateRestTradePositionReq{},
	}
}

func (api *PrivateRestTradePositionAPI) Do() (*XcoinRestRes[PrivateRestTradePositionRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradePosition])
	return xcoinCallApiWithSecret[PrivateRestTradePositionRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 设置杠杆倍数 设置仓位和订单的币种维度交易杠杆
func (client *PrivateRestClient) NewPrivateRestTradeLeverPost() *PrivateRestTradeLeverPostAPI {
	return &PrivateRestTradeLeverPostAPI{
		client: client,
		req:    &PrivateRestTradeLeverPostReq{},
	}
}

func (api *PrivateRestTradeLeverPostAPI) Do() (*XcoinRestRes[PrivateRestTradeLeverCommonRes], error) {
	url := xcoinHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeLeverPost])
	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return xcoinCallApiWithSecret[PrivateRestTradeLeverCommonRes](api.client.c, url, body, POST)
}

// GET 获取杠杆倍数 获取仓位和订单的币种维度交易杠杆
func (client *PrivateRestClient) NewPrivateRestTradeLeverGet() *PrivateRestTradeLeverGetAPI {
	return &PrivateRestTradeLeverGetAPI{
		client: client,
		req:    &PrivateRestTradeLeverGetReq{},
	}
}

func (api *PrivateRestTradeLeverGetAPI) Do() (*XcoinRestRes[PrivateRestTradeLeverCommonRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeLeverGet])
	return xcoinCallApiWithSecret[PrivateRestTradeLeverCommonRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 设置仓位止盈止损 该接口可以设置仓位止盈止损参数，支持全部仓位设置止盈止损和部分仓位止盈止损
func (client *PrivateRestClient) NewPrivateRestTradeStopPosition() *PrivateRestTradeStopPositionAPI {
	return &PrivateRestTradeStopPositionAPI{
		client: client,
		req:    &PrivateRestTradeStopPositionReq{},
	}
}

func (api *PrivateRestTradeStopPositionAPI) Do() (*XcoinRestRes[PrivateRestTradeStopPositionRes], error) {
	url := xcoinHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeStopPosition])
	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return xcoinCallApiWithSecret[PrivateRestTradeStopPositionRes](api.client.c, url, body, POST)
}
