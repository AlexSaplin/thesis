import React from "react";
import {Alert, Modal, Form, Input} from "antd";
import Editor from "@monaco-editor/react";
import {saveData} from "../../../../utils/functions";


class ResultFunctionModal extends React.Component {
    formRef = React.createRef();

    saveArray = () => {
        const extension = this.formRef.current !== null ? this.formRef.current.getFieldValue('extension') : ""
        const name = `${this.props.functionName}_output${extension}`
        saveData(this.props.data, "octet/stream", name)
    }

    render() {
        let innerAreaText = <Editor width="100%" height="50vh" options={{
            minimap: {enabled: false},
            contextmenu: false,
            wordBasedSuggestions: false,
            autoClosingBrackets: false,
            autoIndent: "none",
            readOnly: true
        }}
        defaultValue={this.props.errorMessage !== "" ? this.props.data : Buffer.from(this.props.data, 'binary').toString()}/>
        let width = "80%"

        if (this.props.data.byteLength >= 16 * 1024) {
            const rules = [{pattern: /^(\.[^.]+)+$/, message: "Please, type correct file extension"}]
            innerAreaText = 
            <Form name="basic" ref={this.formRef} initialValues={{extension: ".txt"}}>
                For files with size larger than 16kb there is no text editor.
                Please, type below the extention to save data.
                <Form.Item name="extension" rules={rules} width="100%"><Input/></Form.Item>
            </Form>
            width = "400px"
        }

        return (
            <div>
                <Modal title={`Result of running function ${this.props.functionName}`} visible={this.props.visible} 
                onOk={this.saveArray} confirmLoading={false} onCancel={this.props.handleCancel} width={width}
                centered={true} okText="Download" cancelText="Close" closeIcon={<div></div>}>
                    {this.props.errorMessage !== "" &&
                    <Alert message={this.props.errorMessage} type="error" showIcon style={{marginBottom: "10px"}}/>}
                    {innerAreaText}
                </Modal>
            </div>
        )
    }
}


export default ResultFunctionModal
