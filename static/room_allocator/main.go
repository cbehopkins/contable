package main

import (
	"fmt"
	"github.com/cbehopkins/room_allocation"
	"strconv"
	"strings"
	"syscall/js"
)

func calculate(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	namesBox := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	setOutput := func(v string) {
		js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", v)
	}
	roomCount, err := strconv.Atoi(value1)
	if err != nil {
		fmt.Println("Received conversion error:", err)
		setOutput("Please supply valid Room input")
		return nil
	}

	names := strings.Split(namesBox, "\n")

	if len(names) < 4 {
		setOutput("Please supply at least 4 people")
		return nil
	}

	optCnt := 8
	meetCnt := 1

	peeps := room_allocation.NewPeople(names)
	roomsSchedule, err := peeps.ToMeeting().OptimalMeet(roomCount, meetCnt, 1<<optCnt)
	if err != nil {
		fmt.Println("Error!", err)
		setOutput("Unexpected Errpr")
		return nil
	}
	setOutput(roomsSchedule.String())
	//t.Println(peeps.ListConnections())

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
