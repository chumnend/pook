import { useContext } from 'react';

import AuthContext from '../context/AuthContext';
import type { AuthContextState } from '../context/AuthContext';

const useAuth = (): AuthContextState => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export default useAuth;
