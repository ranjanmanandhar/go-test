package api

type CoinBalanceParam struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int

	CoinBalance int64
}

type error struct {
	Code    int
	Message string
}
