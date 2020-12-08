import { createContext, useContext, useState } from 'react';
import config from '../config';

const AuthContext = createContext();

export const useAuthContext = () => useContext(AuthContext);

const AuthProvider = (props) => {
  const [isLoggedIn, setLoggedIn] = useState(false);
  const [token, setToken] = useState(null);

  const register = (email, password) => {
    const url = config.prefix + '/api/v1/register';

    return new Promise((resolve, reject) => {
      fetch(url, {
        method: 'POST',
        body: JSON.stringify({ email, password }),
      })
        .then((resp) => resp.json())
        .then((data) => resolve(data))
        .catch((err) => reject(err));
    });
  };

  const login = (email, password) => {
    const url = config.prefix + '/api/v1/login';

    return new Promise((resolve, reject) => {
      fetch(url, {
        method: 'POST',
        body: JSON.stringify({ email, password }),
      })
        .then((resp) => resp.json())
        .then((data) => resolve(data))
        .catch((err) => reject(err));
    });
  };

  const logout = () => {
    alert('logging out');
  };

  const auth = {
    isLoggedIn,
    token,
    register,
    login,
    logout,
  };

  return (
    <AuthContext.Provider value={auth}>{props.children}</AuthContext.Provider>
  );
};

export default AuthProvider;
