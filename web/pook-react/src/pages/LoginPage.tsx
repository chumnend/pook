import React, { useState } from 'react';

import Header from '../components/Header';
import styles from '../styles/LoginPage.module.css';

const LoginPage = () => {
  const [formData, setFormData] = useState({
    login: '',
    password: '',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // TODO: Add logic for registration API
    console.log('Form submitted:', formData)
  }

  return (
    <div>
      <Header />
      <div className={styles.container}>
      <h1>Login</h1>
        <form onSubmit={handleSubmit} className={styles.form}>
          <div className={styles.formGroup}>
            <label htmlFor="login">Username / Email</label>
            <input
              type="input"
              id="login"
              name="login"
              value={formData.login}
              onChange={handleChange}
              required
            />
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
          </div>
          <button type="submit" className={styles.submitButton}>
            Login
          </button>
        </form>
      </div>
    </div>
  );
}

export default LoginPage;
