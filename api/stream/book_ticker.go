package stream

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/utils"
)

func BookTicker(symbol string) {
	wsBookTickerHandler := func(event *binance.WsBookTickerEvent) {
		utils.OPrint(event)
	}

	errHandler := func(err error) {
		fmt.Println(err)
	}
	
	doneC, _, err := binance.WsBookTickerServe(symbol, wsBookTickerHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
