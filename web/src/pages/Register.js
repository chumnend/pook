import { Link, useHistory } from 'react-router-dom';
import { useAuth } from '../context/auth';
import {
  AuthButton,
  AuthCard,
  // AuthError,
  AuthForm,
  AuthInput,
} from '../components/Auth';
import { useState } from 'react';

const Register = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [password2, setPassword2] = useState('');

  const history = useHistory();
  const auth = useAuth();

  const validateInput = () => {
    return email.length > 0 && password.length > 0 && password === password2;
  };

  const handleRegister = async (event) => {
    event.preventDefault();

    const isSuccess = await auth.register(email, password);
    if (isSuccess) {
      history.push('/home');
    }
  };

  return (
    <AuthCard>
      <h2>Register</h2>
      {/* <AuthError></AuthError> */}
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
      <p>
        Already have an account? <Link to="/login">Sign In</Link>
      </p>
    </AuthCard>
  );
};

export default Register;
