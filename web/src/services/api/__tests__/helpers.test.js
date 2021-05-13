import axios from 'axios';

import { login } from '../helpers';

jest.mock('axios');
jest.mock('jwt-decode', () => () => ({ ID: 'test_id', Email: 'test_email' }));

describe('login', () => {
  it('returns user object after successful login', async () => {
    axios.post.mockResolvedValue({ data: { token: 'test_token' }});

    const user = await login('test@example.com', 'testpassword');
  
    expect(user.id).toBe('test_id');
    expect(user.email).toBe('test_email');
  });

  it('throws error after unsuccessful login', async () => {
    axios.post.mockImplementation(() => {
      return new Error();
    });

    await expect(login('test@example.com', 'testpassword')).rejects.toThrow();
  });
});