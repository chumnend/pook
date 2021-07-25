import axios from 'axios';

import {
  login,
  register,
  checkAuthState,
  saveAuthState,
  clearAuthState,
  AUTH_STATE_KEY,
} from '../api';

jest.mock('axios');
jest.mock('jwt-decode', () => () => ({ id: 'test_id', email: 'test_email' }));

const fakeLocalStorage = (function () {
  let store = {};

  return {
    getItem: function (key) {
      return store[key] || null;
    },
    setItem: function (key, value) {
      store[key] = value.toString();
    },
    removeItem: function (key) {
      delete store[key];
    },
    clear: function () {
      store = {};
    },
  };
})();

describe('login', () => {
  afterEach(() => {
    delete axios.defaults.headers.common['Authorization'];
  });

  it('returns user object after successful login', async () => {
    axios.post.mockResolvedValue({ data: { token: 'test_token' } });

    const user = await login('test@example.com', 'testpassword');

    expect(user.id).toBe('test_id');
    expect(user.email).toBe('test_email');
    expect(user.token).toBe('test_token');
    expect(axios.defaults.headers.common['Authorization']).toBe(
      'Bearer test_token',
    );
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
    axios.post.mockResolvedValue({ data: { token: 'test_token' } });

    const user = await register(
      'test_user',
      'test_pw',
      'test@example.com',
      'testpassword',
    );

    expect(user.id).toBe('test_id');
    expect(user.email).toBe('test_email');
    expect(user.token).toBe('test_token');
    expect(axios.defaults.headers.common['Authorization']).toBe(
      'Bearer test_token',
    );
  });

  it('throws error after unsuccessful registration', async () => {
    axios.post.mockImplementation(() => {
      return new Error();
    });

    await expect(
      register('test_user', 'test_pw', 'test@example.com', 'testpassword'),
    ).rejects.toThrow();
  });
});

describe('checkAuthState', () => {
  beforeAll(() => {
    Object.defineProperty(window, 'localStorage', {
      value: fakeLocalStorage,
    });
  });

  afterEach(() => {
    localStorage.clear();
  });

  it('expects to return authState if it exists in localStorage', () => {
    const testAuthState = {
      id: 'test_id',
      email: 'test_email',
      token: 'test_token',
    };
    localStorage.setItem(AUTH_STATE_KEY, JSON.stringify(testAuthState));

    const authState = checkAuthState();

    expect(authState).toStrictEqual(testAuthState);
  });

  it('expects to return null if it does not exist in localStorage', () => {
    const authState = checkAuthState();

    expect(authState).toBe(null);
  });
});

describe('saveAuthState', () => {
  beforeAll(() => {
    Object.defineProperty(window, 'localStorage', {
      value: fakeLocalStorage,
    });
  });

  afterEach(() => {
    localStorage.clear();
  });

  it('expects to store data in localStorage', () => {
    const testAuthState = {
      id: 'test_id',
      email: 'test_email',
      token: 'test_token',
    };

    saveAuthState(testAuthState.id, testAuthState.email, testAuthState.token);

    expect(window.localStorage.getItem(AUTH_STATE_KEY)).toBe(
      JSON.stringify(testAuthState),
    );
  });

  it('expects to not save in localStorage if an argument is undefined', () => {
    const testAuthState = {
      id: 'test_id',
      email: 'test_email',
      token: undefined,
    };

    saveAuthState(testAuthState.id, testAuthState.email, testAuthState.token);

    expect(window.localStorage.getItem(AUTH_STATE_KEY)).toBe(null);
  });
});

describe('clearAuthState', () => {
  it('expects to erase Authorization axios header', () => {
    axios.defaults.headers.common['Authorization'] = 'test';

    clearAuthState();

    expect(axios.defaults.headers.common['Authorization']).toBe(undefined);
  });
});
