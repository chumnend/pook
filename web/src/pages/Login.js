import { useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import { useAuth } from '../context/auth';
import {
  AuthButton,
  AuthCard,
  AuthError,
  AuthForm,
  AuthInput,
} from '../components/Auth';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const history = useHistory();
  const auth = useAuth();

  const validateInput = () => {
    return email.length > 0 && password.length > 0;
  };

  const handleLogin = async (event) => {
    event.preventDefault();

    const isSuccess = await auth.login(email, password);
    if (isSuccess) {
      history.push('/home');
    }
  };

  return (
    <AuthCard>
      <h2>Login</h2>
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
      <p>
        Don&apos;t have an account? <Link to="/register">Sign Up</Link>
      </p>
    </AuthCard>
  );
};

export default Login;
