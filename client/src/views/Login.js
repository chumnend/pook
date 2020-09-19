import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import axios from 'axios';
import jwt_decode from 'jwt-decode';

const StyledLogin = styled.div`
  width: 100%;
  height: 100%;
  padding-top: 4rem;
`;

const StyledForm = styled.form`
  width: 50%;
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
  border: 1px solid #000;
  & p {
    text-align: center;
  }
`;

const StyledFormHeader = styled.div`
  margin-bottom: 1rem;
  text-align: center;
  & h2 {
    font-size: 1.5rem;
  }
  & p {
    margin: 1rem 0;
    background: red;
    color: #000;
  }
`;

const StyledFormGroup = styled.div`
  width: 100%;
  display: flex;
  flex-direction: column;
  margin-bottom: 1rem;
  & label {
    margin-bottom: 0.3rem;
    font-size: 1rem;
  }
  & input {
    padding: 0.8rem;
  }
  & small {
    font-size: 0.8rem;
    color: red;
  }
`;

const StyledButton = styled.button`
  width: 100%;
  margin: 1rem auto;
  padding: 0.8rem 1rem;
  font-size: 1rem;
  cursor: pointer;
`;

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
    <StyledLogin>
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
    </StyledLogin>
  );
}

export default Login;
