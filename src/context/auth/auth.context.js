import { createContext, useContext, useReducer } from 'react';
import config from '../../config';
import { initialState, reducer } from './auth.reducer';
import * as actionTypes from './auth.types';

const AuthContext = createContext();

export const useAuthContext = () => useContext(AuthContext);

export const AuthProvider = (props) => {
  const [state, dispatch] = useReducer(reducer, initialState);

  const register = (email, password) => {
    dispatch({ type: actionTypes.AUTHENTICATING });

    const url = config.prefix + '/api/v1/register';

    return new Promise((resolve, reject) => {
      fetch(url, {
        method: 'POST',
        body: JSON.stringify({ email, password }),
      })
        .then((resp) => resp.json())
        .then((data) => {
          if (data.success) {
            dispatch({
              type: actionTypes.AUTH_SUCCESS,
              id: data.payload.id,
              token: data.payload.token,
            });
          } else {
            dispatch({
              type: actionTypes.AUTH_FAIL,
              error: data.message,
            });
          }

          return resolve(data.success);
        })
        .catch(() => {
          // should only reach here if there is a network or browser related error
          dispatch({
            type: actionTypes.AUTH_FAIL,
            error: 'Something went wrong. Pleade try again later.',
          });

          return reject(false);
        });
    });
  };

  const login = (email, password) => {
    dispatch({ type: actionTypes.AUTHENTICATING });

    const url = config.prefix + '/api/v1/login';

    return new Promise((resolve, reject) => {
      fetch(url, {
        method: 'POST',
        body: JSON.stringify({ email, password }),
      })
        .then((resp) => resp.json())
        .then((data) => {
          if (data.success) {
            dispatch({
              type: actionTypes.AUTH_SUCCESS,
              id: data.payload.id,
              token: data.payload.token,
            });
          } else {
            dispatch({
              type: actionTypes.AUTH_FAIL,
              error: data.message,
            });
          }

          return resolve(data.success);
        })
        .catch(() => {
          // should only reach here if there is a network or browser related error
          dispatch({
            type: actionTypes.AUTH_FAIL,
            error: 'Something went wrong. Pleade try again later.',
          });

          return reject(false);
        });
    });
  };

  const logout = () => {
    dispatch({ type: actionTypes.AUTH_LOGOUT });
  };

  const auth = {
    isLoggedIn: state.isLoggedIn,
    id: state.id,
    token: state.token,
    authenticating: state.authenticating,
    error: state.error,
    register,
    login,
    logout,
  };

  return (
    <AuthContext.Provider value={auth}>{props.children}</AuthContext.Provider>
  );
};
