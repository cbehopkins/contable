package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/cbehopkins/wordlist"
	"github.com/cbehopkins/ana"
)

func anagram(this js.Value, i []js.Value) interface{} {
	fmt.Println("Here we are in anagram")
	if len(i) != 1 {
		return map[string]interface{}{
			"error": fmt.Sprint("Error len of input wrong", len(i))}
	}
	if i[0].Type() != js.TypeString {
		return map[string]interface{}{
			"error": fmt.Sprint("Error, wrong type", i[0].Type)}
	}
	stringToCalc := i[0].String()

	data, err := wordlist.Asset("data/wordlist.txt")
	if err != nil {
		return map[string]interface{}{
			"error": err,
		}
	}
	results := ana.AnagramWord(stringToCalc, data)
	resultArray := make([]string, len(results))
	for i, result := range results{
		resultArray[i] = string(result)
	}

	tmp, err := json.Marshal(resultArray)
	if err != nil {
		return map[string]interface{}{
			"error": err,
		}
	}
	return map[string]interface{}{
		"error":    "",
		"anagrams": string(tmp),
	}
}

func registerCallbacks() {
	js.Global().Set("anagram", js.FuncOf(anagram))

}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
