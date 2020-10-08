import React, { useReducer } from 'react';
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import { initialState, authReducer } from './auth.reducer';
import { AUTH_SUCCESS, AUTH_ERROR, LOGOUT } from './auth.types';
import config from '../../config';

axios.defaults.withCredentials = true;

const AuthContext = React.createContext();

const AuthProvider = (props) => {
  const [state, dispatch] = useReducer(authReducer, initialState);

  const context = {
    isLoggedIn: state.isLoggedIn,
    user: state.user,

    async authorizeUser(authType, payload) {
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
    },

    async setUser() {
      try {
        const res = await axios.post(`${config.url}/api/users/validate`);
        const { success, user } = res.data;

        if (success) {
          dispatch({
            type: AUTH_SUCCESS,
            user,
          });
        }
      } catch (err) {
        dispatch({ type: AUTH_ERROR });
        console.error(err);
      }
    },

    async logout() {
      try {
        await axios.post(`${config.url}/api/users/logout`);
        dispatch({ type: LOGOUT });
      } catch (err) {
        dispatch({ type: AUTH_ERROR });
        console.error(err);
      }
    },
  };

  return (
    <AuthContext.Provider value={context}>
      {props.children}
    </AuthContext.Provider>
  );
};

export { AuthContext, AuthProvider };
