import React, { useState } from 'react';

import AuthContext from '../context/AuthContext';
import authService from '../services/auth';
import type UserType from '../types/UserType';
import toSentenceCase from '../utils/toSentenceCase';

type Props = {
  children: React.ReactNode;
}

const AuthProvider = ({ children }: Props) => {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [user, setUser] = useState<UserType | null>(null);
  const [authError, setAuthError] = useState<string | null>(null);

  const register = async (email: string, username: string, password: string): Promise<boolean> => {
    try {
      await authService.register(email, username, password);
      return false;
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
      setUser({ id: data.id, email: data.email, username: data.username, token: data.token });
      return true;
    } catch (error) {
      if (error instanceof Error) {
        setAuthError(toSentenceCase(error.message));
      } else {
        setAuthError('An unknown error occurred');
      }
      setIsLoggedIn(false);
      setUser(null);
      return false;
    }
  }

  const logout = () => {
    setIsLoggedIn(false);
    setUser(null);
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
