package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *WalletHandler) RenderDashboard(c *fiber.Ctx) error {
	portfolio, err := h.walletRepo.GetPortfolioStatus()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error loading portfolio")
	}

	return c.Render("dashboard", fiber.Map{
		"Title":     "Web3 Wallet Tracker - Dashboard",
		"Portfolio": portfolio,
	})
}

func (h *WalletHandler) RenderWallets(c *fiber.Ctx) error {
	wallets, err := h.walletRepo.GetAllWallets()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error loading wallets")
	}

	return c.Render("wallets", fiber.Map{
		"Title":   "Web3 Wallet Tracker - Manage Wallets",
		"Wallets": wallets,
	})
}
