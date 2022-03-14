import React from "react"
import {Alert, Modal, Upload, message} from "antd"
import Editor from "@monaco-editor/react"


const { Dragger } = Upload


class InputFunctionModal extends React.Component {
    render() {
        const props = {
            name: 'file',
            multiple: false,
            showUploadList: {
                showDownloadIcon: false,
            },
            customRequest: this.props.handleFile,
            onChange(info) {
              const { status } = info.file;
              if (status !== 'uploading') {
                console.log(info.file, info.fileList);
              }
              if (status === 'done') {
                message.success(`${info.file.name} file uploaded successfully.`);
              } else if (status === 'error') {
                message.error(`${info.file.name} file upload failed.`);
              }
            },
          }

        return (
            <div>
                <Modal title={`Run function ${this.props.functionName}`} visible={this.props.visible}
                       onOk={this.props.handleRun}
                       confirmLoading={false} onCancel={this.props.handleCancel} width="80%"
                       style={{minHeight: "80%"}} centered={true} okText="Run" destroyOnClose={true} closeIcon={<div></div>}>
                           {this.props.errorMessage !== "" && <Alert message={this.props.errorMessage} type="error" showIcon
                           />}

                            <Dragger {...props} style={{marginBottom: "5px"}}>
                                Click or drag file to this area to upload
                              </Dragger>
                                <Editor width="100%" height="50vh"
                                    options={{
                                        minimap: {enabled: false},
                                        contextmenu: false,
                                        wordBasedSuggestions: false,
                                        autoClosingBrackets: false,
                                        autoIndent: "none"
                                    }}
                                    defaultValue={this.props.data} onChange={this.props.onDataChange}/>
                </Modal>
            </div>
        );
    }
}


export default InputFunctionModal
