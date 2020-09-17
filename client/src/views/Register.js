import React, { useState } from 'react';
import axios from 'axios';
import jwt_decode from 'jwt-decode';

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
    <div>
      <form onSubmit={handleSubmit}>
        {errors.message && <p>{errors.message}</p>}
        <div>
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
        </div>
        <div>
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
        </div>
        <div>
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
        </div>
        <div>
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
        </div>

        <button type="submit" disabled={isLoading}>
          Sign Up
        </button>
      </form>
    </div>
  );
}

export default Register;
