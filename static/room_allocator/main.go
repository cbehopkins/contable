package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

func calculate(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	namesBox := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	roomCount, _ := strconv.Atoi(value1)

	//js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)

	fmt.Println("You have asked for ", roomCount, " Rooms")
	fmt.Println("The people in the room are:")
	names := strings.Split(namesBox, "\n")
	for i, name := range names {
		fmt.Println(i, "Name is:", name)
	}
	return nil
}

func registerCallbacks() {
	js.Global().Set("calculate", js.FuncOf(calculate))
}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
