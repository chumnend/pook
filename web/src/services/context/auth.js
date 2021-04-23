import axios from 'axios';
import jwt_decode from 'jwt-decode';
import PropTypes from 'prop-types';
import { createContext, useState } from 'react';

// By default the react app runs on the same server as the api.
// In development mode, a different server can be pointed to using REACT_APP_API_PREFIX
let apiPrefix = '';
if (process.env.NODE_ENV === 'development') {
  apiPrefix = process.env.REACT_APP_API_PREFIX;
}

const AuthContext = createContext();

const AuthProvider = (props) => {
  const [isLoggedIn, setLoggedIn] = useState(false);
  const [user, setUser] = useState(null);
  const [error, setError] = useState(null);

  const getToken = () => {
    const token = localStorage.getItem('token');
    if (token != null) {
      decodeToken(token);
    }
  };

  const decodeToken = (token) => {
    const decoded = jwt_decode(token);
    const user = {
      id: decoded.ID,
      email: decoded.Email,
    };

    setLoggedIn(true);
    setUser(user);
  };

  const parseError = (error) => {
    if (error.response !== undefined) {
      return error.response.data.error;
    } else {
      return 'Something went wrong. Please try again later.';
    }
  };

  const clearError = () => {
    setError(null);
  };

  const register = async (email, password) => {
    const url = apiPrefix + '/api/v1/register';
    const payload = {
      email,
      password,
    };

    setError(null);

    try {
      const res = await axios.post(url, payload);
      const { token } = res.data;

      decodeToken(token);
      localStorage.setItem('token', token);
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

      return true;
    } catch (err) {
      setUser(null);
      setError(parseError(err));
      localStorage.removeItem('token');
      delete axios.defaults.headers.common['Authorization'];

      return false;
    }
  };

  const login = async (email, password) => {
    const url = apiPrefix + '/api/v1/login';
    const payload = {
      email,
      password,
    };

    setError(null);

    try {
      const res = await axios.post(url, payload);
      const { token } = res.data;

      decodeToken(token);
      localStorage.setItem('token', token);
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

      return true;
    } catch (err) {
      setLoggedIn(false);
      setUser(null);
      setError(parseError(err));
      localStorage.removeItem('token');
      delete axios.defaults.headers.common['Authorization'];

      return false;
    }
  };

  const logout = () => {
    setLoggedIn(false);
    setUser(null);
    setError(null);
    localStorage.removeItem('token');
    delete axios.defaults.headers.common['Authorization'];
  };

  const auth = {
    isLoggedIn,
    user,
    error,
    getToken,
    clearError,
    register,
    login,
    logout,
  };

  return (
    <AuthContext.Provider value={auth}>{props.children}</AuthContext.Provider>
  );
};

AuthProvider.propTypes = {
  children: PropTypes.node,
};

export { AuthContext, AuthProvider };
