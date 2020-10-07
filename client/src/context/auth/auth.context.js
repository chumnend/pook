import React, { createContext, useReducer } from 'react';
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import { initialState, authReducer } from './auth.reducer';
import { AUTH_SUCCESS, AUTH_ERROR, LOGOUT } from './auth.types';
import config from '../../config';

axios.defaults.withCredentials = true;

const AuthContext = createContext();
const AuthProvider = (props) => {
  const [authState, dispatch] = useReducer(authReducer, initialState);

  // sign in the user
  const authorizeUser = async (authType, payload) => {
    try {
      const res = await axios.post(
        `${config.url}/api/users/${authType}`,
        payload,
      );

      // decode the token and set in store
      const token = res.data.token;
      const user = jwt_decode(token);
      dispatch({
        type: AUTH_SUCCESS,
        user,
      });

      return;
    } catch (err) {
      dispatch({ type: AUTH_ERROR });
      throw err;
    }
  };

  const setUser = async () => {
    try {
      const res = await axios.post(`${config.url}/api/users/validate`);
      const { isValid, user } = res.data;

      if (isValid) {
        dispatch({
          type: AUTH_SUCCESS,
          user,
        });
      }
    } catch (err) {
      console.error(err);
    }
  };

  // log out a user
  const logout = async () => {
    try {
      await axios.post(`${config.url}/api/users/logout`);
      dispatch({ type: LOGOUT });
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <AuthContext.Provider value={{ authState, authorizeUser, setUser, logout }}>
      {props.children}
    </AuthContext.Provider>
  );
};

export { AuthContext, AuthProvider };
