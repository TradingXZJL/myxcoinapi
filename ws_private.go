package myxcoinapi

import "errors"

func getPositionArg(businessType string, symbol string) WsSubscribeArg {
	return WsSubscribeArg{
		Stream:       "position",
		BusinessType: businessType,
		Symbol:       symbol,
	}
}

// Symbol 交易对，如 BTC-USDT，不填可订阅所有交易对
func (ws *PrivateWsStreamClient) SubscribePosition(businessType string, symbols ...string) (*Subscription[WsPosition], error) {
	if businessType == "" {
		return nil, errors.New("businessType is required")
	}

	// 如果未鉴权，则先鉴权
	if !ws.isAuth {
		err := ws.Auth()
		if err != nil {
			log.Error("Auth error: ", err)
			return nil, err
		}
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getPositionArg(businessType, ""))
	} else {
		for _, s := range symbols {
			args = append(args, getPositionArg(businessType, s))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, SUBSCRIBE)
	if err != nil {
		return nil, err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}
	log.Infof("SubscribePosition Success: args:%v", doSub.Args)

	sub := &Subscription[WsPosition]{
		SubId:        doSub.SubId,
		Ws:           &ws.WsStreamClient,
		Event:        SUBSCRIBE,
		Args:         args,
		resultChan:   make(chan WsPosition),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		subResultMap: make(map[string]bool),
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(arg)
		ws.positionSubMap.Store(string(keyData), sub)
	}

	return sub, nil
}

func (ws *PrivateWsStreamClient) UnsubscribePosition(businessType string, symbols ...string) error {
	if businessType == "" {
		return errors.New("businessType is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getPositionArg(businessType, ""))
	} else {
		for _, s := range symbols {
			args = append(args, getPositionArg(businessType, s))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, UNSUBSCRIBE)
	if err != nil {
		return err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}
	log.Infof("UnsubscribePosition Success: args:%v", doSub.Args)

	for _, arg := range args {
		doSub.Ws.sendUnSubscribeSuccessToCloseChan([]WsSubscribeArg{arg})
		keyData, _ := json.Marshal(arg)
		ws.positionSubMap.Delete(string(keyData))
	}

	return nil
}

func getOrderArg(businessType string, symbol string) WsSubscribeArg {
	return WsSubscribeArg{
		Stream:       "order",
		BusinessType: businessType,
		Symbol:       symbol,
	}
}

// Symbol 交易对，如 BTC-USDT，不填可订阅所有交易对
func (ws *PrivateWsStreamClient) SubscribeOrder(businessType string, symbols ...string) (*Subscription[WsOrder], error) {
	if businessType == "" {
		return nil, errors.New("businessType is required")
	}

	if !ws.isAuth {
		err := ws.Auth()
		if err != nil {
			log.Error("Auth error: ", err)
			return nil, err
		}
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getOrderArg(businessType, ""))
	} else {
		for _, s := range symbols {
			args = append(args, getOrderArg(businessType, s))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, SUBSCRIBE)
	if err != nil {
		return nil, err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}

	sub := &Subscription[WsOrder]{
		SubId:        doSub.SubId,
		Ws:           &ws.WsStreamClient,
		Event:        SUBSCRIBE,
		Args:         args,
		resultChan:   make(chan WsOrder),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		subResultMap: make(map[string]bool),
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(arg)
		ws.orderSubMap.Store(string(keyData), sub)
	}

	return sub, nil
}

func (ws *PrivateWsStreamClient) UnSubscribeOrder(businessType string, symbols ...string) error {
	if businessType == "" {
		return errors.New("businessType is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getOrderArg(businessType, ""))
	} else {
		for _, s := range symbols {
			args = append(args, getOrderArg(businessType, s))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, UNSUBSCRIBE)
	if err != nil {
		return err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}

	log.Infof("UnSubscribeOrder Success: args:%v", doSub.Args)

	for _, arg := range args {
		doSub.Ws.sendUnSubscribeSuccessToCloseChan([]WsSubscribeArg{arg})
		keyData, _ := json.Marshal(arg)
		ws.orderSubMap.Delete(string(keyData))
	}

	return nil
}

func getTradingAccountArg() WsSubscribeArg {
	return WsSubscribeArg{
		Stream: "trading_account",
	}
}

func (ws *PrivateWsStreamClient) SubscribeTradingAccount() (*Subscription[WsTradingAccount], error) {
	if !ws.isAuth {
		err := ws.Auth()
		if err != nil {
			log.Error("Auth error: ", err)
			return nil, err
		}
	}
	arg := getTradingAccountArg()
	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, []WsSubscribeArg{arg}, SUBSCRIBE)
	if err != nil {
		return nil, err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}

	sub := &Subscription[WsTradingAccount]{
		SubId:        doSub.SubId,
		Ws:           &ws.WsStreamClient,
		Event:        SUBSCRIBE,
		Args:         []WsSubscribeArg{arg},
		resultChan:   make(chan WsTradingAccount),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		subResultMap: make(map[string]bool),
	}

	keyData, _ := json.Marshal(arg)
	ws.tradingAccountSubMap.Store(string(keyData), sub)

	return sub, nil
}

func (ws *PrivateWsStreamClient) UnsubscribeTradingAccount() error {
	if !ws.isAuth {
		err := ws.Auth()
		if err != nil {
			log.Error("Auth error: ", err)
			return err
		}
	}
	arg := getTradingAccountArg()
	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, []WsSubscribeArg{arg}, UNSUBSCRIBE)
	if err != nil {
		return err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}

	log.Infof("UnsubscribeTradingAccount Success: args:%v", doSub.Args)

	doSub.Ws.sendUnSubscribeSuccessToCloseChan([]WsSubscribeArg{arg})
	keyData, _ := json.Marshal(arg)
	ws.tradingAccountSubMap.Delete(string(keyData))

	return nil
}
