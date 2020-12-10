import * as actionTypes from './auth.types';

export const initialState = {
  isLoggedIn: false,
  id: null,
  token: null,
  authenticating: false,
  error: null,
};

export const reducer = (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.AUTHENTICATING:
      return {
        ...state,
        authenticating: true,
        error: null,
      };
    case actionTypes.AUTH_SUCCESS:
      return {
        ...state,
        isLoggedIn: true,
        id: action.id,
        token: action.token,
        authenticating: false,
        error: null,
      };
    case actionTypes.AUTH_FAIL:
      return {
        ...state,
        isLoggedIn: false,
        id: null,
        token: null,
        authenticating: false,
        error: action.error,
      };
    case actionTypes.AUTH_LOGOUT:
      return initialState;
    default:
      return state;
  }
};
