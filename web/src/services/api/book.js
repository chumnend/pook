import axios from 'axios';

// By default the react app runs on the same server as the api.
// In development mode, a different server can be pointed to using REACT_APP_API_PREFIX
let apiPrefix = '';
/* istanbul ignore next */
if (process.env.NODE_ENV === 'development') {
  apiPrefix = process.env.REACT_APP_API_PREFIX;
}

/** user api routes */
export const API_BOOK_ROUTE = 'v1/books';

export const listBooks = () => {};
export const createBook = () => {};
export const getBook = () => {};
export const updateBook = () => {};
export const deleteBook = () => {};
