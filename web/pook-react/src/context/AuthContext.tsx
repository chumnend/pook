import { createContext } from 'react';

export type AuthContextState = {
  user: { email: string; username: string } | null;
  login: (data: { email: string, username: string }) => void;
  logout: () => void;
};

const AuthContext = createContext<AuthContextState | undefined>(undefined);

export default AuthContext;
