import PropTypes from 'prop-types';
import { createContext, useContext, useEffect, useState } from 'react';

import Loader from '../Loader';

export const AuthContext = createContext();

export const useAuth = () => useContext(AuthContext);

const AuthProvider = ({ children }) => {
  const [state, setState] = useState({
    loading: true,
    error: null,
    user: null,
  });

  useEffect(() => {
    console.log('check localstorage?');
    setState((state) => ({ ...state, loading: false }));
  }, []);

  if (state.loading) {
    return <Loader fullPage />;
  }

  // check if user information in localStorage
  // return loader until data is collected

  const login = () => {};
  const register = () => {};
  const logout = () => {};

  const values = {
    login,
    register,
    logout,
  };

  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>;
};

AuthProvider.propTypes = {
  children: PropTypes.node.isRequired,
};

export default AuthProvider;
