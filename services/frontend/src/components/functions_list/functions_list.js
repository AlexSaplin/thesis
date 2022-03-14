import React from "react";
import {connect} from "react-redux";
import {Button, Empty, Row, Spin} from "antd";

import {FunctionActions} from "../../actions";
import {LoadingOutlined} from '@ant-design/icons';

import './functions_list.css'
import FunctionItem from "./function_item";


class FunctionsListUnwrapped extends React.Component {
    componentDidMount() {
        this.props.get_functions_list()
    }

    render() {
        const {functions = [], processing} = this.props;

        const functionsView = functions.length !== 0 ?
            functions.map((func, index) => {
                    return (<FunctionItem key={index} func={func} reloadParent={this.props.get_functions_list}
                                          remove_function={this.props.remove_function}/>);
                }
            )
            : (
                <Row align="middle" justify="center" style={{height: "100%"}}>
                    <Empty description="No functions yet">
                        <Button type="primary">
                            <a href="https://deep-mux.github.io/functions-quickstart/" target="_blank"
                               rel="noopener noreferrer">Get Started</a></Button>
                    </Empty>
                </Row>
            );

        return processing ? (
            <Row align="middle" justify="center" style={{height: "100%"}}>
                <Spin indicator={<LoadingOutlined style={{fontSize: 48}} spin/>}/>
            </Row>
        ) : functionsView;
    }
}

function mapStateToProps(state) {
    return {
        functions: state.func.funcs,
        processing: state.func.processing,
    };
}

const actionCreators = {
    get_functions_list: FunctionActions.get_functions_list,
    remove_function: FunctionActions.remove_function,
};

const FunctionsList = connect(
    mapStateToProps,
    actionCreators
)(FunctionsListUnwrapped);

export default FunctionsList;
