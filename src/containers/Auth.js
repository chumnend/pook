import React from 'react';
import { useState } from 'react';
import PropTypes from 'prop-types';
import { useAuthContext } from '../context/auth';

const Auth = (props) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [password2, setPassword2] = useState('');

  const authContext = useAuthContext();

  const validateForm = () => {
    if (props.login) {
      return email.length > 0 && password.length > 0;
    } else {
      return email.length > 0 && password.length > 0 && password === password2;
    }
  };

  const handleSubmit = (event) => {
    event.preventDefault();

    if (props.login) {
      authContext
        .login(email, password)
        .then((success) => console.log(success))
        .catch(() => console.log('internal error'));
    } else {
      authContext
        .register(email, password)
        .then((success) => console.log(success))
        .catch(() => console.log('internal error'));
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <h2>{props.login ? 'Welcome Back!' : 'Create An Account'}</h2>
        <div>
          <label htmlFor="email">Email</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        {!props.login && (
          <div>
            <label htmlFor="password2">Confirm Password</label>
            <input
              type="password"
              id="password2"
              value={password2}
              onChange={(e) => setPassword2(e.target.value)}
            />
          </div>
        )}
        <button disabled={!validateForm()}>
          {props.login ? 'Login' : 'Register'}
        </button>
      </form>
    </div>
  );
};

Auth.propTypes = {
  login: PropTypes.bool,
};

export default Auth;
