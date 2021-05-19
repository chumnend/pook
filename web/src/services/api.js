import axios from 'axios';
import jwtDecode from 'jwt-decode';

// By default the react app runs on the same server as the api.
// In development mode, a different server can be pointed to using REACT_APP_API_PREFIX
let apiPrefix = '';
/* istanbul ignore next */
if (process.env.NODE_ENV === 'development') {
  apiPrefix = process.env.REACT_APP_API_PREFIX;
}

/** user api routes */
export const API_USER_REGISTER = apiPrefix + '/api/v1/register';
export const API_USER_LOGIN = apiPrefix + '/api/v1/login';

/*
 * Registers a new user.
 * @param {string} email - the user's email
 * @param {string} password - the user's password
 * @return {Object} the users information
 */
export const register = async (firstName, lastName, email, password) => {
  try {
    const res = await axios.post(API_USER_REGISTER, {
      firstname: firstName,
      lastname: lastName,
      email: email,
      password: password,
    });
    const { token } = res.data;

    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    const decoded = jwtDecode(token);
    const user = {
      id: decoded.id,
      email: decoded.email,
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
      id: decoded.id,
      email: decoded.email,
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
