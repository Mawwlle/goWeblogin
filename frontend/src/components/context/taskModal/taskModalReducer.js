import {HIDE_MODAL, SHOW_MODAL} from '../types'

const handlers = {
  [SHOW_MODAL]: (state, {payload}) => ({...payload, visible: true}),
  [HIDE_MODAL]: state => ({...state, visible: false}),
  DEFAULT: state => state
}

export const taskModalReducer = (state, action) => {
  const handle = handlers[action.type] || handlers.DEFAULT
  return handle(state, action)
}
