import { useEffect, useRef, useState } from 'react';
import { Link, useHistory } from 'react-router-dom';

import {
  AuthButton,
  AuthCard,
  AuthError,
  AuthForm,
  AuthInput,
  AuthText,
  AuthTitle,
} from '../components/Auth';
import Page from '../components/Page';
import * as ROUTES from '../constants/routes';
import { useAuth } from '../context/auth';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const history = useHistory();
  const auth = useAuth();
  const authRef = useRef(auth);

  useEffect(() => {
    authRef.current.clearError();
  }, []);

  const validateInput = () => {
    return email.length > 0 && password.length > 0;
  };

  const handleLogin = async (event) => {
    event.preventDefault();

    const isSuccess = await auth.login(email, password);
    if (isSuccess) {
      history.push(ROUTES.HOME);
    }
  };

  return (
    <Page>
      <AuthCard>
        <AuthTitle>Welcome Back</AuthTitle>
        {auth.error && <AuthError>{auth.error}</AuthError>}
        <AuthForm>
          <AuthInput
            type="email"
            placeholder="Your Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <AuthInput
            type="password"
            placeholder="Your Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <AuthButton onClick={handleLogin} disabled={!validateInput()}>
            Login
          </AuthButton>
        </AuthForm>
        <AuthText>
          Don&apos;t have an account? <Link to={ROUTES.REGISTER}>Sign Up</Link>
        </AuthText>
        <AuthText>
          <Link to={ROUTES.LANDING}>Back to Home</Link>
        </AuthText>
      </AuthCard>
    </Page>
  );
};

export default Login;
