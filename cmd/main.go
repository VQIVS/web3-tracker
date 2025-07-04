package main

import "github.com/VQIVS/web3-tracker.git/app"

func main() {
	app, err := app.NewApp("config.yaml")
	if err != nil {
		panic(err)
	}
	if err := app.Start(); err != nil {
		panic(err)
	}
}
