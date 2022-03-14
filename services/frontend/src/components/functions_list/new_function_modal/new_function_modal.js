import React from "react";
import {Alert, Form, Input, Modal} from "antd";
import {internalApiURL} from "../../../config";
import axios from "axios";
import {FunctionActions} from "../../../actions";
import {connect} from "react-redux";
import Editor from "@monaco-editor/react";
import JSZip from "jszip"


const initialCodeText = `from typing import Union


# please, do not edit the function name.
def calling_function(input_data: bytes) -> Union[bytes, str]:
    """
    input_data : bytes object with input data.
    Return : bytes or str
        result of executing, will be sent to the client.
    """
    return input_data
`


function generateZipArchive(functionName, functionInCode, code) {
    let zip = new JSZip()

    let requirements = ""
    let env = "python3.6"
    let callingFunction = functionInCode

    const lineSplitted = code.split('\n')
    let finalCode = ""
    lineSplitted.forEach(element => {
        if (element.startsWith("%%")) {
            let wordSplitted = element.split(' ')
            if (wordSplitted.length === 2) {
                if (wordSplitted[0] === "%%req") {
                    requirements += `${wordSplitted[1]}\n`
                } else if (wordSplitted[0] === "%%env") {
                    env = wordSplitted[1]
                } else if (wordSplitted[0] === "%%calling_func") {
                    callingFunction = wordSplitted[1]
                }
            }
        } else {
            finalCode += `${element}\n`
        }
    });

    zip.file("main.py", finalCode)

    let yamlContents =
        `env: ${env}
name: ${functionName}
python:
  call: main:${callingFunction}
`
    if (requirements !== "") {
        yamlContents += "  requirements: requirements.txt\n"
        zip.file("requirements.txt", requirements)
    }
    zip.file("deepmux.yaml", yamlContents)

    console.log(finalCode)
    console.log(requirements)
    console.log(yamlContents)

    return zip
}


class NewFunctionModalUnwrapped extends React.Component {

    state = {
        confirmLoading: false,
        generalFormError: "",
        codeValue: initialCodeText,
    };

    formRef = React.createRef();

    handleOk = async () => {
        this.setState({
            confirmLoading: true,
        });
        const apiFunctionName = this.formRef.current.getFieldValue('apiFunctionName');
        const codeFunctionName = this.formRef.current.getFieldValue('codeFunctionName');
        const regex = /[a-zA-Z0-9_]+/

        if (apiFunctionName === "" || apiFunctionName === undefined) {
            this.setState({
                generalFormError: "Api function name should not be empty",
                confirmLoading: false,
            })
            return
        } else if (!regex.test(apiFunctionName)) {
            this.setState({
                generalFormError: "Api function name should have only a-z, A-Z, 0-9 and _",
                confirmLoading: false,
            })
            return
        } else if (codeFunctionName === "" || codeFunctionName === undefined) {
            this.setState({
                generalFormError: "Code's function name should not be empty",
                confirmLoading: false,
            })
            return
        } else if (!regex.test(codeFunctionName)) {
            this.setState({
                generalFormError: "Code's function name should have only a-z, A-Z, 0-9 and _",
                confirmLoading: false,
            })
            return
        }


        const endpoint = `${internalApiURL}/v1/function/${apiFunctionName}`;

        const zip = generateZipArchive(apiFunctionName, "calling_function", this.state.codeValue)
        const res = await zip.generateAsync({type: "blob"})
        let bodyFormData = new FormData()
        bodyFormData.append('repo', res)

        try {
            await axios.put(endpoint)
            console.log("Sent put request")
            await axios.post(endpoint, bodyFormData, {
                headers: {"Content-Type": "multipart/form-data"}
            })
            console.log("Sent post request")

            this.props.get_functions_list();
            this.props.hideModal();

            this.setState({
                generalFormError: "",
                confirmLoading: false,
                codeValue: initialCodeText
            });
            this.formRef.current.resetFields();
        } catch (error) {
            console.log(error)
            console.log(error.constructor.name)
            try {
                if (error.response.status === 400) {
                    this.setState({
                        generalFormError: error.response.data.message,
                        confirmLoading: false,
                    })
                }
                if (error.response.status === 409) {
                    this.setState({
                        generalFormError: "Function with this name already exists",
                        confirmLoading: false,
                    })
                } else {
                    this.setState({
                        generalFormError: error.toString(),
                        confirmLoading: false,
                    })
                }
            } catch (_) {
                this.setState({
                    generalFormError: "Unknown error",
                    confirmLoading: false
                })
            }
        }
    };

    handleCancel = () => {
        this.props.hideModal();
    };

    onCodeChange = (newValue, _) => {
        this.setState({
            codeValue: newValue,
        });
    };

    render() {
        const {confirmLoading} = this.state;

        return (
            <div>
                <Modal title="Add New Function" visible={this.props.visible} onOk={this.handleOk}
                       confirmLoading={confirmLoading} onCancel={this.handleCancel} width="80%"
                       style={{minHeight: "80%"}} centered={true}>

                    {this.state.generalFormError !== "" &&
                    <Alert message={this.state.generalFormError} type="error" showIcon style={{marginBottom: "10px"}}/>}
                    <Form name="basic" ref={this.formRef}
                          initialValues={{apiFunctionName: "newFunction", codeFunctionName: "calling_function"}}>
                        <Form.Item
                                    label="Function Name"
                                    name="apiFunctionName"
                                    placeholder="newFunction"
                                    value="newFunction"
                                    rules={[
                                        {
                                            required: true,
                                            message: 'Function should have not empty name',
                                        },
                                        {
                                            type: "regexp",
                                            pattern: /^[a-zA-Z0-9_]+$/,
                                            message: "Api function name should have only a-z, A-Z, 0-9 and _"
                                        }
                                    ]}
                                    width="100%"
                                >
                                    <Input value="newFunction"/>
                                </Form.Item>

                        <Editor width="100%" height="50vh" defaultLanguage="python"
                                options={{minimap: {enabled: false}, contextmenu: false}}
                                defaultValue={initialCodeText} onChange={this.onCodeChange}/>
                    </Form>
                </Modal>
            </div>
        );
    }
}


const actionCreators = {
    get_functions_list: FunctionActions.get_functions_list,
};

const NewFunctionModal = connect((_) => {}, actionCreators,)(NewFunctionModalUnwrapped);

export default NewFunctionModal;
