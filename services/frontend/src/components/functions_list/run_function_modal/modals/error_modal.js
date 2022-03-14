import React from "react"
import {Modal} from "antd"
import CloseCircleOutlined from '@ant-design/icons/CloseCircleOutlined'
import {LogsViewer, saveLogs} from "../../logs";


class ErrorFunctionModal extends React.Component {
    render() {
        if (this.props.logs.length === 0) {
            return (
                <Modal title="Error" visible={this.props.visible} confirmLoading={false} onOk={this.props.handleClose} cancelText="Close" 
                width="300px" centered={true} cancelButtonProps={{style: {display: 'none'}}} closeIcon={<div></div>} showIcon={<CloseCircleOutlined />}>
                    {this.props.error}
                </Modal>
            )
        } else {
            return (
                <Modal title="Error" visible={this.props.visible} confirmLoading={false} onCancel={this.props.handleClose} cancelText="Close"
                       onOk={() => saveLogs(this.props.functionName, this.props.logs)} width="80%" centered={true} closeIcon={<div/>}
                       showIcon={<CloseCircleOutlined />} okText="Download logs">
                    {this.props.error}
                    <LogsViewer data={this.props.logs}/>
                </Modal>
            )
        }
        
    }
}


export default ErrorFunctionModal;
