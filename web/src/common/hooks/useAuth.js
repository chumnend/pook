import { useContext } from 'react';

import { AuthContext } from '../../services/context/auth';

const useAuth = () => {
  return useContext(AuthContext);
};

export default useAuth;
