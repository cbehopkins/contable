package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/cbehopkins/ana"
	cntSlv "github.com/cbehopkins/countdown/cnt_slv"
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
	foundValues := cntSlv.NewNumMap()
	//found_values.SelfTest = true
	foundValues.UseMult = true
	foundValues.PermuteMode = cntSlv.FastMap
	foundValues.SeekShort = false // TBD make this controllable

	fmt.Println("Starting permute")
	returnProofs := foundValues.CountHelper(target, sources)
	for range returnProofs {
		//fmt.Println("Proof Received", v)
		// Here we get the proofs as they are found
	}
	//fmt.Println("Permute Complete", proof_list)
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
	//jsonStr := `[["","","","2","6","","7","",""],["6","8","","","7","","","9",""],["1","9","","","","4","5","",""],["8","2","","1","","","","4",""],["","","4","6","","2","9","",""],["","5","","","","3","","2","8"],["","","9","3","","","","7","4"],["","4","","","5","","","3","6"],["7","","3","","1","8","","",""]]`
	ra, err := jsonTo2dArrayInt(i[0].String())
	// fmt.Println("We have our array now:", ra)
	returnArray := runSudoku(ra)
	tmp, err := json.Marshal(returnArray)
	return map[string]interface{}{
		"error":  err,
		"sudoku": string(tmp),
	}
}

func runSudoku(input [][]int) (output [][]int) {
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
	result := testPuzzle.SelfCheck()
	if result != nil {
		// TBD add error field we can report this to
		return
	}
	testPuzzle.SolveAll()

	result = testPuzzle.SelfCheck()
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
