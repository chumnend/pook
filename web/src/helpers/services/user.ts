const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export async function register(email: string, username: string, password: string) {
  const response = await fetch(`${API_BASE_URL}/v1/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ email, username, password }),
  });

  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || 'Something went wrong');
  }

  return response.json();
}

export const login = async (username: string, password: string) => {
  const response = await fetch(`${API_BASE_URL}/v1/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, password }),
  });

  if(!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || 'Something went wrong');
  }

  return response.json();
};

export default { register, login };
