import React, {useReducer} from 'react'
import {TaskModalContext} from './taskModalContext'
import {taskModalReducer} from './taskModalReducer'
import {HIDE_MODAL, SHOW_MODAL} from '../types'

export const TaskModalState = ({children}) => {
  const [state, dispatch] = useReducer(taskModalReducer, {visible: false})

  const show = (action, task) => {
    dispatch({
      type: SHOW_MODAL,
      payload: {action, task}
    })
  }

  const hide = () => dispatch({type: HIDE_MODAL})

  return (
    <TaskModalContext.Provider value={{
      show, hide,
      modal: state
    }}>
      {children}
    </TaskModalContext.Provider>
  )
}
