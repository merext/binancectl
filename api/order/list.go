package order

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api"
)

func List(symbol string) ([]*binance.Order, error) {
	orders, err := api.Client.NewListOrdersService().Symbol(symbol).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return orders, nil
}

func ListOpen(symbol string) ([]*binance.Order, error) {
	orders, err := api.Client.NewListOpenOrdersService().Symbol(symbol).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return orders, nil
}
