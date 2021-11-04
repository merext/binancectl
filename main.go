package main

import (
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api"
	"github.com/merext/binancectl/v2/cli"
	"github.com/rs/zerolog"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")

	//binance.UseTestnet = true
	api.Client = binance.NewClient(apiKey, secretKey)

	cli.Execute()

}
