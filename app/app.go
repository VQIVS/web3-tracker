package app

import (
	"github.com/VQIVS/web3-tracker.git/api/handlers"
	"github.com/VQIVS/web3-tracker.git/config"
	"github.com/VQIVS/web3-tracker.git/internal/repository"
	"github.com/VQIVS/web3-tracker.git/pkg/sqlite"
	"github.com/VQIVS/web3-tracker.git/service/geth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type App struct {
	config   config.Config
	handlers *handlers.WalletHandler
	fiber    *fiber.App
}

func NewApp(cfgPath string) (*App, error) {
	cfg := config.MustLoad(cfgPath)

	// Create template engine with custom functions
	engine := html.New("./views", ".html")
	engine.AddFunc("slice", func(s string, start, end int) string {
		if len(s) == 0 {
			return ""
		}
		if start < 0 {
			start = len(s) + start
		}
		if end <= 0 {
			end = len(s) + end
		}
		if start < 0 {
			start = 0
		}
		if end > len(s) {
			end = len(s)
		}
		if start >= end || start >= len(s) {
			return ""
		}
		return s[start:end]
	})

	httpServer := fiber.New(fiber.Config{
		Views: engine,
	})

	httpServer.Static("/static", "./web/static")

	db, err := sqlite.SetupDatabase(cfg.Database.Path)
	if err != nil {
		return nil, err
	}

	ethService, err := geth.NewEthereumService(cfg.Ethereum.RPCURL)
	if err != nil {
		return nil, err
	}

	handlers := handlers.NewWalletHandler(
		repository.NewWalletRepository(db.DB),
		ethService,
	)

	return &App{
		config:   cfg,
		handlers: handlers,
		fiber:    httpServer,
	}, nil
}

func (a *App) SetupRoutes() {
	// Frontend routes
	a.fiber.Get("/", a.handlers.RenderWallets)
	// a.fiber.Get("/wallets", a.handlers.RenderWallets)

	api := a.fiber.Group("/api")
	api.Get("/portfolio", a.handlers.GetPortfolioStatus)
	api.Get("/update", a.handlers.UpdateAllBalances)
	api.Post("/update", a.handlers.UpdateAllBalances)
	api.Get("/balance/:address", a.handlers.GetBalance)
	api.Get("/wallets", a.handlers.GetAllWallets)
	api.Post("/wallet", a.handlers.AddWallet)
	api.Delete("/wallet/:address", a.handlers.DeleteWallet)
}

func (a *App) Start() error {
	a.SetupRoutes()
	return a.fiber.Listen(":" + a.config.Server.Port)
}
