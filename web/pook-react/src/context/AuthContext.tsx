import { createContext } from 'react';

import type UserType from '../types/UserType';

export type AuthContextState = {
  user: UserType | null;
  register: (email: string, username: string, password: string) => void;
  login: (username: string, password: string) => void;
  logout: () => void;
};

const AuthContext = createContext<AuthContextState | undefined>(undefined);

export default AuthContext;
