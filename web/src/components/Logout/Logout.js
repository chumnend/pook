import { useEffect, useRef } from 'react';
import { Redirect } from 'react-router-dom';

import * as ROUTES from '../../common/constants/routes';
import useAuth from '../../common/hooks/useAuth';

const Logout = () => {
  const auth = useAuth();
  const authRef = useRef(auth);

  useEffect(() => {
    authRef.current.logout();
  }, []);

  return <Redirect to={ROUTES.LANDING} />;
};

export default Logout;
