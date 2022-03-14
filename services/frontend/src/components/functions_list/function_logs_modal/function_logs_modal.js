import React from "react";
import {Alert, Modal} from "antd";
import {LogsViewer} from "../logs"
import {loadState} from "../../../persist_store";
import {baseApiURL} from "../../../config";


class FunctionLogsModal extends React.Component {
    state = {
        generalFormError: "",
        data: []
    };

    gettingLogsStarted = false
    source = null

    componentWillUnmount() {
        this.stopGettingLogs()
    }

    stopGettingLogs() {
        try {
            this.source.close()
        } catch(_) {}
        this.source = null
        this.gettingLogsStarted = false
        if (this.state.data.length !== 0) {
            this.setState({data: []})
        }
    }

    startGettingLogs = async () => {
        let isAlreadyGetting = this.gettingLogsStarted
        this.gettingLogsStarted = true
        if (isAlreadyGetting === true) {
            return
        }

        this.setState({data: []})

        let urlPrefix = 'ws'
        if (window.location.protocol === 'https:') {
            urlPrefix = 'wss'
        }
        const endpoint = `${urlPrefix}:${baseApiURL}/v1-a/function/${this.props.func.name}/stream_logs`
        const token = loadState().authentication.token
        this.source = new WebSocket(endpoint)

        this.source.onopen = event => {
            this.source.send(token)
        }

        this.source.onmessage = event => {
            if (event.data !== undefined) {
                let json = JSON.parse(event.data)
                if (json.message !== "keepalive") {
                    let data = this.state.data
                    data.push(json)
                    this.setState({data: data})
                }
                this.source.send("keepalive")
            }
        }

        this.source.onerror = event => {
            console.log(event)
        }

        this.source.onping = event => {
            console.log("sdf")
        }
    };

    handleCancel = () => {
        this.setState({generalFormError: ""})
        this.props.hideModal();
    };

    onCodeChange = (newValue, _) => {
        this.setState({
            codeValue: newValue,
        });
    };

    render() {
        if (this.props.visible) {
            this.startGettingLogs()
        } else {
            this.stopGettingLogs()
        }

        return (
            <div>
                <Modal title={`Logs of ${this.props.func.name}`} visible={this.props.visible}
                       onCancel={this.handleCancel} width="90%" style={{minHeight: "80%"}} centered={true}
                       okButtonProps={{style: {display: 'none'}}} cancelText="Close">
                    {this.state.generalFormError !== "" &&
                    <Alert message={this.state.generalFormError} type="error" showIcon style={{marginBottom: "10px"}}/>}
                    <LogsViewer data={this.state.data}/>
                </Modal>
            </div>
        );
    }
}


export default FunctionLogsModal;
