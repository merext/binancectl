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

	//wgDepthHandler := func(event *binance.WsDepthEvent) {
	//	data, _ := json.MarshalIndent(event, "", "  ")
	//	fmt.Println(string(data))
	//}

	errHandler := func(err error) {
		fmt.Println(err)
	}
	//doneC, _, err := binance.WsDepthServe("BTCUSDT", wgDepthHandler, errHandler)
	//doneC, _, err := binance.WsTradeServe("BTCUSDT", wgTradeHandler, errHandler)
	//doneC, _, err := binance.WsAllMarketsStatServe(wgAllMarketsStatHandler, errHandler)
	doneC, _, err := binance.WsBookTickerServe(symbol, wsBookTickerHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
