function numToStringWithPad(value, size) {
    var s = String(value)
    while (s.length < (size || 2)) {
        s = "0" + s
    }
    return s
}


function getTimestampString(seconds, nanoseconds) {
    const date = new Date(seconds * 1000)
    return numToStringWithPad(date.getFullYear(), 4) + "-" + numToStringWithPad(date.getMonth(), 2) + "-" +
        numToStringWithPad(date.getDay(), 2) + " " + numToStringWithPad(date.getHours(), 2) + ":" +
        numToStringWithPad(date.getMinutes(), 2) + ":" + numToStringWithPad(date.getSeconds(), 2)
}


function saveData(data, type, name) {
    var a = document.createElement("a");
    document.body.appendChild(a);
    a.style = "display: none";
    var blob = new Blob([data], {type: type})
    var url = window.URL.createObjectURL(blob);
    a.href = url;
    a.download = name;
    a.click();
    window.URL.revokeObjectURL(url);
}


export {numToStringWithPad, getTimestampString, saveData}