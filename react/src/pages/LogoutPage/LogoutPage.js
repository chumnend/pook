import React, { useEffect, useRef } from 'react';
import { Redirect } from 'react-router-dom';

import { HOME_ROUTE } from '../../components/Router';
import { useAuth } from '../../providers/AuthProvider';

const Logout = () => {
  const auth = useAuth();
  const authRef = useRef(auth);

  useEffect(() => {
    authRef.current.logout();
  }, []);

  return <Redirect to={HOME_ROUTE} />;
};

export default Logout;
