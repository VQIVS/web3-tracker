package entities

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Address    string         `gorm:"unique;not null;index" json:"address"`
	Label      string         `gorm:"size:255" json:"label,omitempty"`
	BalanceWei string         `gorm:"not null" json:"balance_wei"`
	BalanceETH string         `gorm:"not null" json:"balance_eth"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type PortfolioStatus struct {
	TotalWallets   int       `json:"total_wallets"`
	TotalBalance   string    `json:"total_balance_eth"`
	LastUpdated    time.Time `json:"last_updated"`
	WalletBalances []Wallet  `json:"wallet_balances"`
}
