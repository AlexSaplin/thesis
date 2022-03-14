import React from "react"
import {Modal, Row, Spin, Progress} from "antd"
import {LoadingOutlined} from '@ant-design/icons'
// import {internalApiURL} from "../../../../config"
// import {EventSourcePolyfill} from "event-source-polyfill"
// import {loadState} from "../../../../persist_store"
// import {LogsViewer} from "../../logs";


class WaitFunctionModal extends React.Component {
    state = {
        data: []
    };

    // gettingLogsStarted = false
    // source = null

    componentWillUnmount() {
        this.stopGettingLogs()
    }

    stopGettingLogs() {
        try {
            this.source.close()
        } catch (_) {}
        this.source = null
    }

    // startGettingLogs() {
    //     let isAlreadyGetting = this.gettingLogsStarted
    //     this.gettingLogsStarted = true
    //     if (isAlreadyGetting === true) {
    //         return
    //     }
    //
    //     const endpoint = `${internalApiURL}/v1/function/${this.props.functionName}/stream_logs`
    //     this.setState({data: []});
    //
    //     const token = loadState().authentication.token
    //
    //     this.source = new EventSourcePolyfill(endpoint, {
    //         headers: {'X-Token': token, 'Authorization': 'Token ' + token}
    //     })
    //
    //     this.source.onmessage = event => {
    //         let json = JSON.parse(event.data)
    //         let data = this.state.data
    //         data.push(json)
    //         this.setState({data: data})
    //         this.props.logsReceiver(data)
    //     }
    //
    //     this.source.onerror = event => {
    //         if (event.error !== undefined) {
    //             let data = this.state.data
    //             data.push({
    //                 time: {seconds: Math.floor(Date.now() / 1000), nanos: (Date.now() % 1000) * 10},
    //                 message: "Connection error. Reconnecting..."
    //             })
    //             this.setState({data: data})
    //             this.props.logsReceiver(data)
    //         }
    //     }
    //
    //     this.onopen = event => {
    //         let data = this.state.data
    //         data.push({
    //             time: {seconds: Math.floor(Date.now() / 1000), nanos: (Date.now() % 1000) * 10},
    //             message: "Started getting container logs..."
    //         })
    //         this.setState({data: data})
    //         this.props.logsReceiver(data)
    //     }
    // }

    render() {
        const spin = <Spin indicator={<LoadingOutlined style={{ fontSize: 76 }} spin />} width={80}/>
        let stageText = "Uploading data to server"
        let width = "320px"
        // let viewLoadingBars = true

        let uploadItem = null
        if (this.props.uploadProgress !== 100) {
            if (this.props.isUploadProgressChecking === true) {
                uploadItem = <Progress type="circle" percent={this.props.uploadProgress} width={80}/>
            } else {
                uploadItem = spin
            }
        } else {
            uploadItem = <Progress type="circle" percent={100} width={80}/>
            stageText = "Processing function"
        }

        let containerWorkingItem = null
        // let logsArea = null
        if (this.props.uploadProgress !== 100) {
            containerWorkingItem = <Progress type="circle" percent={0} width={80}/>
        } else if (this.props.downloadProgress === 0 && this.props.isDownloadProgressChecking === true) {
            containerWorkingItem = spin
            // this.startGettingLogs()
            // logsArea = <LogsViewer data={this.state.data}/>
            // width = "80%"
            // viewLoadingBars = false
        } else {
            containerWorkingItem = <Progress type="circle" percent={100} width={80}/>
            stageText = "Downloading data from server"
            // this.stopGettingLogs()
        }

        let downloadItem = null
        if (this.props.downloadProgress !== 100) {
            if (this.props.isDownloadProgressChecking === true) {
                downloadItem = <Progress type="circle" percent={this.props.downloadProgress} width={80}/>
            } else {
                downloadItem = spin
            }
        } else {
            downloadItem = <Progress type="circle" percent={100} width={80}/>
            stageText = "Finished"
        }

        return (
            <Modal title={`Running function ${this.props.functionName}: ${stageText}`} visible={this.props.visible}
                   confirmLoading={true} onCancel={this.props.handleCancel} width={width} centered={true}
                   okButtonProps={{style: {display: 'none'}}} closeIcon={<div/>}>
                {/*{viewLoadingBars === true ?*/}
                {/*    <Row align="middle" justify="center" style={{height: "100%"}}>*/}
                {/*        {uploadItem}*/}
                {/*        <div style={{marginLeft: "10px", marginRight: "10px"}}>*/}
                {/*            {containerWorkingItem}*/}
                {/*        </div>*/}
                {/*        {downloadItem}*/}
                {/*    </Row>*/}
                {/*    : logsArea}*/}
                <Row align="middle" justify="center" style={{height: "100%"}}>
                        {uploadItem}
                        <div style={{marginLeft: "10px", marginRight: "10px"}}>
                            {containerWorkingItem}
                        </div>
                        {downloadItem}
                    </Row>
            </Modal>
        );
    }
}


export default WaitFunctionModal
