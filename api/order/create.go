package order

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api"
)

func Create(symbol string, quantity string, price string) (*binance.CreateOrderResponse, error) {
	fmt.Println(">>>", symbol, quantity, price)
	order, err := api.Client.NewCreateOrderService().
		Symbol(symbol).
		Side(binance.SideTypeBuy).
		Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity(quantity).
		Price(price).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return order, nil
}
