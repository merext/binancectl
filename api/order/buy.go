package order

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api"
)

func Buy(ctx context.Context, symbol string, quantity string, price string) (*binance.CreateOrderResponse, error) {
	request := api.Client.NewCreateOrderService().
		Symbol(symbol).
		Side(binance.SideTypeBuy).
		Quantity(quantity)

	if price != "" {
		request = request.
			Type(binance.OrderTypeLimit).
			Price(price).
			TimeInForce(binance.TimeInForceTypeGTC)
	} else {
		request = request.
			Type(binance.OrderTypeMarket)
	}

	order, err := request.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return order, nil
}
