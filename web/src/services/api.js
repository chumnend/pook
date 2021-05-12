import axios from 'axios';
import jwt_decode from 'jwt-decode';

// By default the react app runs on the same server as the api.
// In development mode, a different server can be pointed to using REACT_APP_API_PREFIX
let apiPrefix = '';
if (process.env.NODE_ENV === 'development') {
  apiPrefix = process.env.REACT_APP_API_PREFIX;
}

/*
 * Registers a new user.
 * @param {string} email - the user's email
 * @param {string} password - the user's password
 * @return {Object} the users information
 */
export const register = async (email, password) => {
  const url = apiPrefix + '/api/v1/register';
  const payload = { email, password };

  try {
    const res = await axios.post(url, payload);
    const { token } = res.data;

    localStorage.setItem('token', token);
    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    const decoded = jwt_decode(token);
    const user = {
      id: decoded.ID,
      email: decoded.Email,
    };

    return user;
  } catch (error) {
    localStorage.removeItem('token');
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
  const url = apiPrefix + '/api/v1/login';
  const payload = { email, password };

  try {
    const res = await axios.post(url, payload);
    const { token } = res.data;

    localStorage.setItem('token', token);
    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    const decoded = jwt_decode(token);
    const user = {
      id: decoded.ID,
      email: decoded.Email,
    };

    return user;
  } catch (error) {
    localStorage.removeItem('token');
    delete axios.defaults.headers.common['Authorization'];
    throw error;
  }
};

/*
 * Log user out of browser
 */
export const logout = () => {
  localStorage.removeItem('token');
  delete axios.defaults.headers.common['Authorization'];
};
