# Web3 Wallet Tracker

A simple and elegant Web3 wallet tracker built with Go and Fiber that allows you to monitor Ethereum wallet balances in real-time.

##  Features

- **Portfolio Dashboard** - View total balance across all tracked wallets
- **Wallet Management** - Add, remove, and label your Ethereum wallets
- **Real-time Updates** - Fetch latest balances from Ethereum mainnet
- **SQLite Database** - Lightweight local storage for wallet data

## Quick Start

### Prerequisites

- Go 1.24+ installed
- Ethereum RPC endpoint (Infura, Alchemy, or public node)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/VQIVS/web3-tracker.git
   cd web3-tracker
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Configure your RPC endpoint**
   
   Edit `config.yaml` and replace `YOUR_PROJECT_ID` with your actual Infura project ID:
   ```yaml
   ethereum:
     rpc_url: "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"
   ```

4. **Run the application**
   ```bash
   go run cmd/main.go
   ```

5. **Open your browser**
   
   Navigate to `http://localhost:8080`

## 📁 Project Structure

```
web3-tracker/
├── cmd/
│   └── main.go              # Application entry point
├── app/
│   └── app.go              # Application setup and routing
├── api/handlers/
│   ├── wallet.go           # API handlers
│   ├── frontend.go         # Frontend rendering handlers
│   └── presenter/          # Response models
├── config/
│   ├── config.go           # Configuration management
│   └── read.go             # Wallet file reading utilities
├── internal/
│   ├── entities/           # Database models
│   └── repository/         # Database operations
├── service/geth/
│   └── geth.go             # Ethereum client service
├── pkg/
│   ├── sqlite/             # Database setup
│   └── common/             # Utility functions
├── views/                  # HTML templates
│   ├── layout.html         # Base layout
│   ├── dashboard.html      # Dashboard page
├── config.yaml             # Application configuration
└── README.md
```

## Configuration

The application uses `config.yaml` for configuration:

```yaml
ethereum:
  rpc_url: "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"

database:
  path: "./sqlite.db"

server:
  port: "8080"

wallets:
  file_path: "./wallets.json"
```

### RPC Providers

You can use any of these Ethereum RPC providers:

- **Infura**: `https://mainnet.infura.io/v3/YOUR_PROJECT_ID`
- **Alchemy**: `https://eth-mainnet.alchemyapi.io/v2/YOUR_API_KEY`
- **Cloudflare**: `https://cloudflare-eth.com`
- **Public nodes**: Various free public Ethereum nodes

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Dashboard page |
| GET | `/api/portfolio` | Get portfolio status |
| GET | `/api/wallets` | Get all wallets |
| POST | `/api/wallet` | Add new wallet |
| DELETE | `/api/wallet/:address` | Delete wallet |
| GET | `/api/balance/:address` | Get specific wallet balance |
| POST | `/api/update` | Update all wallet balances |

## Usage

### Adding Wallets

1. Navigate to the "Manage Wallets" page
2. Enter a valid Ethereum address (0x...)
3. Optionally add a label for easy identification
4. Click "Add Wallet"

### Updating Balances

- **Single wallet**: Click the refresh button next to any wallet
- **All wallets**: Click "Update All Balances" on either page

### Portfolio Overview

The dashboard shows:
- Total number of tracked wallets
- Combined ETH balance across all wallets
- Individual wallet balances and last update times

## Development

### Running in Development

```bash
# Run with auto-reload (requires air)
go install github.com/cosmtrek/air@latest
air

# Or run directly
go run cmd/main.go
```

### Building for Production

```bash
# Build binary
go build -o web3-tracker cmd/main.go

# Run binary
./web3-tracker
```

## Database

The application uses SQLite for local storage. The database file (`sqlite.db`) is created automatically on first run.

### Wallet Schema

```sql
CREATE TABLE wallets (
    id INTEGER PRIMARY KEY,
    address TEXT UNIQUE NOT NULL,
    label TEXT,
    balance_wei TEXT NOT NULL,
    balance_eth TEXT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);
```
