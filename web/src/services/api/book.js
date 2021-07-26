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

/**
 * Return list of books for a given user
 * @param {*} userId
 * @returns array of book objects
 */
export const listBooks = async (userId) => {
  try {
    const res = await axios.get(API_BOOK + `?userId=${userId}`);
    const { books } = res.data;

    return books;
  } catch (error) {
    throw error;
  }
};

/**
 * Creates new book object
 * @param {string} title
 * @param {*} userId
 * @returns book object
 */
export const createBook = async (title, userId) => {
  try {
    const payload = {
      title,
      userId,
    };
    const res = await axios.post(API_BOOK, payload);
    const { book } = res.data;

    return book;
  } catch (error) {
    throw error;
  }
};

/**
 * Get a book
 * @param {*} id
 * @returns book object
 */
export const getBook = async (id) => {
  try {
    const res = await axios.get(API_BOOK_ID(id));
    const { book } = res.data;

    return book;
  } catch (error) {
    throw error;
  }
};

/**
 * Update a book
 * @param {*} id
 * @param {string} updatedTitle
 * @returns book object
 */
export const updateBook = async (id, updatedTitle) => {
  try {
    const payload = {
      updatedTitle,
    };
    const res = await axios.put(API_BOOK_ID(id), payload);
    const { book } = res.data;

    return book;
  } catch (error) {
    throw error;
  }
};

/**
 * Delete a book
 * @param {*} id
 */
export const deleteBook = async (id) => {
  try {
    await axios.delete(API_BOOK_ID(id));
  } catch (error) {
    throw error;
  }
};
