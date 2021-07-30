import React, {useContext, useEffect} from 'react';
import {Button} from 'antd';
import {Tasks} from "../components/Tasks";
import {TasksContext} from "../components/context/tasks/tasksContext";
import {Loader} from "../components/Loader";
import {TaskModalContext} from "../components/context/taskModal/taskModalContext";

export const Home = () => {
    const {loading, tasks, fetchTasks} = useContext(TasksContext)
    const {show} = useContext(TaskModalContext)
    useEffect(() => {
        fetchTasks()
    }, [])


    return (
        <div>
            <Button
                type="primary"
                onClick={() => {show('create')}}
            >
                Добавить задачу
            </Button>
            {loading
                ? <Loader/>
                : <Tasks tasks={tasks}/>
            }
        </div>
    );
}