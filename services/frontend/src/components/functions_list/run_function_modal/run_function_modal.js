import React from "react";
import {internalApiURL} from "../../../config";
import axios from "axios";
import {FunctionActions} from "../../../actions";
import {connect} from "react-redux";
import InputFunctionModal from "./modals/input_modal"
import WaitFunctionModal from "./modals/wait_modal"
import ResultFunctionModal from "./modals/result_modal"
import ErrorFunctionModal from "./modals/error_modal"


const initialFormData = "here type input"
const MAX_DATA_SIZE = 64 * 1024 * 1024


const ModalState = {input: 0, wait: 1, result: 2, error: 3}
Object.freeze(ModalState)


class RunFunctionModalUnwrapped extends React.Component {
    state = {
        confirmLoading: false,
        generalFormError: "",
        data: initialFormData,
        currentStep: ModalState.input,
        cancelToken: null,
        fileData: null,
        uploadProgress: 0,
        downloadProgress: 0,
        isUploadProgressChecking: true,
        isDownloadProgressChecking: true,
        isCancelled: false,
        logs: []
    };

    handleRun = async () => {
        await this.startRequest(this.state.data, this.state.data.length)
    }

    restoreState = (data, generalFormError, currentStep, advanced={}) => {
        this.setState({
            confirmLoading: false,
            generalFormError: generalFormError,
            data: data,
            currentStep: currentStep,
            cancelToken: null,
            fileData: null,
            uploadProgress: 0,
            downloadProgress: 0,
            isUploadProgressChecking: true,
            isDownloadProgressChecking: true,
            ...advanced
        });
    }

    logsReceiver = (data) => {
        this.setState({logs: data})
    }

    changeProgresses = (event, valueField, checkingField) => {
        if (event.total === 0) {
            if (this.state[checkingField] === true) {
                this.setState({
                    [checkingField]: false
                });
            }
        } else {
            const percent = Math.floor((event.loaded / event.total) * 100)
            this.setState({
                [valueField]: percent
            });
        }
    }

    startRequest = async (data, dataLength) => {
        if (dataLength > MAX_DATA_SIZE) {
            this.restoreState("Max size of data is 64mb", "Max size of data is 64mb", ModalState.error)
            return
        }

        const endpoint = `${internalApiURL}/v1/function/${this.props.func.name}/run`;
        const cancelToken = axios.CancelToken.source()
        try {
            this.setState({
                currentStep: ModalState.wait,
                cancelToken: cancelToken,
                waitStage: "upload"
            });
            const res = await axios.post(endpoint, data, 
                {
                    cancelToken: cancelToken.token, 
                    transformResponse: (r) => r,
                    onUploadProgress: (event) => {this.changeProgresses(event, "uploadProgress", "isUploadProgressChecking")},
                    onDownloadProgress: (event) => {this.changeProgresses(event, "downloadProgress", "isDownloadProgressChecking")},
                    responseType: 'arraybuffer',
                    timeout: 60000
                })
            this.restoreState(res.data, "", ModalState.result)
        } catch (error) {
            console.log(error.code)
            if (axios.isCancel(error)) {
                return
            }
            try {
                if (error.code === 'ECONNABORTED') {
                    this.restoreState(
                        "Timeout of 60s reached. There can be some server errors. Please, try to reaload the page. For big files, please, use CLI.",
                        "Timeout of 60s reached. There can be some server errors. Please, try to reaload the page. For big files, please, use CLI.",
                        ModalState.error
                    )
                } else if (error.response.status === 400) {
                    this.restoreState(
                        JSON.parse(String.fromCharCode.apply(null, new Uint8Array(error.response.data))).message,
                        "Error occured while running container:",
                        ModalState.result
                    )
                } else if (error.response.status === 500) {
                    this.restoreState(
                        "Can't run because of internal server error",
                        "Can't run because of internal server error",
                        ModalState.error
                    )
                } else if (error.response.status === 408) {
                    console.log("pyisa")
                } else {
                    this.restoreState(error.toString(), error.toString(), ModalState.error)
                }
            } catch (_) {
                this.restoreState("Unknown error", "Unknown error", ModalState.error)
            }
        }
    };

    handleCancel = () => {
        this.props.hideModal();
        try {
            this.state.cancelToken.cancel()
        } catch (_) {}

        this.restoreState(initialFormData, "", ModalState.input, {logs: []})
    };

    onDataChange = (newValue, _) => {
        this.setState({
            data: newValue,
        });
    };

    addFile = async (options) => {
        const { file } = options;
        const reader = new FileReader();
        reader.readAsArrayBuffer(file)
        reader.onload = async (_) => {
            await this.startRequest(reader.result, reader.result.byteLength)
        }
      };

    render() {
        if (this.state.currentStep === ModalState.input) {
            return (<div><InputFunctionModal functionName={this.props.func.name} visible={this.props.visible} 
                errorMessage={this.state.generalFormError} data={this.state.data} onDataChange={this.onDataChange} 
                handleCancel={this.handleCancel} handleRun={this.handleRun} handleFile={this.addFile}/>
                </div>)
        } else if (this.state.currentStep === ModalState.wait) {
            return (
                <div>
                    <WaitFunctionModal functionName={this.props.func.name} visible={this.props.visible} handleCancel={this.handleCancel}
                    uploadProgress={this.state.uploadProgress} isUploadProgressChecking={this.state.isUploadProgressChecking}
                    downloadProgress={this.state.downloadProgress} isDownloadProgressChecking={this.state.isDownloadProgressChecking}
                    logsReceiver={this.logsReceiver}/>
                    </div>
                    )
        } else if (this.state.currentStep === ModalState.result) {
            return (
            <div>
                <ResultFunctionModal functionName={this.props.func.name} visible={this.props.visible} 
                errorMessage={this.state.generalFormError} data={this.state.data} handleCancel={this.handleCancel}/>
            </div>)
        } else {
            return (
                <div>
                    <ErrorFunctionModal functionName={this.props.func.name} visible={this.props.visible}
                    handleClose={this.handleCancel} error={this.state.generalFormError} logs={this.state.logs}/>
                </div>
            );
        }
    }
}


const actionCreators = {
    get_functions_list: FunctionActions.get_functions_list,
};

const RunFunctionModal = connect((_) => {}, actionCreators,)(RunFunctionModalUnwrapped);

export default RunFunctionModal;
