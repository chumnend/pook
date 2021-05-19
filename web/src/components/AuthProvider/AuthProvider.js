import PropTypes from 'prop-types';
import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useState,
} from 'react';

import * as apiHelpers from '../../services/api';
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
    const authState = apiHelpers.checkAuthState();
    setState((state) => ({
      ...state,
      loading: false,
      user: authState,
    }));
  }, []);

  const register = useCallback(async (fname, lname, email, password) => {
    try {
      const user = await apiHelpers.register(fname, lname, email, password);
      setState((state) => ({
        ...state,
        user: user,
      }));

      return true;
    } catch (error) {
      let errorMessage;
      if (error.response) {
        errorMessage = error.response.data.error;
      } else {
        errorMessage = error.message;
      }

      setState((state) => ({
        ...state,
        error: errorMessage,
        user: null,
      }));

      return false;
    }
  }, []);

  const login = useCallback(async (email, password, saveUser = false) => {
    try {
      const user = await apiHelpers.login(email, password);
      setState((state) => ({
        ...state,
        user: user,
      }));

      if (saveUser) {
        apiHelpers.saveAuthState(user.id, user.email, user.token);
      }

      return true;
    } catch (error) {
      let errorMessage;
      if (error.response) {
        errorMessage = error.response.data.error;
      } else {
        errorMessage = error.message;
      }

      setState((state) => ({
        ...state,
        error: errorMessage,
        user: null,
      }));

      return false;
    }
  }, []);

  const logout = useCallback(() => {
    apiHelpers.clearAuthState();
    setState((state) => ({
      ...state,
      error: null,
      user: null,
    }));
  }, []);

  const values = {
    ...state,
    isAuth: state.user !== null,

    register,
    login,
    logout,
  };

  if (state.initializing) {
    return <Loader fullPage />;
  }

  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>;
};

AuthProvider.propTypes = {
  children: PropTypes.node.isRequired,
};

export default AuthProvider;
