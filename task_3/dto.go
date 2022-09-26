package dto

import "fmt"

type Message[T any] interface {
	GetContent() T
}

type HttpMessage[T any] struct {
	Content T
}

type WebsocketMessage[T any] struct {
	Content T
}

func (hm HttpMessage[T]) GetContent() T {
	return hm.Content
}

func (wm WebsocketMessage[T]) GetContent() T {
	return wm.Content
}

type Processor[T any, M Message[T]] interface {
	ProcessMessage(msg M)
}

type HttpProcessor[T any, M HttpMessage[T]] struct {
}

type WebsocketProcessor[T any, M WebsocketMessage[T]] struct {
}

func (hp HttpProcessor[T, M]) ProcessMessage(msg M) {
	fmt.Println("Message processed: http")
}

func (wp WebsocketProcessor[T, M]) ProcessMessage(msg M) {
	fmt.Println("Message processed: websocket")
}
