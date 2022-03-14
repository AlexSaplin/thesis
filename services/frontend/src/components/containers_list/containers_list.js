import React from "react";
import {connect} from "react-redux";
import {Empty, Row, Spin} from "antd";
import {LoadingOutlined} from '@ant-design/icons';
import {ContainerActions} from "../../actions";
import ContainerItem from "./container_item";


class ContainersListUnwrapped extends React.Component {
  componentDidMount() {
    this.props.get_containers_list()
  }



  render() {
    const { containers = [], processing } = this.props;


    const containersView = containers.length !== 0 ?
      containers.map((container, index) => {
          return (<ContainerItem key={index} container={container} reloadParent={this.props.get_containers_list}/>);
        }
      )
      : (
        <Row align="middle" justify="center" style={{height: "100%"}}>
          <Empty description="No containers yet" />
        </Row>
      );

    return processing ? (
      <Row align="middle" justify="center" style={{height: "100%"}}>
        <Spin indicator={<LoadingOutlined style={{ fontSize: 48 }} spin />} />
      </Row>
    ) : containersView;
  }
}

function mapStateToProps(state) {
  return {
    containers: state.container.containers,
    processing: state.container.processing,
  };
}

const actionCreators = {
  get_containers_list: ContainerActions.get_containers_list,
  remove_container: ContainerActions.remove_container,
};

const ContainersList = connect(
  mapStateToProps,
  actionCreators
)(ContainersListUnwrapped);

export default ContainersList;
