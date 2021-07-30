import {ADD_TASK, DELETE_TASK, EDIT_TASK, FETCH_TASKS, SHOW_LOADER} from "../types";

const handlers = {
    [SHOW_LOADER]: state => ({...state, loading: true}),
    [ADD_TASK]: (state, {payload}) => ({
        ...state,
        tasks: [...state.tasks, payload]
    }),
    [FETCH_TASKS]: (state, {payload}) => ({
        ...state,
        tasks: payload,
        loading: false
    }),
    [DELETE_TASK]: (state, {payload}) => ({
        ...state,
        tasks: state.tasks.filter(task => task.id !== payload)
    }),
    [EDIT_TASK]: (state, {payload}) => ({
        ...state,
        tasks: state.tasks.map((task) => task.id === payload.id ? payload : task)
    }),
    DEFAULT: state => state
}

export const tasksReducer = (state, action) => {
    const handle = handlers[action.type] || handlers.DEFAULT
    return handle(state, action)
}