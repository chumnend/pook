import React, { createContext, useReducer } from 'react';
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import { initialState, authReducer } from './auth.reducer';
import { AUTH_SUCCESS, AUTH_ERROR, LOGOUT, SET_USER } from './auth.types';
import config from '../../config';

const AuthContext = createContext();

function AuthProvider(props) {
  const [authState, dispatch] = useReducer(authReducer, initialState);

  // sign in the user
  const authorizeUser = async (authType, payload) => {
    try {
      const res = await axios.post(
        `${config.uri}/api/users/${authType}`,
        payload,
      );
      const token = res.data.token;

      // save the token
      localStorage.setItem('token', token);
      axios.defaults.headers.common['Authorization'] = token;

      // decode the token and set in store
      const user = jwt_decode(token);
      dispatch({
        type: AUTH_SUCCESS,
        user,
        token,
      });

      return;
    } catch (err) {
      localStorage.removeItem('token');
      dispatch({ type: AUTH_ERROR });
      throw err;
    }
  };

  const setUser = (token) => {
    const user = jwt_decode(token);
    dispatch({
      type: SET_USER,
      user,
      token,
    });
  };

  // log out a user
  const logout = () => {
    localStorage.removeItem('token');
    delete axios.defaults.headers.common['Authorization'];
    dispatch({ type: LOGOUT });
  };

  return (
    <AuthContext.Provider value={{ authState, authorizeUser, setUser, logout }}>
      {props.children}
    </AuthContext.Provider>
  );
}

export { AuthContext, AuthProvider };
