package handlers

import (
	"fmt"

	"github.com/VQIVS/web3-tracker.git/api/handlers/presenter"
	"github.com/VQIVS/web3-tracker.git/internal/repository"
	"github.com/VQIVS/web3-tracker.git/pkg/common"
	"github.com/VQIVS/web3-tracker.git/service/geth"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
)

type WalletHandler struct {
	walletRepo *repository.WalletRepository
	ethService *geth.EthereumService
}

func NewWalletHandler(walletRepo *repository.WalletRepository, ethService *geth.EthereumService) *WalletHandler {
	return &WalletHandler{
		walletRepo: walletRepo,
		ethService: ethService,
	}
}
func (h *WalletHandler) GetPortfolioStatus(c *fiber.Ctx) error {
	portfolio, err := h.walletRepo.GetPortfolioStatus()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presenter.APIResponse{
			Success: false,
			Message: "Failed to retrieve portfolio status",
		})
	}

	return c.JSON(presenter.APIResponse{
		Success: true,
		Message: "Portfolio status retrieved successfully",
		Data:    portfolio,
	})
}

func (h *WalletHandler) UpdateAllBalances(c *fiber.Ctx) error {
	wallets, err := h.walletRepo.GetAllWallets()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presenter.APIResponse{
			Success: false,
			Message: "Failed to retrieve wallets",
		})
	}

	for _, wallet := range wallets {
		balance, err := h.ethService.GetBalance(wallet.Address)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.APIResponse{
				Success: false,
				Message: fmt.Sprintf("Failed to get balance for address %s", wallet.Address),
			})
		}
		balanceETH := common.WeiToETH(balance)
		if err := h.walletRepo.UpsertWallet(wallet.Address, wallet.Label, balance.String(), balanceETH); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.APIResponse{
				Success: false,
				Message: fmt.Sprintf("Failed to update wallet %s", wallet.Address),
			})
		}
	}

	return c.JSON(presenter.APIResponse{
		Success: true,
		Message: "All balances updated successfully",
	})
}

func (h *WalletHandler) GetBalance(c *fiber.Ctx) error {
	address := c.Params("address")
	balance, err := h.ethService.GetBalance(address)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presenter.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to get balance for address %s", address),
		})
	}
	response := presenter.BalanceResponse{
		Address:    address,
		BalanceWei: balance.String(),
		BalanceETH: common.WeiToETH(balance),
	}
	return c.JSON(presenter.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Balance for address %s retrieved successfully", address),
		Data:    response,
	})
}

func (h *WalletHandler) AddWallet(c *fiber.Ctx) error {
	var req presenter.UpdateBalanceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presenter.APIResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}
	if !ethCommon.IsHexAddress(req.Address) {
		return c.Status(fiber.StatusBadRequest).JSON(presenter.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Invalid Ethereum address: %s", req.Address),
		})
	}
	balance, err := h.ethService.GetBalance(req.Address)
	if err != nil {
		fmt.Printf("Error fetching balance for address %s: %v\n", req.Address, err)
		return c.Status(fiber.StatusInternalServerError).JSON(presenter.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to get balance for address %s: %v", req.Address, err),
		})
	}
	balanceETH := common.WeiToETH(balance)
	if err := h.walletRepo.UpsertWallet(req.Address, req.Label, balance.String(), balanceETH); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presenter.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to add wallet %s", req.Address),
		})
	}
	return c.JSON(presenter.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Wallet %s added successfully", req.Address),
		Data: presenter.BalanceResponse{
			Address:    req.Address,
			BalanceWei: balance.String(),
			BalanceETH: balanceETH,
		},
	})

}

func (h *WalletHandler) DeleteWallet(c *fiber.Ctx) error {
	address := c.Params("address")
	if err := h.walletRepo.DeleteWallet(address); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presenter.APIResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to delete wallet %s", address),
		})
	}
	return c.JSON(presenter.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Wallet %s deleted successfully", address),
	})

}

func (h *WalletHandler) GetAllWallets(c *fiber.Ctx) error {
	wallets, err := h.walletRepo.GetAllWallets()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presenter.APIResponse{
			Success: false,
			Message: "Failed to retrieve wallets",
		})
	}
	return c.JSON(presenter.APIResponse{
		Success: true,
		Message: "Wallets retrieved successfully",
		Data:    wallets,
	})
}
