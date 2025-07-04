package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type WalletAddress struct {
	Address string `json:"address"`
	Label   string `json:"label,omitempty"`
}

func ReadWalletAddresses(filename string) ([]WalletAddress, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read wallet file %s: %v", filename, err)
	}

	var addresses []WalletAddress
	if err := json.Unmarshal(data, &addresses); err != nil {
		return nil, fmt.Errorf("failed to unmarshal wallet JSON: %v", err)
	}

	return addresses, nil
}
