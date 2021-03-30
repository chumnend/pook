import { useState, createContext, useContext } from 'react';
import PropTypes from 'prop-types';
import axios from 'axios';
import jwt_decode from 'jwt-decode';

const apiPrefix = process.env.REACT_APP_API_PREFIX;

const AuthContext = createContext();

const useAuth = () => {
  return useContext(AuthContext);
};

const AuthProvider = (props) => {
  const [user, setUser] = useState(null);

  const decodeToken = (token) => {
    const decoded = jwt_decode(token);
    const user = {
      id: decoded.ID,
      email: decoded.Email,
    };

    console.log(decoded);

    setUser(user);
  };

  const getToken = () => {
    const token = localStorage.getItem('token');
    if (token != null) {
      decodeToken(token);
    }
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

      decodeToken(token);
      localStorage.setItem('token', token);
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

      return true;
    } catch (err) {
      setUser(null);
      localStorage.removeItem('token');
      delete axios.defaults.headers.common['Authorization'];

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

      decodeToken(token);
      localStorage.setItem('token', token);
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

      return true;
    } catch (err) {
      setUser(null);
      localStorage.removeItem('token');
      delete axios.defaults.headers.common['Authorization'];

      return false;
    }
  };

  const logout = () => {
    localStorage.removeItem('token');
    setUser(null);
  };

  const auth = {
    user,
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
