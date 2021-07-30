import {Alert} from "antd";
import React from "react";

export const TaskAlert = ({message, type}) =>{
    return (
        <Alert
            message="Задача успешно создана"
            type="success"
            showIcon
        />
    )
}
