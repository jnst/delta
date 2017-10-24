package main

import "golang.org/x/net/websocket"

func main() {

	ws, err := websocket.Dial("wss://socket.etherdelta.com/", "", "http://localhost/")
	if err != nil {
		return err
	}

}
