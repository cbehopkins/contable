const form = document.querySelector('#room-form');
const roomResultList = document.querySelector('.collection');
const displayBtn = document.querySelector('.display-results');
const nextPageBtn = document.querySelector('.next_page');
//# says ID of
const filter = document.querySelector('#InputText');
const roomCountInput = document.querySelector('#RoomCount');

loadEventListeners();

function loadEventListeners(){
    displayBtn.addEventListener('click', displayResults);
    nextPageBtn.addEventListener('click', nextPage);
}

const srcJsonHard = `[
    [{"People":[{"Name":"bob","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"}]},{"Name":"this","Connections":[{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]}]},{"People":[{"Name":"other","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"bob"}]},{"Name":"that","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]}]}],[{"People":[{"Name":"bob","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"}]},{"Name":"that","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]}]},{"People":[{"Name":"this","Connections":[{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]},{"Name":"other","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"bob"}]}]}],[{"People":[{"Name":"bob","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"}]},{"Name":"other","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"bob"}]}]},{"People":[{"Name":"this","Connections":[{"Count":1,"PerLink":"that"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]},{"Name":"that","Connections":[{"Count":1,"PerLink":"this"},{"Count":1,"PerLink":"other"},{"Count":1,"PerLink":"bob"}]}]}]
]
`
let disCount = 0;
let numMeetings = 0;
function displayResults(e){
    calculate('RoomCount', 'InputText', 'result');
    disCount = 0;
    while(roomResultList.firstChild) {
        roomResultList.removeChild(roomResultList.firstChild);
    }
    let roomCount = roomCountInput.value;
    for (i=0; i< roomCount; i++){
        addRoomDisplay(i);
    }
    populateRoomDisplay();
}
function nextPage(e){
    console.log(disCount);
    disCount++;
    if (disCount >= numMeetings) {
        disCount = 0;
    }
    console.log(disCount);
    populateRoomDisplay();
}
function addRoomDisplay(item) {
    const rw = document.createElement('tr');
    rw.className = 'collection-item';

    let ta = document.createElement("textarea");
    ta.value = "Hi\nThere";
    rw.appendChild(ta);

    roomResultList.appendChild(rw);
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
    document.getElementById("room_display").textContent = `Meeting Number: ${disCount + 1}`

    let srcJson = document.querySelector("#result").value;
    if (false) {
        srcJson = srcJsonHard;
    }
    //console.log(`Got ${srcJson}`)
    if (srcJson == "") {
        return;
    }
    let resultObj = JSON.parse(srcJson);
    numMeetings = resultObj.length;
    let selectedItem = resultObj[disCount];
    let selectedRoom = roomResultList.firstElementChild;
    for (let i=0; i<selectedItem.length; i++){
        peopleinRoom = resultObj[disCount][i]["People"];        
        selectedRoom.firstElementChild.value = renderPeople(peopleinRoom);
        selectedRoom = selectedRoom.nextElementSibling
    }
}
