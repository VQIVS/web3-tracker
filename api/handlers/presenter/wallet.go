package presenter

import (
	"time"

	"github.com/VQIVS/web3-tracker.git/internal/entities"
)

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
type PortfolioStatus struct {
	TotalWallets   int               `json:"total_wallets"`
	TotalBalance   string            `json:"total_balance_eth"`
	LastUpdated    time.Time         `json:"last_updated"`
	WalletBalances []entities.Wallet `json:"wallet_balances"`
}
