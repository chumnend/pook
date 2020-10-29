import React, { useState, useContext } from 'react';
import { Link } from 'react-router-dom';
import { AuthContext } from '../../context/auth';
import {
  Page,
  StyledForm,
  StyledFormGroup,
  StyledFormHeader,
  StyledButton,
} from './styles';

function Login(props) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isLoading, setLoading] = useState(false);
  const [errors, setErrors] = useState({});
  const { authorizeUser } = useContext(AuthContext);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);

    try {
      await authorizeUser('login', {
        email,
        password,
      });

      props.history.push('/');
    } catch (error) {
      if (error.response !== undefined) {
        const errorObj = error.response.data.error;
        setErrors(errorObj);
      } else {
        setErrors({
          message: 'Server Error. Try again later.',
        });
      }

      setLoading(false);
    }
  };

  return (
    <Page>
      <StyledForm onSubmit={handleSubmit}>
        <StyledFormHeader>
          <h2>Sign in to your account</h2>
          {errors.message && <p>{errors.message}</p>}
        </StyledFormHeader>
        <StyledFormGroup>
          <label htmlFor="email">Email</label>
          <input
            required
            type="email"
            id="email"
            name="email"
            placeholder="Enter your email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          {errors.extra && <small>{errors.extra.email}</small>}
        </StyledFormGroup>
        <StyledFormGroup>
          <label htmlFor="password">Password</label>
          <input
            required
            type="password"
            id="password"
            name="password"
            placeholder="Enter new password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          {errors.extra && <small>{errors.extra.password}</small>}
        </StyledFormGroup>

        <StyledButton type="submit" disabled={isLoading}>
          Login
        </StyledButton>
        <p>
          Don't have an account? <Link to="/register">Sign Up</Link>
        </p>
      </StyledForm>
    </Page>
  );
}

export default Login;
