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

/** local storage key */
export const AUTH_STATE_KEY = 'authState';

/*
 * Registers a new user.
 * @param {string} email - the user's email
 * @param {string} password - the user's password
 * @return {Object} the users information
 */
export const register = async (firstName, lastName, email, password) => {
  try {
    const res = await axios.post(API_USER_REGISTER, {
      firstName: firstName,
      lastName: lastName,
      email: email,
      password: password,
    });
    const { token } = res.data;

    const decoded = jwtDecode(token);
    const authState = {
      id: decoded.id,
      email: decoded.email,
      token: token,
    };

    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    return authState;
  } catch (error) {
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

    const decoded = jwtDecode(token);
    const authState = {
      id: decoded.id,
      email: decoded.email,
      token: token,
    };

    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    return authState;
  } catch (error) {
    throw error;
  }
};

/*
 * Checks for any saved authentication information
 * @return {Object} saved user information ie. { id: '12122', email: 'user@example.com', token: 'sdadsdasd2e2e1' }
 */
export const checkAuthState = () => {
  const authState = localStorage.getItem(AUTH_STATE_KEY);
  if (authState) {
    axios.defaults.headers.common[
      'Authorization'
    ] = `Bearer ${authState.token}`;
    return JSON.parse(authState);
  }

  return null;
};

/*
 * Saves authentication information
 * @param {string} id - the user's identifier
 * @param {string} email - the user's email
 * @param {string} token - the user's token
 */
export const saveAuthState = (id, email, token) => {
  if (!id || !email || !token) {
    return null;
  }

  localStorage.setItem(AUTH_STATE_KEY, JSON.stringify({ id, email, token }));
};

/*
 * Erases all saved authentication information
 */
export const clearAuthState = () => {
  localStorage.removeItem(AUTH_STATE_KEY);
  delete axios.defaults.headers.common['Authorization'];
};
