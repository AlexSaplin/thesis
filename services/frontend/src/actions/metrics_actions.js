// import { metricsActionConstants } from '../constants';
// import { metricBaseURL } from "../config";
// import { loadState } from "../persist_store";
//
// import axios from "axios";
//
//
// const ONE_DAY = 86400000
//
//
// export function get_metrics_list(daysFromNow) {
//    return dispatch => {
//         dispatch({type: metricsActionConstants.GET_METRICS_REQUEST});
//         const state = loadState()
//         console.log("ooooooo")
//         console.log(state.authentication.user)
//         // const GET_METRICS_ENDPOINT = `http://${metricBaseURL}/api/v1/query?query={__name__=~"run_duration_sum|run_duration_count",owner_id="${state.authentication.user.uuid}"}`;
//         //
//         // return new Promise((resolve, reject) => {axios.get(GET_METRICS_ENDPOINT, {transformRequest: (data, headers) => {
//         //         delete headers.common['X-Token'];
//         //     }})
//         //     .then((response) => {
//         //         console.log("------")
//         //         console.log(response.data)
//         //         let functionsInfo = {}
//         //
//         //         const date = new Date()
//         //         const today = new Date(date.getFullYear(), date.getMonth(), date.getDate()).getTime()
//         //         // for (let i = 0; i < daysFromNow; ++i) {
//         //         //     functionsInfo[today - i * ONE_DAY] = {count: 0, time: 0}
//         //         // }
//         //
//         //         response.data.data.result.forEach((item, index) => {
//         //             let functionName = item.metric.function_id
//         //             if (!(functionName in functionsInfo)) {
//         //                 functionsInfo[functionName] = {count: 0, totalTime: 0, details: {}}
//         //                 for (let i = 0; i < daysFromNow; ++i) {
//         //                     functionsInfo[functionName].details[today - i * ONE_DAY] = {count: 0, totalTime: 0}
//         //                 }
//         //             }
//         //         })
//         //
//         //         response.data.data.result.forEach((item, index) => {
//         //             let functionName = item.metric.function_id
//         //             let metricsName = item.metric.__name__
//         //             let localTime = new Date(Number.parseInt(item.metric.date) * 1000)
//         //             let localDate = new Date(localTime.getFullYear(), localTime.getMonth(), localTime.getDate()).getTime()
//         //
//         //             if (metricsName === "run_duration_count") {
//         //                 functionsInfo[functionName].count += Number.parseInt(item.value[1])
//         //             } else if (metricsName === "run_duration_sum") {
//         //                 functionsInfo[functionName].totalTime += Number.parseFloat(item.value[1])
//         //             }
//         //
//         //             if (localDate <= today - daysFromNow * ONE_DAY) {
//         //                 return
//         //             }
//         //
//         //             if (metricsName === "run_duration_count") {
//         //                 functionsInfo[functionName].details[localDate].count += Number.parseInt(item.value[1])
//         //             } else if (metricsName === "run_duration_sum") {
//         //                 functionsInfo[functionName].details[localDate].totalTime += Number.parseFloat(item.value[1])
//         //             }
//         //         })
//         //         console.log("+------+")
//         //         console.log(functionsInfo)
//         //
//         //         dispatch({type: metricsActionConstants.GET_METRICS_SUCCESS, payload: functionsInfo});
//         //         resolve(response);
//         //     })
//         //     .catch((error) => {
//         //         dispatch({type: metricsActionConstants.GET_METRICS_FAILURE});
//         //         reject(error);
//         //     });
//         // });
//        const GET_METRICS_ENDPOINT = `http://${metricBaseURL}/api/v1/query?query={__name__=~"run_duration_sum|run_duration_count",owner_id="${state.authentication.user.uuid}"}`;
//
//        return new Promise((resolve, reject) => {axios.get(GET_METRICS_ENDPOINT, {transformRequest: (data, headers) => {
//                delete headers.common['X-Token'];
//            }})
//            .then((response) => {
//                console.log("------")
//                console.log(response.data)
//                let functionsInfo = {}
//
//                const date = new Date()
//                const today = new Date(date.getFullYear(), date.getMonth(), date.getDate()).getTime()
//                // for (let i = 0; i < daysFromNow; ++i) {
//                //     functionsInfo[today - i * ONE_DAY] = {count: 0, time: 0}
//                // }
//
//                response.data.data.result.forEach((item, index) => {
//                    let functionName = item.metric.function_id
//                    if (!(functionName in functionsInfo)) {
//                        functionsInfo[functionName] = {count: 0, totalTime: 0, details: {}}
//                        for (let i = 0; i < daysFromNow; ++i) {
//                            functionsInfo[functionName].details[today - i * ONE_DAY] = {count: 0, totalTime: 0}
//                        }
//                    }
//                })
//
//                response.data.data.result.forEach((item, index) => {
//                    let functionName = item.metric.function_id
//                    let metricsName = item.metric.__name__
//                    // let localTime = new Date(Number.parseInt(item.metric.date) * 1000)
//                    // let localDate = new Date(localTime.getFullYear(), localTime.getMonth(), localTime.getDate()).getTime()
//
//                    if (metricsName === "run_duration_count") {
//                        functionsInfo[functionName].count += Number.parseInt(item.value[1])
//                    } else if (metricsName === "run_duration_sum") {
//                        functionsInfo[functionName].totalTime += Number.parseFloat(item.value[1])
//                    }
//
//                    // if (localDate <= today - daysFromNow * ONE_DAY) {
//                    //     return
//                    // }
//
//                    // if (metricsName === "run_duration_count") {
//                    //     functionsInfo[functionName].details[localDate].count += Number.parseInt(item.value[1])
//                    // } else if (metricsName === "run_duration_sum") {
//                    //     functionsInfo[functionName].details[localDate].totalTime += Number.parseFloat(item.value[1])
//                    // }
//
//                    // const GET_DETAILS = `http://${metricBaseURL}/api/v1/query?query={__name__=~"run_duration_sum|run_duration_count",owner_id="${state.authentication.user.uuid}"}`;
//                    // new Promise((resolve, reject) => {axios.get(GET_DETAILS, {transformRequest: (data, headers) => {
//                    //         delete headers.common['X-Token'];
//                    //     }}).then((response) => {
//                    //
//                    // }
//                    // }
//                })
//                console.log("+------+")
//                console.log(functionsInfo)
//
//                dispatch({type: metricsActionConstants.GET_METRICS_SUCCESS, payload: functionsInfo});
//                resolve(response);
//            })
//            .catch((error) => {
//                dispatch({type: metricsActionConstants.GET_METRICS_FAILURE});
//                reject(error);
//            });
//        });
//    }
// }
