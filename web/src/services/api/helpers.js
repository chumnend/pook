import axios from 'axios';
import jwtDecode from 'jwt-decode';

import { API_USER_LOGIN, API_USER_REGISTER } from './constants';

/*
 * Registers a new user.
 * @param {string} email - the user's email
 * @param {string} password - the user's password
 * @return {Object} the users information
 */
export const register = async (email, password) => {
  try {
    const res = await axios.post(API_USER_REGISTER, { email, password });
    const { token } = res.data;

    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    const decoded = jwtDecode(token);
    const user = {
      id: decoded.ID,
      email: decoded.Email,
      token: token,
    };

    return user;
  } catch (error) {
    delete axios.defaults.headers.common['Authorization'];
    throw error;
  }
};

/*
 * Login a new user.
 * @param {string} email - the user's email
 * @param {string} password - the user's password
 * @return {Object} the users information
 */
export const login = async (email, password) => {
  try {
    const res = await axios.post(API_USER_LOGIN, { email, password });
    const { token } = res.data;

    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    const decoded = jwtDecode(token);
    const user = {
      id: decoded.ID,
      email: decoded.Email,
      token: token,
    };

    return user;
  } catch (error) {
    delete axios.defaults.headers.common['Authorization'];
    throw error;
  }
};

/*
 * Log user out of browser
 */
export const logout = () => {
  delete axios.defaults.headers.common['Authorization'];
};
