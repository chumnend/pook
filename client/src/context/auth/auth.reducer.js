import { AUTH_SUCCESS, AUTH_ERROR, LOGOUT } from './auth.types';

const initialState = {
  isLoggedIn: false,
  user: {},
};

const authReducer = (state = initialState, action) => {
  switch (action.type) {
    case AUTH_SUCCESS:
      return {
        isLoggedIn: true,
        user: action.user,
      };
    case AUTH_ERROR:
      return {
        isLoggedIn: false,
        user: {},
      };
    case LOGOUT:
      return {
        isLoggedIn: false,
        user: {},
      };
    default:
      return state;
  }
};

export { initialState, authReducer };
