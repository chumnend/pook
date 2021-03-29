import { Link } from 'react-router-dom';

import {
  AuthButton,
  AuthCard,
  // AuthError,
  AuthForm,
  AuthInput,
} from '../components/Auth';

const Register = () => {
  return (
    <AuthCard>
      <h2>Login</h2>
      {/* <AuthError></AuthError> */}
      <AuthForm>
        <AuthInput type="email" placeholder="Your Email" />
        <AuthInput type="password" placeholder="Your Password" />
        <AuthInput type="password2" placeholder="Confirm Password" />
        <AuthButton>Login</AuthButton>
      </AuthForm>
      <p>
        Already have an account? <Link to="/login">Sign In</Link>
      </p>
    </AuthCard>
  );
};

export default Register;
