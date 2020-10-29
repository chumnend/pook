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

function Register(props) {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [password2, setPassword2] = useState('');
  const [isLoading, setLoading] = useState(false);
  const [errors, setErrors] = useState({});
  const { authorizeUser } = useContext(AuthContext);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);

    try {
      await authorizeUser('register', {
        name,
        email,
        password,
        password2,
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
          {errors.extra && <small>{errors.extra.name}</small>}
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
          {errors.extra && <small>{errors.extra.password2}</small>}
        </StyledFormGroup>

        <StyledButton type="submit" disabled={isLoading}>
          Register
        </StyledButton>
        <p>
          Already have an account? <Link to="/login">Sign In</Link>
        </p>
      </StyledForm>
    </Page>
  );
}

export default Register;
