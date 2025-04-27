import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

import Header from '../components/Header';
import useAuth from '../hooks/useAuth';
import styles from '../styles/LoginPage.module.css';

const LoginPage = () => {
  const [formData, setFormData] = useState({
    login: '',
    password: '',
  });
  const [errors, setErrors] = useState({
    login: '',
    password: '',
  });
  const { authError, clearAuthError, login } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    return () => {
      if (authError) {
        clearAuthError();
      }
    };
  }, [authError, clearAuthError]);

  const validateField = (name: string, value: string) => {
    let error = '';
    if (name === 'login') {
      if (value.trim() === '') {
        error = 'Username is required';
      }
    } else if (name === 'password') {
      if (value.trim() === '') {
        error = 'Password is required';
      }
    }
    return error;
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });

    const error = validateField(name, value);
    setErrors({ ...errors, [name]: error });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const newErrors = {
      login: validateField('login', formData.login),
      password: validateField('password', formData.password),
    };
    setErrors(newErrors);

    // Check if there are any errors
    const hasErrors = Object.values(newErrors).some((error) => error !== '');
    if (hasErrors) {
      return;
    }

    const isSuccess = await login(formData.login, formData.password);
    if (isSuccess) {
      navigate('/');
    }
  };

  return (
    <div>
      <Header />
      <div className={styles.container}>
        <h1>Login</h1>
        {authError && <p style={{ color: 'red' }}>{authError}</p>}
        <form onSubmit={handleSubmit} className={styles.form}>
          <div className={styles.formGroup}>
            <label htmlFor="login">Username</label>
            <input
              type="input"
              id="login"
              name="login"
              value={formData.login}
              onChange={handleChange}
              required
            />
            {errors.login && <p style={{ color: 'red' }}>{errors.login}</p>}
          </div>
          <div className={styles.formGroup}>
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              value={formData.password}
              onChange={handleChange}
              required
            />
            {errors.password && <p style={{ color: 'red' }}>{errors.password}</p>}
          </div>
          <button
            type="submit"
            className={styles.submitButton}
            disabled={
              Object.values(errors).some((error) => error !== '') || 
              Object.values(formData).some((value) => value === '')
            }
          >
            Login
          </button>
        </form>
      </div>
    </div>
  );
};

export default LoginPage;
