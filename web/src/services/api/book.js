import axios from 'axios';

// By default the react app runs on the same server as the api.
// In development mode, a different server can be pointed to using REACT_APP_API_PREFIX
let apiPrefix = '';
/* istanbul ignore next */
if (process.env.NODE_ENV === 'development') {
  apiPrefix = process.env.REACT_APP_API_PREFIX;
}

/** user api routes */
export const API_BOOK = apiPrefix + 'v1/books';
export const API_BOOK_ID = (id) => apiPrefix + `v1/books/${id}`;

export const listBooks = async (userId) => {
  try {
    const res = await axios.get(API_BOOK + `?userId=${userId}`);
    const { books } = res.data;

    return books;
  } catch (error) {
    throw error;
  }
};

export const createBook = async (title, userId) => {};

export const getBook = async (id) => {};

export const updateBook = async (id, updatedTitle) => {};

export const deleteBook = async (id) => {};
