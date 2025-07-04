package common

import "math/big"

func WeiToETH(wei *big.Int) string {
	ethValue := new(big.Float).SetInt(wei)
	ethValue.Quo(ethValue, big.NewFloat(1e18))
	return ethValue.Text('f', 6)
}
