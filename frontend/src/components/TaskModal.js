import React, {useContext} from 'react';
import {Modal, Form, Input, Radio} from 'antd';
import {TasksContext} from "./context/tasks/tasksContext";
import {TaskModalContext} from "./context/taskModal/taskModalContext";

export const TaskModal = () => {
    const [form] = Form.useForm();
    const modalInfo = {
        title: "Добавить новую задачу",
        action: "Добавить"
    }
    const {addTask, editTask} = useContext(TasksContext)
    const {modal, hide} = useContext(TaskModalContext)

    if (modal.action === 'edit') {
        modalInfo.title = "Изменить задачу"
        modalInfo.action = "Изменить"
    }
    const onOk = values => {
        console.log('Received values of form: ', values);
        if (modal.action === 'edit') {
            values.id = modal.task.id
            editTask(values)
        } else {
            addTask(values).then(() => {
                console.log("Success")
            }).catch(() => {
                console.log("warning")
            })
        }
        hide()
    }

    return (
        <Modal
            visible={modal.visible}
            title={modalInfo.title}
            okText={modalInfo.action}
            cancelText="Отменить"
            onCancel={() => hide()}
            onOk={() => {
                form
                    .validateFields()
                    .then((values) => {
                        form.resetFields();
                        onOk(values);
                    })
                    .catch((info) => {
                        console.log('Validate Failed:', info);
                    });
            }}
        >
            <Form
                form={form}
                layout="vertical"
                name="form_in_modal"
                initialValues={{}}
            >
                <Form.Item
                    name="title"
                    label="Название"
                    rules={[
                        {
                            required: true,
                            message: 'Пожалуйста, введите название задачи',
                        },
                    ]}
                >
                    <Input/>
                </Form.Item>
                <Form.Item
                    name="description"
                    label="Описание"
                >
                    <Input type="textarea"/>
                </Form.Item>
                <Form.Item name="priority" className="collection-create-form_last-form-item" label="Приоритет">
                    <Radio.Group>
                        {
                            [1, 2, 3].map(value => (
                                <Radio.Button value={value}
                                              checked={
                                                  modal.task?.priority
                                                  ? modal.task.priority === value
                                                  : false}
                                >
                                    {value}
                                </Radio.Button>
                            ))
                        }

                    </Radio.Group>
                </Form.Item>
            </Form>
        </Modal>
    );
}