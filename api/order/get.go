package order

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api"
)

func Get(ctx context.Context, symbol string, orderId int64) (*binance.Order, error) {
	order, err := api.Client.NewGetOrderService().
		Symbol(symbol).
		OrderID(orderId).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return order, nil
}
