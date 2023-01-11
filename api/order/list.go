package order

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api"
)

func List(ctx context.Context, symbol string) ([]*binance.Order, error) {
	orders, err := api.Client.NewListOrdersService().
		Symbol(symbol).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return orders, nil
}

func ListOpen(ctx context.Context, symbol string) ([]*binance.Order, error) {
	orders, err := api.Client.NewListOpenOrdersService().
		Symbol(symbol).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return orders, nil
}
