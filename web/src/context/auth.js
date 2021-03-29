import { useState, createContext, useContext } from 'react';
import PropTypes from 'prop-types';

// const apiPrefix = process.env.REACT_APP_API_PREFIX;

const AuthContext = createContext();

const useAuth = () => {
  return useContext(AuthContext);
};

const AuthProvider = (props) => {
  const [token, setToken] = useState(null);

  const getToken = () => {
    const token = sessionStorage.getItem('token');
    setToken(token);
  };

  const register = () => {};
  const login = () => {};
  const logout = () => {};

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
