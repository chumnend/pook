const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const createBook = async (userId: string, title: string) => {
  const response = await fetch(`${API_BASE_URL}/books`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ userId, title }),
  });

  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || `HTTP Error: ${response.status}`);
  }

  return response.json();
};

export const getAllBooks = async() => {
  const response = await fetch(`${API_BASE_URL}/books` , {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || `HTTP Error: ${response.status}`);
  }

  return response.json();
}

export const getAllBooksByUserId = async (userId: string) => {
  const response = await fetch(`${API_BASE_URL}/books?user_id=${encodeURIComponent(userId)}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || `HTTP Error: ${response.status}`);
  }

  return response.json();
};

export const getBookById = async (bookId: string) => {
  const response = await fetch(`${API_BASE_URL}/books/${encodeURIComponent(bookId)}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || `HTTP Error: ${response.status}`);
  }

  return response.json();
};

export const updateBook = async (bookId: string, title: string) => {
  const response = await fetch(`${API_BASE_URL}/books/${encodeURIComponent(bookId)}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ title }),
  });

  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || `HTTP Error: ${response.status}`);
  }

  return response.json();
};

export const deleteBook = async (bookId: string) => {
  const response = await fetch(`${API_BASE_URL}/books/${encodeURIComponent(bookId)}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || `HTTP Error: ${response.status}`);
  }

  return response.json();
};
