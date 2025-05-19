import { createContext } from 'react';

import type UserType from '../types/UserType';

export type AuthContextState = {
  /* Is a user currently logged into the app */
  isLoggedIn: boolean;
  /* Infformation regarding the authenticated user if logged in */
  user: UserType | null;
  /* API related error message */
  authError: string | null;
  /* Function to make API call to create a new user */
  register: (email: string, username: string, password: string) => Promise<boolean>;
  /* Function to authenticate a user */
  login: (username: string, password: string) => Promise<boolean>;
  /* Function to  remove user information from app */
  logout: () => void;
  /* Clears the saved auth error message */
  clearAuthError: () => void;
};

const AuthContext = createContext<AuthContextState | undefined>(undefined);

export default AuthContext;
