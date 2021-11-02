package trade

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api"
)

func List(symbol string) ([]*binance.SymbolPrice, error) {
	prices, err := api.Client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return prices, nil
}
