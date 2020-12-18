import React from 'react';
import PropTypes from 'prop-types';
import { useAuthContext } from '../context/auth';
import SignInForm from '../components/SignInForm';
import SignUpForm from '../components/SignUpForm';
import Logout from '../components/Logout';

const Auth = (props) => {
  const authContext = useAuthContext();

  const login = (email, password) => {
    authContext
      .login(email, password)
      .then((success) => {
        if (success) props.history.push('/');
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const register = (email, password) => {
    authContext
      .register(email, password)
      .then((success) => {
        if (success) props.history.push('/');
      })
      .catch((err) => {
        console.error(err);
      });
  };

  switch (props.type) {
    case 'login':
      return <SignInForm login={login} />;
    case 'register':
      return <SignUpForm register={register} />;
    case 'logout':
      return <Logout logout={authContext.logout} />;
    default:
      return null;
  }
};

Auth.propTypes = {
  history: PropTypes.object,
  type: PropTypes.string,
};

export default Auth;
