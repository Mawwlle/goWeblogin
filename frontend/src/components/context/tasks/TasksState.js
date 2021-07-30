import React, {useReducer} from "react";
import {TasksContext} from "./tasksContext";
import {tasksReducer} from "./tasksReducer";
import axios from "axios";
import {ADD_TASK, DELETE_TASK, EDIT_TASK, FETCH_TASKS, SHOW_LOADER} from "../types";

const url = process.env.REACT_APP_DB_URL

export const TasksState = ({children}) => {
    const initialState = {
        tasks: [],
        loading: false
    }
    const [state, dispatch] = useReducer(tasksReducer, initialState)

    const showLoader = () => dispatch({type: SHOW_LOADER})

    const fetchTasks = async () => {
        showLoader()
        const res = await axios.get(`${url}/tasks.json`)
        const payload = Object.keys(res.data || []).map(key => {
            return {
                ...res.data[key],
                id: key
            }
        })
        dispatch({
            type: FETCH_TASKS,
            payload
        })
    }

    const addTask = async data => {
        const task = {
            title: data.title,
            description: data.description,
            priority: data.priority,
            date: new Date().toJSON()
        }
        try {
            const res = await axios.post(`${url}/tasks.json`, task)
            const payload = {
                ...task,
                id: res.data.name
            }
            dispatch({
                type: ADD_TASK,
                payload
            })
        } catch (e) {
            throw new Error(e.message)
        }

    }

    const deleteTask = async id => {
        await axios.delete(`${url}/tasks/${id}.json`)
        dispatch({
            type: DELETE_TASK,
            payload: id
        })
    }

    const editTask = async data => {
        const task = {
            title: data.title,
            description: data.description,
            priority: data.priority,
            date: new Date().toJSON(),
            id: data.id
        }
        try {
            const res = await axios.put(`${url}/tasks/${task.id}.json`, task)
            const payload = {
                ...task,
            }
            dispatch({
                type: EDIT_TASK,
                payload
            })
        } catch (e) {
            throw new Error(e.message)
        }
    }
    return (
        <TasksContext.Provider value={{
            showLoader, addTask, deleteTask, fetchTasks, editTask,
            loading: state.loading,
            tasks: state.tasks
        }}>
            {children}
        </TasksContext.Provider>
    )
}