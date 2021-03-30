import { useEffect, useRef } from 'react';
import { Redirect } from 'react-router-dom';
import { useAuth } from '../context/auth';

const Logout = () => {
  const auth = useAuth();
  const authRef = useRef(auth);

  useEffect(() => {
    authRef.current.logout();
  }, []);

  return <Redirect to={'/'} />;
};

export default Logout;
