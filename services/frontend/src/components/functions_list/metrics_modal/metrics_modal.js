import React from "react";
import {Alert, Card, Col, Form, Input, Modal, Row, Select, Spin, Statistic} from "antd";
import {internalApiURL} from "../../../config";
import axios from "axios";
import {FunctionActions} from "../../../actions";
import {connect} from "react-redux";
import Editor from "@monaco-editor/react";
import JSZip from "jszip"
import Chart from "../../statistics/chart";
import {LoadingOutlined} from "@ant-design/icons";
import {
    HorizontalGridLines,
    LineSeries,
    makeWidthFlexible,
    VerticalGridLines,
    VerticalRectSeries,
    XAxis, XYPlot,
    YAxis
} from "react-vis";
import moment from "moment";


const ONE_DAY = 86400000
const NEAR_HALF_DAY = 86400000 * 0.4
const FlexibleXYPlot = makeWidthFlexible(XYPlot);


function isoStringToDateOnlyTimestamp(date) {
    const isoDateObj = new Date(date);
    isoDateObj.setHours(0);
    isoDateObj.setMinutes(0);
    isoDateObj.setSeconds(0);
    const dateOnly = new Date(isoDateObj.getFullYear(), isoDateObj.getMonth(), isoDateObj.getDate())
    return dateOnly.getTime()
}


const STATISTIC_VARIANTS = {
    totalRunTime: array => {
        return Array.from(array, element => {return {date: isoStringToDateOnlyTimestamp(element.day), y: element.runTime.toFixed(3)}})
    },
    runCount: array => {
        return Array.from(array, element => {return {date: isoStringToDateOnlyTimestamp(element.day), y: element.runCount}})
    },
    averageRunTime: array => {
        return Array.from(array, element => {return {date: isoStringToDateOnlyTimestamp(element.day), y: (element.runTime / element.runCount || 0).toFixed(3)}})
    },
    totalLoadTime: array => {
        return Array.from(array, element => {return {date: isoStringToDateOnlyTimestamp(element.day), y: element.loadTime.toFixed(3)}})
    },
    loadCount: array => {
        return Array.from(array, element => {return {date: isoStringToDateOnlyTimestamp(element.day), y: element.loadCount}})
    },
    averageLoadTime: array => {
        return Array.from(array, element => {return {date: isoStringToDateOnlyTimestamp(element.day), y: (element.loadTime / element.loadCount || 0).toFixed(3)}})
    },
}


const CHART_VARIANT = {
    columns: array => {
        return Array.from(array, element => {return {x0: element.date - NEAR_HALF_DAY, x: element.date + NEAR_HALF_DAY, y: element.y}})
    },
    line: array => {
        return Array.from(array, element => {return {x: element.date, y: element.y}})
    }
}


class ViewFunctionMetricsModalUnwrapped extends React.Component {

    state = {
        confirmLoading: false,
        generalFormError: "",
        currentLoadStatus: 0, // 0 - not started, 1 - loading, 2 - success, 3 - error
        functionInfo: null,
        modifiedFunctionInfo: null,
        currentChartType: "columns",
        currentStatisticsType: "totalRunTime",
        period: 14
    };

    changeCurrentChartType = value => {
        this.setState({currentChartType: value})
    }

    changeCurrentStatisticsType = value => {
        this.setState({currentStatisticsType: value})
    }

    changeCurrentPeriod = value => {
        this.setState({period: Number.parseInt(value)})
    }

    handleCancel = () => {
        this.props.hideModal();
    };

    onCodeChange = (newValue, _) => {
        this.setState({
            codeValue: newValue,
        });
    };

    loadData = async () => {
        const endpoint = `${internalApiURL}/v1/function/${this.props.func.name}/metrics?offset=30`
        // const endpoint = `${internalApiURL}/v1/function/newFunction`

        this.setState({currentLoadStatus: 1});

        try {
            let res = await axios.get(endpoint)

            let data = res.data

            // data = {
            //     totalRunTime: 12.345345,
            //     totalRunCount: 8,
            //     totalLoadTime: 55.55,
            //     totalLoadCount: 4,
            //     details: [
            //         {
            //             day: "2021-04-28T19:31:00.000Z",
            //             runTime: 8.555,
            //             runCount: 5,
            //             loadTime: 20.6,
            //             loadCount: 3
            //         },
            //         {
            //             day: "2021-04-27T19:31:00.000Z",
            //             runTime: 5.555,
            //             runCount: 3,
            //             loadTime: 20.6,
            //             loadCount: 3
            //         }
            //     ]
            // }

            data.details.sort(function(a, b) {
                return (new Date(a.day)).getTime() - (new Date(b.day)).getTime();
            });

            this.setState({currentLoadStatus: 2, functionInfo: data});
        } catch (error) {
            console.log(error)
            console.log(error.constructor.name)
            try {
                if (error.response.status === 400) {
                    this.setState({
                        generalFormError: error.response.data.message,
                        currentLoadStatus: 3,
                    })
                }
                if (error.response.status === 409) {
                    this.setState({
                        generalFormError: "Function with this name already exists",
                        currentLoadStatus: 3,
                    })
                } else {
                    this.setState({
                        generalFormError: error.toString(),
                        currentLoadStatus: 3,
                    })
                }
            } catch (_) {
                this.setState({
                    generalFormError: "Unknown error",
                    currentLoadStatus: 3
                })
            }
        }
    }

    tickFormat = value => {
        let m = moment(value)
        if (m.hours() === 0 && m.minutes() === 0 && m.seconds() === 0 && m.milliseconds() === 0) {
            return m.format('DD.MM')
        } else {
            return ""
        }
    }

    yTickFormat = value => {

        if (this.state.currentStatisticsType === "runCount" || this.state.currentStatisticsType === "loadCount") {
            return Math.round(value) === value ? value : ""
        } else {
            return value
        }
    }

    render() {
        const {confirmLoading} = this.state;

        if (this.props.visible === false) {
            if (this.state.currentLoadStatus !== 0) {
                this.setState({currentLoadStatus: 0})
            }
            return <Modal title="Metrics" visible={this.props.visible} onOk={this.handleOk}
                          confirmLoading={confirmLoading} onCancel={this.handleCancel} width="90%"
                          style={{minHeight: "80%"}} centered={true}>
                <Row align="middle" justify="center" style={{height: "100%"}}>
                    <Spin indicator={<LoadingOutlined style={{fontSize: 48}} spin/>}/>
                </Row>
            </Modal>
        }

        if (this.state.currentLoadStatus === 0) {
            this.loadData().then(r => {})
        }

        if (this.state.currentLoadStatus === 0 || this.state.currentLoadStatus === 1) {
            return <Modal title="Metrics" visible={this.props.visible} onOk={this.handleCancel}
                   confirmLoading={confirmLoading} onCancel={this.handleCancel} width="90%"
                   style={{minHeight: "80%"}} centered={true} okButtonProps={{ style: { display: 'none' } }} cancelText="Close">
                <Row align="middle" justify="center" style={{height: "100%"}}>
                    <Spin indicator={<LoadingOutlined style={{fontSize: 48}} spin/>}/>
                </Row>
            </Modal>
        } else if (this.state.currentLoadStatus === 3) {
            return <Modal title="Metrics" visible={this.props.visible} onOk={this.handleCancel}
                          confirmLoading={confirmLoading} onCancel={this.handleCancel} width="90%"
                          style={{minHeight: "80%"}} centered={true} okButtonProps={{ style: { display: 'none' } }} cancelText="Close">
                <Alert message={this.state.generalFormError} type="error" showIcon style={{marginBottom: "10px"}}/>
            </Modal>

        }

        let currentTimestamp = new Date()
        const today = new Date(currentTimestamp.getFullYear(), currentTimestamp.getMonth(), currentTimestamp.getDate()).getTime()

        const slicedData = CHART_VARIANT[this.state.currentChartType](STATISTIC_VARIANTS[this.state.currentStatisticsType](this.state.functionInfo.details.slice(-this.state.period)))
        let maxY = 0
        slicedData.forEach(function (item, index) {
            maxY = Math.max(maxY, item.y)
        });
        if (this.state.currentStatisticsType === "count") {
            maxY = Math.max(maxY, 5)
            maxY = Math.ceil(maxY)
        }
        maxY += 1
        let currentChart = null
        let beginPeriod = 0
        let endPeriod = 0

        if (this.state.currentChartType === "columns") {
            currentChart = <VerticalRectSeries data={slicedData} style={{stroke: '#fff', fill: "#041527"}}/>
            beginPeriod = today - this.state.period * ONE_DAY + NEAR_HALF_DAY
            endPeriod = today + NEAR_HALF_DAY
        } else {
            currentChart = <LineSeries data={slicedData} style={{ fill: 'none', stroke: "#041527" }}/>
            beginPeriod = today - this.state.period * ONE_DAY + NEAR_HALF_DAY
            endPeriod = today + NEAR_HALF_DAY
        }

        return (
            <div>
                <Modal title={`Metrics of ${this.props.func.name}`} visible={this.props.visible} onOk={this.handleCancel}
                       confirmLoading={confirmLoading} onCancel={this.handleCancel} width="90%"
                       style={{minHeight: "80%"}} centered={true} okButtonProps={{ style: { display: 'none' } }} cancelText="Close">
                        <Row gutter={16} style={{marginTop: "10px"}}>
                            <Col span={8}>
                                <Card>
                                    <Statistic title="Total run time" value={`${this.state.functionInfo.totalRunTime.toFixed(3)} seconds`}/>
                                </Card>
                            </Col>
                            <Col span={8}>
                                <Card>
                                    <Statistic title="Total run count" value={this.state.functionInfo.totalRunCount}/>
                                </Card>
                            </Col>
                            <Col span={8}>
                                <Card>
                                    <Statistic title="Average run time" value={`${(this.state.functionInfo.totalRunTime / this.state.functionInfo.totalRunCount || 0).toFixed(3)} seconds`}/>
                                </Card>
                            </Col>
                        </Row>
                    {/*<Row gutter={16} style={{marginTop: "10px"}}>*/}
                    {/*    <Col span={8}>*/}
                    {/*        <Card>*/}
                    {/*            <Statistic title="Total load time" value={`${this.state.functionInfo.totalLoadTime.toFixed(3)} seconds`}/>*/}
                    {/*        </Card>*/}
                    {/*    </Col>*/}
                    {/*    <Col span={8}>*/}
                    {/*        <Card>*/}
                    {/*            <Statistic title="Total load count" value={this.state.functionInfo.totalLoadCount}/>*/}
                    {/*        </Card>*/}
                    {/*    </Col>*/}
                    {/*    <Col span={8}>*/}
                    {/*        <Card>*/}
                    {/*            <Statistic title="Average load time" value={`${(this.state.functionInfo.totalLoadTime / this.state.functionInfo.totalLoadCount || 0).toFixed(3)} seconds`}/>*/}
                    {/*        </Card>*/}
                    {/*    </Col>*/}
                    {/*</Row>*/}
                            <Card style={{marginTop: 10}}>
                                <div className="ant-statistic-title" style={{fontSize: 18}}>
                                    Details
                                    <Select defaultValue={this.state.currentChartType} style={{ width: 120, marginBottom: 10, marginLeft: 10, right: 0 }} onChange={this.changeCurrentChartType}>
                                        <Select.Option value="columns">Columns</Select.Option>
                                        <Select.Option value="line">Line</Select.Option>
                                    </Select>
                                    <Select defaultValue={this.state.currentStatisticsType} style={{ width: 160, marginBottom: 10, marginLeft: 10, right: 0 }} onChange={this.changeCurrentStatisticsType}>
                                        <Select.Option value="totalRunTime">Total run time</Select.Option>
                                        <Select.Option value="runCount">Run count</Select.Option>
                                        <Select.Option value="averageRunTime">Average run time</Select.Option>
                                        {/*<Select.Option value="totalLoadTime">Total load time</Select.Option>*/}
                                        {/*<Select.Option value="loadCount">Load count</Select.Option>*/}
                                        {/*<Select.Option value="averageLoadTime">Average load time</Select.Option>*/}
                                    </Select>
                                    <Select defaultValue={String(this.state.period)} style={{ width: 120, marginBottom: 10, marginLeft: 10, right: 0 }} onChange={this.changeCurrentPeriod}>
                                        <Select.Option value="7">Week</Select.Option>
                                        <Select.Option value="14">Two weeks</Select.Option>
                                        <Select.Option value="30">Month</Select.Option>
                                    </Select>
                                </div>
                                <FlexibleXYPlot xDomain={[beginPeriod, endPeriod]} yDomain={[0, maxY]} xType="time" height={400} style={{padding: "5px"}}>
                                    <VerticalGridLines />
                                    <HorizontalGridLines />
                                    <XAxis tickFormat={this.tickFormat}/>
                                    <YAxis tickFormat={this.yTickFormat}/>
                                    {currentChart}
                                </FlexibleXYPlot>
                            </Card>

                </Modal>
            </div>
        );
    }
}


const actionCreators = {
    get_functions_list: FunctionActions.get_functions_list,
};

const ViewFunctionMetricsModal = connect((_) => {}, actionCreators,)(ViewFunctionMetricsModalUnwrapped);

export default ViewFunctionMetricsModal;
