import React from "react"
import {getTimestampString, saveData} from "../../utils/functions"
import AlwaysScrollToBottom from "../../utils/components/AlwaysScrollToBottom"


function saveLogs(functionName, logs) {
    const name = `${functionName}_logs.txt`
    let result = ""
    logs.forEach(item => {
        result += `[${getTimestampString(item.time.seconds, item.time.nanos)}] ${item.message}\n`
    })
    saveData(result, 'text/plain', name)
}


function LogLine(props) {
    return (<div style={{color: "white", marginTop: "5px", fontSize: "10pt", fontFamily: "monaco"}}>
        [{getTimestampString(props.item.time.seconds, props.item.time.nanos)}] {props.item.message}
    </div>)
}


function LogsViewer(props) {
    return (<div style={{width: "100%", height: "70vh", overflow: "scroll", backgroundColor: "black", padding: "10px"}}>
        {props.data.map((item, i) => <LogLine key={i} item={item}/>)}
        <AlwaysScrollToBottom />
    </div>)
}


export {saveLogs, LogsViewer}
