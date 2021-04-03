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

const Register = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [password2, setPassword2] = useState('');

  const history = useHistory();
  const auth = useAuth();
  const authRef = useRef(auth);

  useEffect(() => {
    authRef.current.clearError();
  }, []);

  const validateInput = () => {
    return email.length > 0 && password.length > 0 && password === password2;
  };

  const handleRegister = async (event) => {
    event.preventDefault();

    const isSuccess = await auth.register(email, password);
    if (isSuccess) {
      history.push(ROUTES.HOME);
    }
  };

  return (
    <Page>
      <AuthCard>
        <AuthTitle>Let&apos;s Get Started</AuthTitle>
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
          <AuthInput
            type="password"
            placeholder="Confirm Password"
            value={password2}
            onChange={(e) => setPassword2(e.target.value)}
          />
          <AuthButton onClick={handleRegister} disabled={!validateInput()}>
            Register
          </AuthButton>
        </AuthForm>
        <AuthText>
          Already have an account? <Link to={ROUTES.LOGIN}>Sign In</Link>
        </AuthText>
        <AuthText>
          <Link to={ROUTES.LANDING}>Back to Home</Link>
        </AuthText>
      </AuthCard>
    </Page>
  );
};

export default Register;
