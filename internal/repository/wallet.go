package repository

import (
	"math/big"
	"time"

	"github.com/VQIVS/web3-tracker.git/api/handlers/presenter"
	"github.com/VQIVS/web3-tracker.git/internal/entities"
	"github.com/VQIVS/web3-tracker.git/pkg/common"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) CreateWallet(wallet *entities.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *WalletRepository) UpdateWallet(wallet *entities.Wallet) error {
	return r.db.Save(wallet).Error
}

func (r *WalletRepository) UpsertWallet(address, label, balanceWei, balanceETH string) error {
	var wallet entities.Wallet
	err := r.db.Where("address = ?", address).First(&wallet).Error

	if err == gorm.ErrRecordNotFound {
		wallet = entities.Wallet{
			Address:    address,
			Label:      label,
			BalanceWei: balanceWei,
			BalanceETH: balanceETH,
		}
		return r.CreateWallet(&wallet)
	} else if err != nil {
		return err
	}
	wallet.Label = label
	wallet.BalanceWei = balanceWei
	wallet.BalanceETH = balanceETH
	return r.UpdateWallet(&wallet)
}

func (r *WalletRepository) GetWalletByAddress(address string) (*entities.Wallet, error) {
	var wallet entities.Wallet
	err := r.db.Where("address = ?", address).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) GetAllWallets() ([]entities.Wallet, error) {
	var wallets []entities.Wallet
	err := r.db.Order("updated_at DESC").Find(&wallets).Error
	return wallets, err
}

func (r *WalletRepository) DeleteWallet(address string) error {
	return r.db.Unscoped().Where("address = ?", address).Delete(&entities.Wallet{}).Error
}

func (r *WalletRepository) GetPortfolioStatus() (*presenter.PortfolioStatus, error) {
	wallets, err := r.GetAllWallets()
	if err != nil {
		return nil, err
	}

	var totalWei big.Int
	var lastUpdated time.Time

	for _, wallet := range wallets {
		wei, ok := new(big.Int).SetString(wallet.BalanceWei, 10)
		if ok {
			totalWei.Add(&totalWei, wei)
		}
		if wallet.UpdatedAt.After(lastUpdated) {
			lastUpdated = wallet.UpdatedAt
		}
	}

	return &presenter.PortfolioStatus{
		TotalWallets:   len(wallets),
		TotalBalance:   common.WeiToETH(&totalWei),
		LastUpdated:    lastUpdated,
		WalletBalances: wallets,
	}, nil
}

func (r *WalletRepository) GetWalletCount() (int64, error) {
	var count int64
	err := r.db.Model(&entities.Wallet{}).Count(&count).Error
	return count, err
}
