import axios from 'axios'; 

import { login, register, logout } from '../helpers';

jest.mock('axios');
jest.mock('jwt-decode', () => () => ({ ID: 'test_id', Email: 'test_email' }));

describe('login', () => {
  afterEach(() => {
    delete axios.defaults.headers.common['Authorization'];
  });

  it('returns user object after successful login', async () => {
    axios.post.mockResolvedValue({ data: { token: 'test_token' }});

    const user = await login('test@example.com', 'testpassword');

    expect(user.id).toBe('test_id');
    expect(user.email).toBe('test_email');
    expect(user.token).toBe('test_token');
    expect(axios.defaults.headers.common['Authorization']).toBe('Bearer test_token');
  });

  it('throws error after unsuccessful login', async () => {
    axios.post.mockImplementation(() => {
      return new Error();
    });

    await expect(login('test@example.com', 'testpassword')).rejects.toThrow();
  });
});

describe('register', () => {
  afterEach(() => {
    delete axios.defaults.headers.common['Authorization'];
  });

  it('expects to successfully create a new user', async () => {
    axios.post.mockResolvedValue({ data: { token: 'test_token' }});

    const user = await register('test@example.com', 'testpassword');

    expect(user.id).toBe('test_id');
    expect(user.email).toBe('test_email');
    expect(user.token).toBe('test_token');
    expect(axios.defaults.headers.common['Authorization']).toBe('Bearer test_token');
  });

  it('throws error after unsuccessful registration', async () => {
    axios.post.mockImplementation(() => {
      return new Error();
    });

    await expect(register('test@example.com', 'testpassword')).rejects.toThrow();
  });
})

describe('logout', () => {
  it('expects to erase Authorization header on logout', () => {
    axios.defaults.headers.common['Authorization'] = 'test';

    logout();

    expect(axios.defaults.headers.common['Authorization']).toBe(undefined);
  });
});