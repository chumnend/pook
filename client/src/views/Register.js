import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import axios from 'axios';
import jwt_decode from 'jwt-decode';

const StyledRegister = styled.div`
  width: 100%;
  height: 100%;
  padding-top: 2rem;
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

function Register() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [password2, setPassword2] = useState('');
  const [isLoading, setLoading] = useState(false);
  const [errors, setErrors] = useState({});

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);

    try {
      const res = await axios.post('http://localhost:8081/api/users/register', {
        name,
        email,
        password,
        password2,
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
    <StyledRegister>
      <StyledForm onSubmit={handleSubmit}>
        <StyledFormHeader>
          <h2>Start an account</h2>
          {errors.message && <p>{errors.message}</p>}
        </StyledFormHeader>
        <StyledFormGroup>
          <label htmlFor="name">Name</label>
          <input
            required
            type="text"
            id="name"
            name="name"
            placeholder="Enter your name"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          {errors.name && <small>{errors.name}</small>}
        </StyledFormGroup>
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
        <StyledFormGroup>
          <label htmlFor="password2">Confirm password</label>
          <input
            required
            type="password"
            id="password2"
            name="password2"
            placeholder="Confirm new password"
            value={password2}
            onChange={(e) => setPassword2(e.target.value)}
          />
          {errors.password2 && <small>{errors.password2}</small>}
        </StyledFormGroup>

        <StyledButton type="submit" disabled={isLoading}>
          Register
        </StyledButton>
        <p>
          Already have an account? <Link to="/login">Sign In</Link>
        </p>
      </StyledForm>
    </StyledRegister>
  );
}

export default Register;
