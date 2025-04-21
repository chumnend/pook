import React, { useState } from 'react';

import AuthContext from '../context/AuthContext';

type Props = {
  children: React.ReactNode;
}

const AuthProvider = ({ children }: Props) => {
  const [user, setUser] = useState<{ email: string; username: string } | null>(null);

  const login = (data: { email: string; username: string}) => {
    setUser(data);
  }

  const logout = () => {
    setUser(null);
  }

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider;
