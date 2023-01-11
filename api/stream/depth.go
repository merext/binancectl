package stream

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
)

func Depth(symbol string) {
	wsDepthHandler := func(event *binance.WsDepthEvent) {
		b, err := json.Marshal(event)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	}

	errHandler := func(err error) {
		fmt.Println(err)
	}

	doneC, _, err := binance.WsDepthServe(symbol, wsDepthHandler, errHandler)
	if err != nil {
		fmt.Print(err)
		return
	}

	// use stopC to exit
	go func() {
		time.Sleep(5 * time.Second)
	}()
	<-doneC
}
