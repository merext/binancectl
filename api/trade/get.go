package trade

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api"
)

func Get(symbol string) ([]*binance.BookTicker, error) {
	price, err := api.Client.NewListBookTickersService().Symbol(symbol).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return price, nil
}
