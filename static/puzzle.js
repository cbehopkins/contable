const selectAnagramMenu = document.getElementById("select_anagram");
const selectCountdownMenu = document.getElementById("select_countdown");
const selectSudokuMenu = document.getElementById("select_sudoku");

const roomResultList = document.getElementById("results");

const anagramRow = document.getElementById("anagram-input");
const anagramBtn = document.getElementById("anagram-btn");
const anagramInput = document.getElementById("anagram-get-it-here");

const countdownRow = document.getElementById("countdown-input");
const countdownBtn = document.getElementById("countdown-btn");
const countdownTable = document.getElementById("countdown-get-it-here");
const countdownTarget = document.getElementById('countdown-target-input');

const sudokuRow = document.getElementById("sudoku-input");
const sudokuBtn = document.getElementById("sudoku-btn");
const sudokuTable = document.getElementById("sudoku-get-it-here");

const resultRow = document.getElementById("result-row");
loadEventListeners();

function loadEventListeners(){
    selectAnagramMenu.addEventListener('click', selectAnagram);
    selectCountdownMenu.addEventListener('click', selectCountdown);
    selectSudokuMenu.addEventListener('click', selectSudoku);

    anagramBtn.addEventListener('click', calculateAnagram);
    countdownBtn.addEventListener('click', calculateCountdown);
    sudokuBtn.addEventListener('click', calculateSudoku);

}
function hideAllViews(){
    [resultRow, anagramRow, countdownRow, sudokuRow].forEach(x => x.style="display:none");
}

///////////////////////////////////////
// Select the puzzle in quesiotn and display
///////////////////////////////////////
function selectAnagram (e) {
    hideAllViews();
    console.log("Select the anagram");
    anagramRow.style = ""; // Display the page

    e.preventDefault();
}

function selectCountdown (e) {
    hideAllViews()
    console.log("Select the countdown");
    inputTablePopulate(countdownTable, 6, 1, "countdown", "number")
    countdownRow.style= ""; // Display the page
    e.preventDefault();
}
function selectSudoku (e) {
    hideAllViews();
    console.log("Select the sudoku");
    inputTablePopulate(sudokuTable, 9, 9, "sudoku", "number")
    sudokuRow.style= ""; // Display the page
    e.preventDefault();
}

///////////////////////////////////////
// Run the Calculations
///////////////////////////////////////
function calculateAnagram(e){
    returnStruct = anagram(anagramInput.value);
    console.log(returnStruct)
    if (returnStruct["error"] != "") {
        console.log(`Got an error ${returnStruct["err"]}`)
        return;
    }
    resultArray = rePopulate(JSON.parse(returnStruct["anagrams"]), 3);
    populateTab(roomResultList, resultArray);
    e.preventDefault();
}
function calculateCountdown(e){
    // The first input number is always the target
    // let testInput = ["4", "5", "10", "100", "50", "1"]
    let inputTable = inputTableRetrieve(countdownTable); 

    // We want the first (and only) row of the table
    returnStruct = countdown(countdownTarget.value, JSON.stringify(inputTable[0]));
    console.log(returnStruct)
    if (returnStruct["error"] != "") {
        console.log(`Got an error ${returnStruct["err"]}`)
        return;
    }
    let rs = returnStruct["countdown"]
    console.log(`Result:${rs}`)
    // resultArray = rePopulate(returnStruct["countdown"], 3);
    // populateTab(roomResultList, resultArray);
    e.preventDefault();
}
function calculateSudoku(e){
    inputTable = inputTableRetrieve(sudokuTable);
    returnStruct = sudoku(inputTable);
    console.log(returnStruct)
    if (returnStruct["error"] != "") {
        console.log(`Got an error ${returnStruct["err"]}`)
        return;
    }
    resultArray = rePopulate(JSON.parse(returnStruct["sudoku"]), 3);
    populateTab(roomResultList, resultArray);
    e.preventDefault();
}
///////////////////////////////////////
// The generate table population functions
// Live below here
///////////////////////////////////////


function clearTable(element) {
    while(element.firstChild) {
        element.removeChild(element.firstChild);
    }
}
function populateTab(element, inputArray) {
    resultRow.style = "display:none"

    clearTable(element)
    for (let i=0; i<inputArray.length; i++) {
        let row = document.createElement("tr")
        populateTabRow(row, inputArray[i]);
        element.appendChild(row)
    }
    resultRow.style = "";
}
function populateTabRow (row, inputArray) {
    for (let i=0; i<inputArray.length; i++) {
        let cell = document.createElement("td")
        cell.innerText  = inputArray[i];
        row.appendChild(cell)
    }
}

function rePopulate(table, length){
    // Take a 1 dimensional input table
    // and return a 2d one with max length
    // per row
    let resultTable = Array();
    for (let i=0; i<table.length; i+=length)
    {
        resultTable.push(table.slice(i, i+length));
    }
    return resultTable
}


function inputTablePopulate(element, x, y, label, type){
    for (let i=0; i<y; i++){
        let row = document.createElement('tr')
        for (let j=0;j<x;j++){
            ip = document.createElement('input')
            ip.id = `${label}__input__${j}__${i}`;
            ip.type = type
            td = document.createElement('td')
            td.appendChild(ip)
            row.appendChild(td)
        }
        element.appendChild(row)
    }
}

function inputTableRetrieve(element){
    let returnTable = Array();
    // let alerted = false;
    for (let row = element.firstChild; row!=null; row=row.nextElementSibling) {
        console.log(`Got a new row ${row}`);
        let rowEntry = Array();
        for (let cel=row.firstChild; cel!=null; cel=cel.nextElementSibling){
            console.log(`got a new column ${cel}`);
            const contents = cel.firstChild.value;
            // if (!alerted && (contents === "")){
            //     alert(`Cell has null contents`);
            //     alerted = true
            // }
            rowEntry.push(contents);
        }
        returnTable.push(rowEntry);
    }
    return returnTable;
}