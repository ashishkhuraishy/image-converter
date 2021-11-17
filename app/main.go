package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("Web assembly")
	js.Global().Set("update", js.FuncOf(update))
	<-c
}

func update(this js.Value, args []js.Value) interface{} {

	doc := js.Global().Get("document")
	data := doc.Call("getElementById", "data")
	data.Set("innerHTML", "<h1>Hello From Golang</h1>")

	return nil
}
