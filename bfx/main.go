package main

import (
	"context"
	"time"

	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/labstack/gommon/log"
)

func main() {
	c := bitfinex.NewClient()

	err := c.Websocket.Connect()
	if err != nil {
		log.Fatal("failed to connect websocket : ", err)
	}
	c.Websocket.SetReadTimeout(time.Second * 2)

	c.Websocket.AttachEventHandler(func(ev interface{}) {
		log.Printf("EVENT: %#v\n", ev)
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	msg := &bitfinex.PublicSubscriptionRequest{
		Event:   "subscribe",
		Channel: bitfinex.ChanTicker,
		Symbol:  bitfinex.TradingPrefix + bitfinex.BTCUSD,
	}

	err = c.Websocket.Subscribe(ctx, msg, func(ev interface{}) {
		switch v := ev.(type) {
		case [][]float64:
			log.Printf("  BID: %#v (%#v)\n", v[0][0], v[0][1])
			log.Printf("  ASK: %#v (%#v)\n", v[0][2], v[0][3])
			log.Printf("  CHANGE: %#v (%#v)\n", v[0][4], v[0][5])
			log.Printf("  LAST: %#v\n", v[0][6])
			log.Printf("  VOL: %#v\n", v[0][7])
			log.Printf("  HIGH: %#v\n", v[0][8])
			log.Printf("  LOW: %#v\n", v[0][9])
		default:
			log.Fatalf("invalid event : %#v\n", v)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-c.Websocket.Done():
			log.Printf("channel closed : %s\n", c.Websocket.Err())
			return
		default:

		}
	}
}
