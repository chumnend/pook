import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import jwt_decode from 'jwt-decode';
import {
  PageContent,
  StyledForm,
  StyledFormGroup,
  StyledFormHeader,
  StyledButton,
} from './styles';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isLoading, setLoading] = useState(false);
  const [errors, setErrors] = useState({});

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);

    try {
      const res = await axios.post('http://localhost:8081/api/users/login', {
        email,
        password,
      });

      const user = jwt_decode(res.data.token);
      console.log(user);
    } catch (error) {
      if (error.response !== undefined) {
        const errorObj = error.response.data.error;
        setErrors({
          message: errorObj.message,
          ...errorObj.errors,
        });
      } else {
        setErrors({
          message: 'Server Error. Try again later.',
        });
      }

      setLoading(false);
    }
  };

  return (
    <PageContent>
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
          {errors.email && <small>{errors.email}</small>}
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
          {errors.password && <small>{errors.password}</small>}
        </StyledFormGroup>

        <StyledButton type="submit" disabled={isLoading}>
          Login
        </StyledButton>
        <p>
          Don't have an account? <Link to="/register">Sign Up</Link>
        </p>
      </StyledForm>
    </PageContent>
  );
}

export default Login;
