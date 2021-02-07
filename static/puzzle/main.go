package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
	"strconv"

	"github.com/cbehopkins/wordlist"
	"github.com/cbehopkins/ana"
	cntSlv "github.com/cbehopkins/countdown/cnt_slv"

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

func jsonToArrayInt(inputText string) (retArray []int, err error){
	var inputValues []string
	err = json.Unmarshal([]byte(inputText), &inputValues)
	if err != nil {
		return nil, err
	}
	retArray = make([]int, len(inputValues))
	for i, v := range inputValues {
		val, err := strconv.Atoi(v)
		if err != nil{
			return nil, err
		}
		retArray[i] = val
	}
	return retArray, nil
}
func jsonTo2dArrayInt(inputText string) (retArray [][]int, err error){
	var inputValues [][]string
	err = json.Unmarshal([]byte(inputText), &inputValues)
	if err != nil {
		return nil, err
	}
	retArray = make([][]int, len(inputValues))
	for j, rowIn := range inputValues {
		rowOut := make([]int, len(rowIn))
		for i, v := range rowIn {
			val, err := strconv.Atoi(v)
			if err != nil{
				return nil, err
			}
			rowOut[i] = val
		}
		retArray[j] = rowOut
	}
	return retArray, nil
}
func RunCountdown(target int, sources []int) string {
		foundValues := cntSlv.NewNumMap()
		//found_values.SelfTest = true
		foundValues.UseMult = true
		foundValues.PermuteMode = cntSlv.FastMap
		foundValues.SeekShort = false // TBD make this controllable

		fmt.Println("Starting permute")
		returnProofs := foundValues.CountHelper(target, sources)
		for range returnProofs {
			//fmt.Println("Proof Received", v)
		}
		//fmt.Println("Permute Complete", proof_list)
		return foundValues.GetProof(target)

}
func countdown(this js.Value, i []js.Value) interface{} {
	fmt.Println("Here we are in countdown")
	if len(i) != 2 {
		return map[string]interface{}{
			"error": fmt.Sprint("Error len of input wrong", len(i))}
	}

	if i[0].Type() != js.TypeString {
		return map[string]interface{}{
			"error": fmt.Sprint("Error, wrong type", i[0].Type)}
	}

	if i[1].Type() != js.TypeString {
		return map[string]interface{}{
			"error": fmt.Sprint("Error, wrong type", i[1].Type)}
	}

	target, err := strconv.Atoi(i[0].String())
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprint("Error, Target int parse fail", err)}
	}
	fmt.Println("We're being asked to target:", target)
	fmt.Println("Our JASON input is:", i[1].String())
	inputValues, err := jsonToArrayInt(i[1].String())
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprint("Error, input values int parse fail", err)}
	}
	fmt.Println("We have the inputs:", inputValues)
	return map[string]interface{}{
		"error":    "",
		"countdown": RunCountdown(target, inputValues),
	}
}
func sudoku(this js.Value, i []js.Value) interface{} {
	fmt.Println("Here we are in sudoku")
	if len(i) != 1 {
		return map[string]interface{}{
			"error": fmt.Sprint("Error len of input wrong", len(i))}
	}



	return map[string]interface{}{
		"error":    "",
	}
}
func registerCallbacks() {
	js.Global().Set("anagram", js.FuncOf(anagram))
	js.Global().Set("countdown", js.FuncOf(countdown))
	js.Global().Set("sudoku", js.FuncOf(sudoku))

}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
