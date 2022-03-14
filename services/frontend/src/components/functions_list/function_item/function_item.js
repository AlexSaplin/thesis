import React from "react";
import {Button, Card, Col, Popconfirm, Row, Tag, Tooltip, Typography} from "antd";
import {CheckCircleOutlined, CloseCircleOutlined, SyncOutlined} from '@ant-design/icons';
import RunFunctionModal from "../run_function_modal/run_function_modal";
import FunctionLogsModal from "../function_logs_modal";
import ViewFunctionMetricsModal from "../metrics_modal";


const {Title} = Typography;


const init_tag = (
    <Tag color="processing"
         style={{fontSize: '15px', lineHeight: '26px', marginTop: '5px', marginBottom: '5px'}} alt="Uplo">
        INITIALIZED
    </Tag>
);

const processing_tag = (
    <Tag icon={<SyncOutlined spin style={{fontSize: '18px'}}/>} color="processing"
         style={{fontSize: '15px', lineHeight: '26px', marginTop: '5px', marginBottom: '5px'}}>
        PROCESSING
    </Tag>
);

const ready_tag = (
    <Tag icon={<CheckCircleOutlined style={{fontSize: '18px'}}/>} color="success"
         style={{fontSize: '15px', lineHeight: '26px', marginTop: '5px', marginBottom: '5px'}}>
        READY
    </Tag>
);

const invalid_tag = (
    <Tag icon={<CloseCircleOutlined style={{fontSize: '18px'}}/>} color="error"
         style={{fontSize: '15px', lineHeight: '26px', marginTop: '5px', marginBottom: '5px'}}>
        ERROR
    </Tag>
)


class FunctionItem extends React.Component {

    state = {
        functionRunModalVisible: false,
        functionLogsModalVisible: false,
        functionMetricsModalVisible: false,
        details: null,
    }

    showFunctionRunModal = () => {
        this.setState({
            functionRunModalVisible: true,
        });
    };

    hideFunctionRunModal = () => {
        this.setState({
            functionRunModalVisible: false,
        })
    };

    showFunctionLogsModal = () => {
        this.setState({
            functionLogsModalVisible: true,
        });
    };

    hideFunctionLogsModal = () => {
        this.setState({
            functionLogsModalVisible: false,
        })
    };

    showFunctionMetricsModal = () => {
        this.setState({
            functionMetricsModalVisible: true,
        });
    };

    hideFunctionMetricsModal = () => {
        this.setState({
            functionMetricsModalVisible: false,
        })
    };


    render() {
        let {name, state, error} = this.props.func;

        return (
            <Card key={name} className="functionslist_model_json" bodyStyle={{padding: "12px 16px 12px 20px"}}>
                <Row align="middle" justify="space-between">
                    <Col>
                        <Title level={4} style={{
                            margin: 0,
                            height: '100%',
                            float: 'left',
                            marginRight: '10px',
                            marginTop: '5px',
                            marginBottom: '5px'
                        }}>
                            Function: {name}
                        </Title>
                        {state === "INIT" && init_tag}
                        {state === "PROCESSING" && processing_tag}
                        {state === "READY" && ready_tag}
                        {state === "INVALID" && <Tooltip title={error}>{invalid_tag}</Tooltip>}
                    </Col>
                    <div>
                        {state === "READY" && <>
                            <Button type="primary" style={{
                                marginTop: '5px',
                                marginBottom: '5px',
                                marginRight: '10px',
                                backgroundColor: '#4092f7',
                                borderColor: '#4092f7'
                            }} float="right" danger onClick={this.showFunctionRunModal}>
                                Run
                            </Button>
                            <Button style={{
                                marginTop: '5px',
                                marginBottom: '5px',
                                marginRight: '10px',
                                borderColor: '#4092f7',
                                color: '#4092f7'
                            }} float="right" danger onClick={this.showFunctionMetricsModal}>
                                Metrics
                            </Button>
                            <Button style={{
                                marginTop: '5px',
                                marginBottom: '5px',
                                marginRight: '10px',
                                borderColor: '#4092f7',
                                color: '#4092f7'
                            }} float="right" danger onClick={this.showFunctionLogsModal}>
                                Logs
                            </Button>
                        </>}

                        <Popconfirm
                            title="Are you sure you want delete this function?"
                            onConfirm={() => this.props.remove_function(name)}
                            onCancel={(e) => {
                                e.stopPropagation()
                            }}
                            cancelText="No"
                            okText="Yes"
                            placement="bottomRight"
                        >
                            <Button type="primary" style={{marginTop: '5px', marginBottom: '5px'}} danger
                                    onClick={(e) => {
                                        e.stopPropagation()
                                    }}>
                                Remove
                            </Button>

                        </Popconfirm>
                    </div>

                </Row>
                <FunctionLogsModal
                    func={this.props.func}
                    visible={this.state.functionLogsModalVisible}
                    hideModal={this.hideFunctionLogsModal}
                    instance={true}
                />
                <RunFunctionModal
                    func={this.props.func}
                    visible={this.state.functionRunModalVisible}
                    hideModal={this.hideFunctionRunModal}
                    instance={true}
                />
                <ViewFunctionMetricsModal
                    func={this.props.func}
                    visible={this.state.functionMetricsModalVisible}
                    hideModal={this.hideFunctionMetricsModal}
                    instance={true}
                />
            </Card>

        )
    }
}

export default FunctionItem;