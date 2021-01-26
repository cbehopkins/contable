const form = document.querySelector('#room-form');
const roomResultList = document.querySelector('.collection');
const displayBtn = document.querySelector('.display-results');
const nextPageBtn = document.querySelector('.next_page');
//# says ID of
const filter = document.querySelector('#InputText');
const roomCountInput = document.querySelector('#RoomCount');
let tx = document.getElementsByTagName('textarea');

loadEventListeners();

function loadEventListeners(){
    displayBtn.addEventListener('click', displayResults);
    nextPageBtn.addEventListener('click', nextPage);
}

// Here there are 4 people in 2 rooms
// which takes 3 meetings for them all to see each other
const srcJsonHard = `[
    [{"People":[{"Name":"bob","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"}]},
                {"Name":"this","Connections":[{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]}]},
     {"People":[{"Name":"other","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"bob"}]},
                {"Name":"that","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]}]}],
    [{"People":[{"Name":"bob","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"}]},
                {"Name":"that","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]}]},
     {"People":[{"Name":"this","Connections":[{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]},
                {"Name":"other","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"bob"}]}]}],
    [{"People":[{"Name":"bob","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"}]},
                {"Name":"other","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"bob"}]}]},
     {"People":[{"Name":"this","Connections":[{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]},
                {"Name":"that","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]}]}]
]
`
let disCount = 0;
let numMeetings = 0;
let useWasm = true;

function displayResults(e){
    if (useWasm) {
        calculate('RoomCount', 'InputText', 'result');
    }
    disCount = 0;
    populateRoomDisplay();
}
function nextPage(e){
    console.log(disCount);
    disCount++;
    if (disCount >= (numMeetings-1)) {
        disCount = 0;
    }
    populateRoomDisplay();
}

function renderPeople(people) {
    roomText = ""
    let spacer = ""
    for (let i=0; i<people.length; i++){
        roomText += spacer + people[i].Name;
        spacer = "\n";
    }
    return roomText;
}

function populateRoomDisplay(){
    

    let srcJson = document.querySelector("#result").value;
    if (!useWasm) {
        srcJson = srcJsonHard;
    }
    //console.log(`Got ${srcJson}`)
    if (srcJson == "") {
        return;
    }
    // srcJson = srcJsonHard;
    let resultObj = JSON.parse(srcJson);
    numMeetings = resultObj.length;
    document.getElementById("room_display").textContent = `Meeting Number: ${disCount + 1} of ${numMeetings}`
    let resTab = generateRoomMappingFromInput(resultObj)
    createDisplayTable(resTab);
    populateResultsTable(resTab);

}

function roomMapping (first, roomList) {
    // first is of the format:  ["other", "bob"]
    //  room list is somethig like:
    // (2) [Array(2), Array(2)]
    // 0: (2) ["bob", "this"]
    // 1: (2) ["other", "that"]
    // We are being asked to find which room "other" is in
    // In which case we should return 1
    // so we should return [1, 0]
    return first.map(x => roomList.findIndex(y => y.includes(x)));
}
function determineMovement(first, last){
    // first and last come in looking like:  ["bob↵this", "other↵that"]
    first_column_split = first.map(x => x.split("\n"));
    last_column_split = last.map(x => x.split("\n"));
    // For each entry in the first list which is a list of rooms
    // Each room being a list of people
    // i.e. [["bob", "this"], ["other", "that"]]
    // The last list is the room they will be in next
    // So construct a list of the room they will be in next:
    // e.g. [[0, 1], [1,0]]
    let mapping = first_column_split.map(x => roomMapping(x, last_column_split));
    let resultA = Array()
    for (let i=0; i< mapping.length; i++){
        let text = ""
        let split = ""
        for (let j=0; j< mapping[i].length; j++){
            if (mapping[i][j] == i) {
                // Don't print anything if we stay in the same place
                continue;
            }
            text += split + `${first_column_split[i][j]} => ${mapping[i][j] + 1}`;
            split = "\n";
        }
        resultA.push(text);
    }
    return resultA;
}

function generateRoomMappingFromInput(resultObj){
    let first_column = resultObj[disCount].map(x => renderPeople(x["People"]))
    let last_column = resultObj[disCount+1].map(x => renderPeople(x["People"]))
    let move_column = determineMovement(first_column, last_column);
    let resultArray = Array()
    for (let i = 0; i< resultObj[0].length; i++) {
        resultArray.push ([first_column[i], move_column[i], last_column[i]])
    }
    return resultArray
}
function autoResize(th){
    th.style.height = 'auto';
    th.style.height = (th.scrollHeight) + 'px';
}
function populateResultsTable(inputArray) {
    let selectedRow = roomResultList.firstElementChild;
    for (let i=0; i<inputArray.length; i++) {
        populateResultsRow(selectedRow, inputArray[i]);
        selectedRow = selectedRow.nextElementSibling;
    }
}
function populateResultsRow(selectedRow, inputArray) {
    selectedElement = selectedRow.firstElementChild
    for (let i=0; i<inputArray.length; i++) {
        selectedElement.firstElementChild.value = inputArray[i];
        autoResize(selectedElement.firstElementChild);
        selectedElement = selectedElement.nextElementSibling
    }
}
function createDisplayTable(resTab){
    while(roomResultList.firstChild) {
        roomResultList.removeChild(roomResultList.firstChild);
    }
    for (i=0; i< resTab.length; i++){
        addRoomDisplay(resTab[i]);
    }
}
function addRoomDisplay(item) {
    let rw = createColumns(item.length);
    roomResultList.appendChild(rw);
}
function createColumns(cnt){
    const rw = document.createElement('tr');
    rw.className = 'collection-item';
    for (let i=0; i<cnt; i++) {
        const td = document.createElement('td');
        let ta = document.createElement("textarea");
        ta.value = "";
        td.appendChild(ta);
        rw.appendChild(td);
    }
    return rw
}

