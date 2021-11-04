package account

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api"
)

func Info() (*binance.Account, error) {
	res, err := api.Client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var balances []binance.Balance
	for _, b := range res.Balances {
		if b.Free != "0.00000000" && b.Free != "0.00" {
			balances = append(balances, b)
		}
	}
	res.Balances = balances

	return res, nil
}
