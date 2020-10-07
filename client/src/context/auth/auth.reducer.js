import { AUTH_SUCCESS, AUTH_ERROR, LOGOUT, SET_USER } from './auth.types';

const initialState = {
  isLoggedIn: false,
  user: {},
  token: localStorage.getItem('token') || '',
};

const authReducer = (state = initialState, action) => {
  switch (action.type) {
    case AUTH_SUCCESS:
      return {
        isLoggedIn: true,
        user: action.user,
        token: action.token,
      };
    case AUTH_ERROR:
      return {
        isLoggedIn: false,
        user: {},
        token: '',
      };
    case LOGOUT:
      return {
        isLoggedIn: false,
        user: {},
        token: '',
      };
    case SET_USER:
      return {
        isLoggedIn: true,
        user: action.user,
        token: action.token,
      };
    default:
      return state;
  }
};

export { initialState, authReducer };
