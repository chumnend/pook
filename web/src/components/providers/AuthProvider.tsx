import React, { useState, useEffect } from 'react';

import AuthContext from '../../helpers/context/AuthContext';
import authService from '../../helpers/services/user';
import type { User as UserType } from '../../helpers/types';
import { toSentenceCase } from '../../helpers/utils';

type Props = {
  children: React.ReactNode;
}

const AuthProvider = ({ children }: Props) => {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [user, setUser] = useState<UserType | null>(null);
  const [authError, setAuthError] = useState<string | null>(null);

  useEffect(() => {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      const parsedUser = JSON.parse(storedUser);
      setUser(parsedUser);
      setIsLoggedIn(true);
    }
  }, []);

  const register = async (email: string, username: string, password: string): Promise<boolean> => {
    try {
      await authService.register(email, username, password);
      return true;
    } catch (error) {
      if (error instanceof Error) {
        setAuthError(toSentenceCase(error.message));
      } else {
        setAuthError('An unknown error occurred');
      }
      return false;
    }
  }

  const login = async (username: string, password: string): Promise<boolean> => { 
    try {
      const data = await authService.login(username, password);
      setIsLoggedIn(true);
      const userData = { id: data.id, email: data.email, username: data.username, token: data.token };
      setUser(userData);
      localStorage.setItem('user', JSON.stringify(userData));
      return true;
    } catch (error) {
      if (error instanceof Error) {
        setAuthError(toSentenceCase(error.message));
      } else {
        setAuthError('An unknown error occurred');
      }
      setIsLoggedIn(false);
      setUser(null);
      localStorage.removeItem('user');
      return false;
    }
  }

  const logout = () => {
    setIsLoggedIn(false);
    setUser(null);
    localStorage.removeItem('user');
  }

  const clearAuthError = () => {
    setAuthError(null);
  }

  return (
    <AuthContext.Provider value={{
      isLoggedIn,
      user,
      authError,
      register,
      login,
      logout,
      clearAuthError,
    }}>
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider;
