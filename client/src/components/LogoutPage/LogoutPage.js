import React, { useEffect, useRef } from 'react';
import { Redirect } from 'react-router-dom';

import { useAuth } from '../AuthProvider';
import { HOME_ROUTE } from '../Router';

const Logout = () => {
  const auth = useAuth();
  const authRef = useRef(auth);

  useEffect(() => {
    authRef.current.logout();
  }, []);

  return <Redirect to={HOME_ROUTE} />;
};

export default Logout;
