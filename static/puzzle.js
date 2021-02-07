const selectAnagramMenu = document.getElementById("select_anagram");
const selectCountdownMenu = document.getElementById("select_countdown");
const selectSudokuMenu = document.getElementById("select_sudoku");

const roomResultList = document.getElementById("results");

const anagramRow = document.getElementById("anagram-input");
const anagramBtn = document.getElementById("anagram-btn");
const anagramInput = document.getElementById("anagram-get-it-here");

const resultRow = document.getElementById("result-row");
loadEventListeners();

function loadEventListeners(){
    selectAnagramMenu.addEventListener('click', selectAnagram);
    selectCountdownMenu.addEventListener('click', selectCountdown);
    selectSudokuMenu.addEventListener('click', selectSudoku);

    anagramBtn.addEventListener('click', calculateAnagram);
}
function hideAllViews(){
    [resultRow, anagramRow].forEach(x => x.style="display:none");
}
function selectAnagram (e) {
    hideAllViews();
    console.log("Select the anagram");
    anagramRow.style = ""; // Display the page

    e.preventDefault();
}
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

function selectCountdown (e) {
    hideAllViews()
    console.log("Select the countdown");
    e.preventDefault();
}
function selectSudoku (e) {
    hideAllViews();
    console.log("Select the sudoku");
    e.preventDefault();
}

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
