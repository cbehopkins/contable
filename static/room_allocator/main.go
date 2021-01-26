package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/cbehopkins/room_allocation"
)

func calculate(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	namesBox := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	roomCount, err := strconv.Atoi(value1)
	if err != nil {
		fmt.Println("Received conversion error:", err)
		return map[string]interface{}{"error": "Please supply valid Room input"}
	}

	names := strings.Split(namesBox, "\n")

	if len(names) < 4 {
		return map[string]interface{}{"error": "Please supply at least 4 people"}
	}

	optCnt := i[3].Int()
	meetCnt := i[2].Int()

	peeps := room_allocation.NewPeople(names)
	roomsSchedule, err := peeps.ToMeeting().OptimalMeet(roomCount, meetCnt, 1<<optCnt)
	if err != nil {
		fmt.Println("Error!", err)
		return map[string]interface{}{"error": "Unexpected Error"}
	}
	v, err := json.Marshal(roomsSchedule)
	if err != nil {
		return map[string]interface{}{"error": "Error marshalling the structure"}
	}

	return map[string]interface{}{
		"error":       "",
		"meetings":    string(v),
		"connections": peeps.ListConnections(),
	}
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
