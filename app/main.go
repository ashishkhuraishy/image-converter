package main

import (
	"encoding/base64"
	"fmt"
	"strings"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("Web assembly")
	js.Global().Set("update", js.FuncOf(update))
	js.Global().Set("loadImage", js.FuncOf(loadImage))
	<-c
}

func update(this js.Value, args []js.Value) interface{} {
	doc := js.Global().Get("document")
	data := doc.Call("getElementById", "data")
	data.Set("innerHTML", "<h1>Hello From Golang</h1>")

	return nil
}

func loadImage(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return nil
	}

	fmt.Println(args[0])

	data := args[0]
	splitted := strings.Split(data.String(), ",")
	if len(splitted) < 2 {
		fmt.Println("invalid data")
		return nil
	}

	_ = splitted[0]
	encoding := splitted[1]

	unbased, err := base64.StdEncoding.DecodeString(encoding)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println(unbased)

	return nil
}
