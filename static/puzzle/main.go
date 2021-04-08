package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"syscall/js"

	"github.com/cbehopkins/ana"
	"github.com/cbehopkins/boggle"

	cntslv "github.com/cbehopkins/countdown/cnt_slv"
	"github.com/cbehopkins/sod"
	"github.com/cbehopkins/wordlist"
)

func anagram(this js.Value, i []js.Value) interface{} {
	// fmt.Println("Here we are in anagram")
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
	for i, result := range results {
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

func jsonToArrayInt(inputText string) (retArray []int, err error) {
	var inputValues []string
	err = json.Unmarshal([]byte(inputText), &inputValues)
	if err != nil {
		return nil, err
	}
	retArray = make([]int, len(inputValues))
	for i, v := range inputValues {
		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		retArray[i] = val
	}
	return retArray, nil
}
func jsonTo2dArrayInt(inputText string) (retArray [][]int, err error) {
	var inputValues [][]string
	fmt.Println("About to unmarshal")
	err = json.Unmarshal([]byte(inputText), &inputValues)
	if err != nil {
		return nil, err
	}
	fmt.Println("Unmarshalled")
	retArray = make([][]int, len(inputValues))
	for j, rowIn := range inputValues {
		rowOut := make([]int, len(rowIn))
		for i, v := range rowIn {
			if v == "" {
				continue
			}
			val, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			rowOut[i] = val
		}
		retArray[j] = rowOut
	}
	fmt.Println("All converted")
	return retArray, nil
}
func runCountdown(target int, sources []int) string {
	foundValues := cntslv.NewNumMap()
	foundValues.SelfTest = false
	foundValues.UseMult = true
	foundValues.SeekShort = false // TBD make this controllable

	fmt.Println("Starting permute")
	returnProofs := foundValues.CountHelper(target, sources)
	for range returnProofs {
		//fmt.Println("Proof Received", v)
		// Here we get the proofs as they are found
	}
	// Here we return the "best" solution
	return foundValues.GetProof(target)
}
func countdown(this js.Value, i []js.Value) interface{} {
	// fmt.Println("Here we are in countdown")
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
	// fmt.Println("We're being asked to target:", target)
	// fmt.Println("Our JASON input is:", i[1].String())
	inputValues, err := jsonToArrayInt(i[1].String())
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprint("Error, input values int parse fail", err)}
	}
	// fmt.Println("We have the inputs:", inputValues)
	ra := runCountdown(target, inputValues)
	return map[string]interface{}{
		"error":     "",
		"countdown": string(ra),
	}
}
func sudoku(this js.Value, i []js.Value) interface{} {
	// fmt.Println("Here we are in sudoku")
	if len(i) != 1 {
		return map[string]interface{}{
			"error": fmt.Sprint("Error len of input wrong", len(i))}
	}
	// TBD rip this out into a test function
	//jsonStr := `[["","","","2","6","","7","",""],["6","8","","","7","","","9",""],["1","9","","","","4","5","",""],["8","2","","1","","","","4",""],["","","4","6","","2","9","",""],["","5","","","","3","","2","8"],["","","9","3","","","","7","4"],["","4","","","5","","","3","6"],["7","","3","","1","8","","",""]]`
	ra, err := jsonTo2dArrayInt(i[0].String())
	// fmt.Println("We have our array now:", ra)
	returnArray, err := runSudoku(ra)
	if err != nil {
		return map[string]interface{}{
			"error": err,
		}
	}
	tmp, err := json.Marshal(returnArray)
	return map[string]interface{}{
		"error":  err,
		"sudoku": string(tmp),
	}
}

func runSudoku(input [][]int) (output [][]int, err error) {
	var testPuzzle *sod.Puzzle
	size := len(input)
	output = make([][]int, size)
	for i, arr := range input {

		if len(arr) != size {
			return
		}
		output[i] = make([]int, size)
	}

	testPuzzle = sod.NewPuzzle()

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			val := input[y][x]
			if val != 0 {
				tc := sod.Coord{x, y}
				testPuzzle.SetVal(sod.Value(val), tc)
			}
		}
	}
	err = testPuzzle.SelfCheck()
	if err != nil {
		return
	}
	testPuzzle.SolveAll()

	err = testPuzzle.SelfCheck()
	if err != nil {
		return
	}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			vals := testPuzzle.GetCel(sod.Coord{x, y}).Values()
			if len(vals) == 1 {
				val := int(vals[0])
				if val != 0 {
					output[y][x] = val
				}
			}
		}
	}
	return
}

func toBoggleArray(input string) ([][]rune, error) {
	var inputValues [][]string
	err := json.Unmarshal([]byte(input), &inputValues)
	if err != nil {
		return nil, err
	}
	for _, v := range inputValues {
		if len(v) != len(inputValues) {
			return nil, errors.New("Input array was not square")
		}
	}
	ra := make([][]rune, len(inputValues))
	for i, v := range inputValues {
		ra[i] = make([]rune, len(inputValues))
		for j, w := range v {
			raa := []rune(w)
			if len(raa) != 1 {
				fmt.Println("Length of rune was not 1, on ", i, ",", j)
			}
			ra[i][j] = raa[0]
		}
	}
	return ra, nil
}

func bogglePromise(this js.Value, argsOuter []js.Value) interface{} {
	// Handler for the Promise
	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		// Run this code asynchronously
		go func(ip string) {
			res, err := boggleRunner(ip)
			if err == nil {
				resolve.Invoke(map[string]interface{}{
					"boggle": res,
					"error":  err,
				})
			} else {
				// Create a JS Error object and pass it to the reject function
				// The constructor for Error accepts a string,
				// so we need to get the error message as string from "err"
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New(err.Error())
				reject.Invoke(errorObject)
			}
		}(argsOuter[0].String())

		// The handler of a Promise doesn't return any value
		return nil
	})

	// Create and return the Promise object
	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func countdownPromise(this js.Value, argsOuter []js.Value) interface{} {
	targetS := argsOuter[0].String()
	inputS := argsOuter[1].String()
	// Handler for the Promise
	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		// Run this code asynchronously
		go func(targetS, inputsS string) {
			errorConstructor := js.Global().Get("Error")

			target, err := strconv.Atoi(targetS)
			if err != nil {
				reject.Invoke(errorConstructor.New(fmt.Errorf("Error, Target int parse fail:%w", err).Error()))

			}

			inputValues, err := jsonToArrayInt(inputsS)
			if err != nil {
				reject.Invoke(errorConstructor.New(fmt.Errorf("Error, input values int parse fail:%w", err).Error()))
			}
			// fmt.Println("We have the inputs:", inputValues)
			ra := runCountdown(target, inputValues)

			resolve.Invoke(map[string]interface{}{
				"error":     err,
				"countdown": string(ra),
			})

		}(targetS, inputS)

		// The handler of a Promise doesn't return any value
		return nil
	})

	// Create and return the Promise object
	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func boggleRun(this js.Value, i []js.Value) interface{} {
	res, err := boggleRunner(i[0].String())
	if err != nil {
		return map[string]interface{}{
			"error": err,
		}
	}
	return map[string]interface{}{
		"error":  err,
		"boggle": string(res),
	}
}
func boggleRunner(jsonIn string) (string, error) {
	ra, err := toBoggleArray(jsonIn)
	if err != nil {
		return "", err
	}
	data, err := wordlist.Asset("data/wordlist.txt")
	if err != nil {
		return "", err
	}

	dic := boggle.NewDictMap([]string{})
	dic.PopulateFromBa(data)

	sortedResult := boggle.NewPuzzleSolve(ra, dic)
	log.Println(sortedResult)
	tmp, err := json.Marshal(sortedResult)
	return string(tmp), err

}
func registerCallbacks() {
	js.Global().Set("anagram", js.FuncOf(anagram))
	//js.Global().Set("boggleRun", js.FuncOf(boggleRun))
	//js.Global().Set("countdown", js.FuncOf(countdown))
	js.Global().Set("countdownPromise", js.FuncOf(countdownPromise))
	js.Global().Set("sudoku", js.FuncOf(sudoku))
	js.Global().Set("bogglePromise", js.FuncOf(bogglePromise))
}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
