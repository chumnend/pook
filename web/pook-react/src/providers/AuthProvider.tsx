import React, { useState } from 'react';

import AuthContext from '../context/AuthContext';
import authService from '../services/auth';
import type UserType from '../types/UserType';

type Props = {
  children: React.ReactNode;
}

const AuthProvider = ({ children }: Props) => {
  const [user, setUser] = useState<UserType | null>(null);

  const register = async (email: string, username: string, password: string) => {
    try {
      const data = await authService.register(email, username, password);
      console.log('Registration successful:', data);
    } catch (error) {
      console.error('Registration error:', error);
    }
    setUser(null);
  }

  const login = async (username: string, password: string) => { 
    try {
      const data = await authService.login(username, password);
      console.log('Login successful:', data);
    } catch (error) {
      console.error('Login error:', error);
    }
    setUser(null);
  }

  const logout = () => {
    setUser(null);
  }

  return (
    <AuthContext.Provider value={{ user, register, login, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider;
