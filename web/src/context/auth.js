import { useState, createContext, useContext } from 'react';
import PropTypes from 'prop-types';
import axios from 'axios';

const apiPrefix = process.env.REACT_APP_API_PREFIX;

const AuthContext = createContext();

const useAuth = () => {
  return useContext(AuthContext);
};

const AuthProvider = (props) => {
  const [token, setToken] = useState(null);

  const getToken = () => {
    const token = localStorage.getItem('token');
    setToken(token);
  };

  const register = async (email, password) => {
    const url = apiPrefix + '/v1/register';
    const payload = {
      email,
      password,
    };

    try {
      const res = await axios.post(url, payload);
      const { token } = res.data;

      localStorage.setItem('token', token);
      setToken(token);
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

      return true;
    } catch (err) {
      localStorage.removeItem('token');
      setToken(null);

      return false;
    }
  };

  const login = async (email, password) => {
    const url = apiPrefix + '/v1/login';
    const payload = {
      email,
      password,
    };

    try {
      const res = await axios.post(url, payload);
      const { token } = res.data;

      setToken(token);
      localStorage.setItem('token', token);
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

      return true;
    } catch (err) {
      localStorage.removeItem('token');
      setToken(null);
      delete axios.defaults.headers.common['Authorization'];

      return false;
    }
  };

  const logout = () => {
    localStorage.removeItem('token');
    setToken(null);
  };

  const auth = {
    token,
    getToken,
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

export { AuthContext, AuthProvider, useAuth };
