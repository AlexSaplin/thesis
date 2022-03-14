import React from "react";
import {Select, Row, Col} from "antd";
import {
    XYPlot,
    XAxis,
    YAxis,
    VerticalGridLines,
    HorizontalGridLines,
    makeWidthFlexible,
    VerticalRectSeries,
    LineSeries
} from "react-vis";
import moment from "moment";


const FlexibleXYPlot = makeWidthFlexible(XYPlot);


// class ChartInternal extends React.Component {
//     state = {
//         currentFunction: "##total",
//         currentChartType: "columns"
//     };
//
//     changeCurrentFunction = value => {
//         this.setState({currentFunction: value})
//     }
//
//     changeCurrentChartType = value => {
//         this.setState({currentChartType: value})
//     }
//
//     tickFormat = value => {
//         let m = moment(value)
//         if (m.hours() === 0 && m.minutes() === 0 && m.seconds() === 0 && m.milliseconds() === 0) {
//             return m.format('DD.MM')
//         } else {
//             return ""
//         }
//     }
//
//     render() {
//         const ONE_DAY = 86400000
//         let currentChart = <VerticalRectSeries data={this.props.data[this.state.currentFunction]} style={{stroke: '#fff'}}/>
//         if (this.state.currentChartType === "line") {
//             currentChart = <LineSeries data={this.props.noWidthData[this.state.currentFunction]} style={{ fill: 'none' }}/>
//         }
//         return (
//             <>
//                 <Select defaultValue={this.state.currentFunction} style={{ width: 120, marginBottom: 10, right: 0 }} onChange={this.changeCurrentFunction}>
//                     <Select.Option value="##total">Total</Select.Option>
//                     {this.props.functionNames.map((val, index) => <Select.Option value={val} key={index}>{val}</Select.Option>)}
//                 </Select>
//                 <Select defaultValue={this.state.currentChartType} style={{ width: 120, marginBottom: 10, marginLeft: 10, right: 0 }} onChange={this.changeCurrentChartType}>
//                     <Select.Option value="columns">Columns</Select.Option>
//                     <Select.Option value="line">Line</Select.Option>
//                 </Select>
//                 <FlexibleXYPlot xDomain={[this.props.beginPeriod, this.props.endPeriod]} yDomain={[0, 5]} xType="time" height={200}>
//                     <VerticalGridLines />
//                     <HorizontalGridLines />
//                     <XAxis tickFormat={this.tickFormat}/>
//                     <YAxis />
//                     {currentChart}
//                 </FlexibleXYPlot>
//             </>
//         );
//     }
// }
//
//
// class Chart extends React.Component {
//     generateDictForInterval(daysFromNow) {
//         const ONE_DAY = 86400000
//         const date = new Date()
//         const today = new Date(date.getFullYear(), date.getMonth(), date.getDate()).getTime()
//         console.log("sdfdsf")
//         console.log(today)
//         let data = {}
//         for (let i = 0; i < daysFromNow; ++i) {
//             data[today - i * ONE_DAY] = 0
//         }
//         return data
//     }
//
//     setDataDictFromArray(dataDict, dataArray) {
//         dataArray.forEach(function (item, index) {
//             if (dataDict[item.date] !== undefined) {
//                 dataDict[item.date] += item.y
//             }
//         })
//     }
//
//     dictToArray(dataDict) {
//         let dataArray = []
//
//         for (const key in dataDict) {
//             if (dataDict.hasOwnProperty(key)) {
//                 dataArray.push({date: parseInt(key), y: dataDict[key]})
//             }
//         }
//
//         dataArray.sort((a, b) => a.date - b.date)
//
//         return dataArray
//     }
//
//     configureDataArray(dataArray, daysFromNow) {
//         let dataDict = this.generateDictForInterval(daysFromNow)
//         this.setDataDictFromArray(dataDict, dataArray)
//         return this.dictToArray(dataDict)
//     }
//
//     render() {
//         const ONE_DAY = 86400000
//
//         let totalData = this.generateDictForInterval(10)
//         let processedData = {}
//         let processedTotal = []
//
//         let processedDataNoWidth = {}
//         let processedTotalNoWidth = []
//
//         let functionNames = []
//
//         for (const key in this.props.data) {
//             if (this.props.data.hasOwnProperty(key)) {
//                 let processedArray = []
//                 let processedArrayNoWidth = []
//                 let initialArray = this.configureDataArray(this.props.data[key], 10)
//                 this.setDataDictFromArray(totalData, initialArray)
//                 initialArray.forEach(function (item, index) {
//                     processedArray.push({x0: item.date - ONE_DAY * 0.4, x: item.date + ONE_DAY * 0.4, y: item.y})
//                     processedArrayNoWidth.push({x0: item.date, x: item.date, y: item.y})
//                 })
//                 processedData[key] = processedArray
//                 processedDataNoWidth[key] = processedArrayNoWidth
//                 functionNames.push(key)
//             }
//         }
//
//         for (const key in totalData) {
//             if (totalData.hasOwnProperty(key)) {
//                 processedTotal.push({x0: parseInt(key) - ONE_DAY * 0.4, x: parseInt(key) + ONE_DAY * 0.4, y: totalData[key]})
//                 processedTotalNoWidth.push({x0: parseInt(key), x: parseInt(key), y: totalData[key]})
//             }
//         }
//
//         processedTotal.sort((a, b) => a.x0 - b.x0)
//         processedTotalNoWidth.sort((a, b) => a.x0 - b.x0)
//         processedData["##total"] = processedTotal
//         processedDataNoWidth["##total"] = processedTotalNoWidth
//
//         const endPeriod = (new Date()).getTime()
//         const beginPeriod = endPeriod - ONE_DAY * 20
//
//         console.log(processedData)
//
//         return (
//             <ChartInternal data={processedData} noWidthData={processedDataNoWidth} functionNames={functionNames} beginPeriod={beginPeriod} endPeriod={endPeriod}/>
//         );
//     }
// }


const ONE_DAY = 86400000
const NEAR_HALF_DAY = 86400000 * 0.4

const STATISTIC_VARIANTS = {
    totalTime: array => {
        return Array.from(array, element => {return {date: element.date, y: element.stats.totalTime}})
    },
    count: array => {
        return Array.from(array, element => {return {date: element.date, y: element.stats.count}})
    },
    averageTime: array => {
        return Array.from(array, element => {return {date: element.date, y: element.stats.totalTime / element.stats.count || 0}})
    }
}


const CHART_VARIANT = {
    columns: array => {
        return Array.from(array, element => {return {x0: element.date - NEAR_HALF_DAY, x: element.date + NEAR_HALF_DAY, y: element.y}})
    },
    line: array => {
        return Array.from(array, element => {return {x: element.date, y: element.y}})
    }
}


class ChartInternal extends React.Component {
    state = {
        currentChartType: "columns",
        currentStatisticsType: "totalTime",
        period: 14
    };

    changeCurrentFunction = value => {
        this.setState({currentFunction: value})
    }

    changeCurrentChartType = value => {
        this.setState({currentChartType: value})
    }

    changeCurrentStatisticsType = value => {
        this.setState({currentStatisticsType: value})
    }

    changeCurrentPeriod = value => {
        this.setState({period: Number.parseInt(value)})
    }

    tickFormat = value => {
        let m = moment(value)
        if (m.hours() === 0 && m.minutes() === 0 && m.seconds() === 0 && m.milliseconds() === 0) {
            return m.format('DD.MM')
        } else {
            return ""
        }
    }

    render() {
        console.log("utauta")
        console.log(this.props.data)
        const slicedData = CHART_VARIANT[this.state.currentChartType](STATISTIC_VARIANTS[this.state.currentStatisticsType](this.props.data.details.slice(-this.state.period)))
        console.log(slicedData)
        let maxY = 0.1
        slicedData.forEach(function (item, index) {
            maxY = Math.max(maxY, item.y)
        });
        if (this.state.currentStatisticsType === "count") {
            maxY = Math.max(maxY, 5)
            maxY = Math.ceil(maxY)
        }
        let currentChart = null
        let beginPeriod = 0
        let endPeriod = 0
        if (this.state.currentChartType === "columns") {
            currentChart = <VerticalRectSeries data={slicedData} style={{stroke: '#fff', fill: "#041527"}}/>
            beginPeriod = slicedData[0].x - ONE_DAY
            endPeriod = slicedData[slicedData.length - 1].x
        } else {
            currentChart = <LineSeries data={slicedData} style={{ fill: 'none', stroke: "#041527" }}/>
            beginPeriod = slicedData[0].x - NEAR_HALF_DAY
            endPeriod = slicedData[slicedData.length - 1].x + NEAR_HALF_DAY
        }

        return (
            <>
                <div className="ant-statistic-title" style={{fontSize: 18}}>
                    Details
                    <Select defaultValue={this.state.currentChartType} style={{ width: 120, marginBottom: 10, marginLeft: 10, right: 0 }} onChange={this.changeCurrentChartType}>
                        <Select.Option value="columns">Columns</Select.Option>
                        <Select.Option value="line">Line</Select.Option>
                    </Select>
                    <Select defaultValue={this.state.currentStatisticsType} style={{ width: 130, marginBottom: 10, marginLeft: 10, right: 0 }} onChange={this.changeCurrentStatisticsType}>
                        <Select.Option value="totalTime">Total time</Select.Option>
                        <Select.Option value="count">Run count</Select.Option>
                        <Select.Option value="averageTime">Average time</Select.Option>
                    </Select>
                    <Select defaultValue={String(this.state.period)} style={{ width: 120, marginBottom: 10, marginLeft: 10, right: 0 }} onChange={this.changeCurrentPeriod}>
                        <Select.Option value="7">Week</Select.Option>
                        <Select.Option value="14">Two weeks</Select.Option>
                        <Select.Option value="30">Month</Select.Option>
                    </Select>
                </div>
                <FlexibleXYPlot xDomain={[beginPeriod, endPeriod]} yDomain={[0, maxY]} xType="time" height={300} style={{padding: "10px"}}>
                    <VerticalGridLines />
                    <HorizontalGridLines />
                    <XAxis tickFormat={this.tickFormat}/>
                    <YAxis />
                    {currentChart}
                </FlexibleXYPlot>
            </>
        );
    }
}


class Chart extends React.Component {
    render() {
        return (
            <ChartInternal data={this.props.data} beginPeriod={this.props.beginPeriod} endPeriod={this.props.endPeriod}/>
        );
    }
}



export default Chart
