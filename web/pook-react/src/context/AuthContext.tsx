import { createContext } from 'react';

import type UserType from '../types/UserType';

export type AuthContextState = {
  user: UserType | null;
  authError: string | null;
  register: (email: string, username: string, password: string) => Promise<boolean>;
  login: (username: string, password: string) => Promise<boolean>;
  logout: () => void;
};

const AuthContext = createContext<AuthContextState | undefined>(undefined);

export default AuthContext;
