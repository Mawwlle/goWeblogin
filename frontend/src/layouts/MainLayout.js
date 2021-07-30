import React from "react";
import {Col, Layout, Row} from "antd";
import {Navbar} from "../components/Navbar";
import {BrowserRouter, Route, Switch} from "react-router-dom";
import {Home} from "../pages/Home";
import {Login} from "../pages/Login";
import {TaskModal} from "../components/TaskModal";

const {Header, Content, Footer} = Layout;

export const MainLayout = () => {
    return (
        <BrowserRouter>
            <Layout className="layout">

                <Header>
                    <div className="logo"/>
                    <Navbar/>
                </Header>
                <Row justify="center">
                    <Col span={18}>
                        <Content>
                            <Switch>
                                <Route path={'/'} exact component={Home}/>
                                <Route path={'/login'} component={Login}/>
                            </Switch>
                            <TaskModal/>
                        </Content>
                    </Col>
                </Row>
                <Footer>Это футер</Footer>

            </Layout>

        </BrowserRouter>

    )
}