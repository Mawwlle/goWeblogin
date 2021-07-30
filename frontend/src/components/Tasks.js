import React, {useContext} from "react";
import {Card, Row, Col, Typography, Avatar} from "antd";
import {DeleteOutlined, EditOutlined} from "@ant-design/icons"
import {TasksContext} from "./context/tasks/tasksContext";
import {TaskModalContext} from "./context/taskModal/taskModalContext";
import {TransitionGroup, CSSTransition} from "react-transition-group";

const {Text, Paragraph} = Typography
export const Tasks = ({tasks}) => {
    const {deleteTask} = useContext(TasksContext)
    const {show} = useContext(TaskModalContext)

    const edit = task => {
        show('edit', task)
    }
    return (
        <Row gutter={[32, 32]} className="tasks">
            {
                tasks.map(task => {
                        const a = <Avatar className="task__priority">{task.priority}</Avatar>
                        return (
                            <Col span={8} key={task.id}>
                                <Card
                                    className="task"
                                    hoverable
                                    title={task.title}
                                    extra={a}
                                    actions={[
                                        <DeleteOutlined
                                            key="remove"
                                            onClick={() => deleteTask(task.id)}
                                        />,
                                        <EditOutlined
                                            key="edit"
                                            onClick={()=>show('edit', task)}
                                        />]}
                                >
                                    <Paragraph>{task.description}</Paragraph>
                                    <Paragraph type="secondary"><Text strong type="secondary">Дата
                                        создания:</Text> {task.date}
                                    </Paragraph>

                                </Card>
                            </Col>
                        )
                    }
                )
            }
        </Row>
    )
}