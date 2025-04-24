import React, { useState } from 'react';

import AuthContext from '../context/AuthContext';
import authService from '../services/auth';
import type UserType from '../types/UserType';
import toSentenceCase from '../utils/toSentenceCase';

type Props = {
  children: React.ReactNode;
}

const AuthProvider = ({ children }: Props) => {
  const [user, setUser] = useState<UserType | null>(null);
  const [authError, setAuthError] = useState<string | null>(null);

  const clearAuthError = () => {
    setAuthError(null);
  }

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
      setUser({ id: data.id, email: data.email, username: data.username, token: data.token });
      return true;
    } catch (error) {
      if (error instanceof Error) {
        setAuthError(toSentenceCase(error.message));
      } else {
        setAuthError('An unknown error occurred');
      }
      setUser(null);
      return false;
    }
  }

  const logout = () => {
    setUser(null);
  }

  return (
    <AuthContext.Provider value={{ user, authError, clearAuthError, register, login, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider;
