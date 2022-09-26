package main

import (
	"dto"
)

// Task: create Processors for HttpMessage and for WebsocketMessage
// Processors must have method ProcessMessage(msg T) with call of fmt.Println("Message processed")
func main() {
	httpProcessor := dto.HttpProcessor[string, dto.HttpMessage[string]]{}
	msg := dto.HttpMessage[string]{
		Content: "qwerty",
	}
	httpProcessor.ProcessMessage(msg)

	websocketProcessor := dto.WebsocketProcessor[[]byte, dto.WebsocketMessage[[]byte]]{}
	msg2 := dto.WebsocketMessage[[]byte]{
		Content: []byte("asdfgh"),
	}
	websocketProcessor.ProcessMessage(msg2)
}
