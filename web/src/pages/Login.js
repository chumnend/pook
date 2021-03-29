import { Link } from 'react-router-dom';

import {
  AuthButton,
  AuthCard,
  // AuthError,
  AuthForm,
  AuthInput,
} from '../components/Auth';

const Login = () => {
  return (
    <AuthCard>
      <h2>Login</h2>
      {/* <AuthError></AuthError> */}
      <AuthForm>
        <AuthInput type="email" placeholder="Your Email" />
        <AuthInput type="password" placeholder="Your Password" />
        <AuthButton>Login</AuthButton>
      </AuthForm>
      <p>
        Don&apos;t have an account? <Link to="/register">Sign Up</Link>
      </p>
    </AuthCard>
  );
};

export default Login;
