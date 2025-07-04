package presenter

type BalanceResponse struct {
	Address    string `json:"address"`
	BalanceWei string `json:"balance_wei"`
	BalanceETH string `json:"balance_eth"`
}

type UpdateBalanceRequest struct {
	Address string `json:"address"`
	Label   string `json:"label,omitempty"`
}

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
