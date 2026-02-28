package myxcoinapi

type PrivateRestAccountBalanceRes struct {
	AccountName           string `json:"accountName"`           // 账户名称
	Pid                   string `json:"pid"`                   // 账户唯一ID
	Uid                   string `json:"uid"`                   // 操作员唯一ID
	Cid                   string `json:"cid"`                   // 操作主体唯一ID
	TotalEquity           string `json:"totalEquity"`           // 总权益，单位为USDT
	TotalMarginBalance    string `json:"totalMarginBalance"`    // 总保证金，单位为USDT
	TotalAvailableBalance string `json:"totalAvailableBalance"` // 总可用保证金，单位为USDT
	TotalEffectiveMargin  string `json:"totalEffectiveMargin"`  // 总有效保证金，单位为USDT
	TotalPositionValue    string `json:"totalPositionValue"`    // 总仓位价值，单位为USDT
	TotalIm               string `json:"totalIm"`               // 总初始保证金，单位为USDT
	TotalMm               string `json:"totalMm"`               // 总维持保证金，单位为USDT
	TotalOpenLoss         string `json:"totalOpenLoss"`         // 保证金开平仓损失，单位为USDT
	Mmr                   string `json:"mmr"`                   // 总维持保证金率
	Imr                   string `json:"imr"`                   // 总初始保证金率
	AccountLeverage       string `json:"accountLeverage"`       // 账户杠杆
	ContractUpl           string `json:"contractUpl"`           // 合约未实现盈亏
	TotalUpl              string `json:"totalUpl"`              // 账户层面未实现盈亏
	FlexibleEquity        string `json:"flexibleEquity"`        // 活期权益，USDT估值
	AutoSubscribe         bool   `json:"autoSubscribe"`         // 现金宝开关
	FlexiblePnl           string `json:"flexiblePnl"`           // 活期收益
	Details               []struct {
		Currency               string `json:"currency"`               // 币种，如 BTC
		Equity                 string `json:"equity"`                 // 币种总权益
		TotalBalance           string `json:"totalBalance"`           // 币种余额，包含现金宝部分
		CashBalance            string `json:"cashBalance"`            // 净现金余额，不包含现金宝部分
		SavingBalance          string `json:"savingBalance"`          // 现金宝余额
		LeftPersonalQuota      string `json:"leftPersonalQuota"`      // 现金宝剩余可申购额度
		SavingTotalPnl         string `json:"savingTotalPnl"`         // 现金宝累计收益
		SavingLastPnl          string `json:"savingLastPnl"`          // 昨日收益
		SavingHoldDays         string `json:"savingHoldDays"`         // 目前已持有天数
		SavingTotalAPR         string `json:"savingTotalAPR"`         // 累计年利率
		SavingLastAPR          string `json:"savingLastAPR"`          // 昨日年利率
		RealLiability          string `json:"realLiability"`          // 真实负债
		PotentialLiability     string `json:"potentialLiability"`     // 潜在负债
		AccruedInterest        string `json:"accruedInterest"`        // 累计利息
		Upl                    string `json:"upl"`                    // 未实现盈亏，仅包含合约和交割合约未实现盈亏
		PositionInitialMargin  string `json:"positionInitialMargin"`  // 合约仓位占用保证金
		OrderInitialMargin     string `json:"orderInitialMargin"`     // 订单占用保证金
		Frozen                 string `json:"frozen"`                 // 冻结资产
		LiabilityInitialMargin string `json:"liabilityInitialMargin"` // 负债占用保证金
		InitialMargin          string `json:"initialMargin"`          // 币种占用保证金
	} `json:"details"` // 币种级别的资金数据列表
}

type PrivateRestAccountTransferBalanceResRow struct {
	AccountName string `json:"accountName"` // 账户名称
	Pid         string `json:"pid"`         // 账户唯一ID
	Uid         string `json:"uid"`         // 管理员唯一ID
	Cid         string `json:"cid"`         // 操作主体唯一ID
	Currency    string `json:"currency"`    // 币种，如 BTC
	MaxTransfer string `json:"maxTransfer"` // 最大可转
}

type PrivateRestAccountTransferBalanceRes []PrivateRestAccountTransferBalanceResRow

type PrivateRestAccountAvailableBalanceRes struct {
	AccountName        string `json:"accountName"`        // 账户名称
	Pid                string `json:"pid"`                // 账户唯一ID
	Uid                string `json:"uid"`                // 管理员唯一ID
	Cid                string `json:"cid"`                // 操作主体唯一ID
	Symbol             string `json:"symbol"`             // 交易对名称，如 BTC-USDT
	Side               string `json:"side"`               // 订单方向 buy：买，sell：卖
	MaxBaseCcyBalance  string `json:"maxBaseCcyBalance"`  // 基于基础币种，最大可申报数量
	MaxQuoteCcyBalance string `json:"maxQuoteCcyBalance"` // 基于计价币种，最大可申报数量
}
