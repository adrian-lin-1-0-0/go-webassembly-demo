package main

import (
	"syscall/js"
)

func main() {

	js.Global().Get("WebSocket").New("ws://localhost:8080/echo")

}
