import React from "react";
import { connect } from "react-redux";
import {Statistic, Card, Row, Col, Select, Spin} from "antd";
import {LoadingOutlined} from "@ant-design/icons";
import {FunctionActions, MetricsActions} from '../../actions';
import {XYPlot, makeWidthFlexible} from "react-vis";
import Chart from "./chart"


const FlexibleXYPlot = makeWidthFlexible(XYPlot);


class StatisticsUnwrapped extends React.Component {
    state = {
        currentFunction: "##total",
        currentChartType: "columns"
    };

    componentDidMount() {
        this.props.get_functions_list()
        this.props.get_metrics_list(30)
    }

    changeCurrentFunction = value => {
        this.setState({currentFunction: value})
    }

    changeCurrentChartType = value => {
        this.setState({currentChartType: value})
    }

    generateDictForInterval(daysFromNow) {
        const ONE_DAY = 86400000
        const date = new Date()
        const today = new Date(date.getFullYear(), date.getMonth(), date.getDate()).getTime()
        let data = {}
        for (let i = 0; i < daysFromNow; ++i) {
            data[today - i * ONE_DAY] = {count: 0, time: 0}
        }
        return data
    }

    setDataDictFromArray(dataDict, dataArray) {

        dataArray.forEach(function (item, index) {
            if (dataDict[item.date] !== undefined) {
                dataDict[item.date].count += item.stats.count
                dataDict[item.date].time += item.stats.time
            }
        })
    }

    dictToArray(dataDict) {
        let dataArray = []

        for (const key in dataDict) {
            if (dataDict.hasOwnProperty(key)) {
                dataArray.push({date: parseInt(key), stats: dataDict[key]})
            }
        }

        dataArray.sort((a, b) => a.date - b.date)

        return dataArray
    }

    configureDataArray(dataArray, daysFromNow) {
        let dataDict = this.generateDictForInterval(daysFromNow)
        this.setDataDictFromArray(dataDict, dataArray)
        return this.dictToArray(dataDict)
    }

    render() {
        console.log("pppp")
        console.log(this.props.functionsProcessing)
        if (this.props.metricsProcessing === 1 || this.props.functionsProcessing === 1) {
            return <Row align="middle" justify="center" style={{height: "100%"}}>
                <Spin indicator={<LoadingOutlined style={{fontSize: 48}} spin/>}/>
            </Row>
        }

        console.log("7-7-7-7-7")
        console.log(this.props.metrics)
        console.log(this.props.functions)
        let functionTranslatorDict = {}
        const newData = {}
        let functionNames = []

        this.props.functions.forEach((item, index) => {
            functionTranslatorDict[item.id] = item.name
            newData[item.name] = {count: 0, totalTime: 0, details: []}
            functionNames.push(item.name)
        })

        for (let key in this.props.metrics) {
            if (this.props.metrics.hasOwnProperty(key)) {
                let dataArray = []
                let details = this.props.metrics[key].details

                for (const date in details) {
                    if (details.hasOwnProperty(date)) {
                        dataArray.push({date: parseInt(date), stats: details[date]})
                    }
                }

                dataArray.sort((a, b) => a.date - b.date)

                newData[functionTranslatorDict[key]].details = dataArray
                newData[functionTranslatorDict[key]].count = this.props.metrics[key].count
                newData[functionTranslatorDict[key]].totalTime = this.props.metrics[key].totalTime
            }
        }

        console.log("u-u-u-u-u")
        console.log(newData)

        const date = new Date()
        const ONE_DAY = 86400000;
        const timestamp = new Date(date.getFullYear(), date.getMonth(), date.getDate()).getTime() - 9 * ONE_DAY

        // const dataa = {
        //     func1: [
        //         {date: timestamp + ONE_DAY * 3, stats: {count: 1, time: 3.3}},
        //         {date: timestamp + ONE_DAY * 7, stats: {count: 2, time: 4.3}},
        //     ],
        //     func2: [
        //         {date: timestamp + ONE_DAY * 4, stats: {count: 1, time: 3.3}},
        //         {date: timestamp + ONE_DAY * 7, stats: {count: 1, time: 3.3}},
        //     ]
        // }

        const daysFromNow = 30

        const today = new Date(date.getFullYear(), date.getMonth(), date.getDate()).getTime()

        let totalData = this.generateDictForInterval(daysFromNow)

        console.log("ppppp")
        // console.log(dataa)

        let processedData = {}
        let processedTotal = []



        let totalTime = 0
        let totalCount = 0
        let totalDetailsDict = {}
        for (let i = 0; i < daysFromNow; ++i) {
            totalDetailsDict[today - i * ONE_DAY] = {count: 0, totalTime: 0}
        }

        for (const key in newData) {
            if (newData.hasOwnProperty(key)) {
                totalTime += newData[key].totalTime
                totalCount += newData[key].count

                newData[key].details.forEach(function (item, index) {
                    if (!(item.date in totalDetailsDict)) {
                        totalDetailsDict[item.date] = {count: 0, totalTime: 0}
                    }
                    totalDetailsDict[item.date].count += item.stats.count
                    totalDetailsDict[item.date].totalTime += item.stats.totalTime
                })
            }
        }

        for (const key in newData) {
            if (newData.hasOwnProperty(key)) {
                if (newData[key].details.length === 0) {
                    for (let i = 1; i <= daysFromNow; ++i) {
                        newData[key].details.push({date: today - daysFromNow * ONE_DAY + i * ONE_DAY, stats: {count: 0, totalTime: 0}})
                    }
                }
            }
        }


        newData["##total"] = {count: totalCount, totalTime: totalTime, details: this.dictToArray(totalDetailsDict)}

        console.log("????")
        console.log(newData)


        // for (const key in totalData) {
        //     if (totalData.hasOwnProperty(key)) {
        //         processedTotal.push({date: parseInt(key), stats: totalData[key]})
        //     }
        // }
        //
        // processedTotal.sort((a, b) => a.date - b.date)
        // processedData["##total"] = processedTotal

        const endPeriod = (new Date()).getTime()
        const beginPeriod = endPeriod - ONE_DAY * daysFromNow

        return (
            <>
                <Card title="">
                    <h2 style={{ display: "inline" }}>Functions usage</h2>
                    <Select showSearch defaultValue={this.state.currentFunction} style={{ width: 300, marginBottom: 10, right: 0, marginLeft: "20px" }} onChange={this.changeCurrentFunction}>
                        <Select.Option value="##total">Total</Select.Option>
                        {functionNames.map((val, index) => <Select.Option value={val} key={index}>{val}</Select.Option>)}
                    </Select>
                </Card>
                <Row gutter={16} style={{marginTop: "10px"}}>
                    <Col span={8}>
                        <Card>
                            <Statistic title="Total time" value={newData[this.state.currentFunction].totalTime}/>
                        </Card>
                    </Col>
                    <Col span={8}>
                        <Card>
                            <Statistic title="Total count" value={newData[this.state.currentFunction].count}/>
                        </Card>
                    </Col>
                    <Col span={8}>
                        <Card>
                            <Statistic title="Average time" value={newData[this.state.currentFunction].totalTime / newData[this.state.currentFunction].count || 0}/>
                        </Card>
                    </Col>
                </Row>
                <Card style={{marginTop: 10}}>
                    <Chart data={newData[this.state.currentFunction]} beginPeriod={beginPeriod} endPeriod={endPeriod}/>
                </Card>
            </>
        );
    };
}

function mapStateToProps(state) {
    return {
        functions: state.func.funcs,
        functionsProcessing: state.func.processing,
        metrics: state.metrics.metrics,
        metricsProcessing: state.metrics.processing

    };
}

const actionCreators = {
    get_functions_list: FunctionActions.get_functions_list,
    get_metrics_list: MetricsActions.get_metrics_list
};

const Statistics = connect(mapStateToProps, actionCreators)(StatisticsUnwrapped);

export default Statistics;
