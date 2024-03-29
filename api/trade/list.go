package trade

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api"
)

func List(ctx context.Context, symbol string) ([]*binance.TradeV3, error) {
	res, err := api.Client.NewListTradesService().
		Symbol(symbol).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}
