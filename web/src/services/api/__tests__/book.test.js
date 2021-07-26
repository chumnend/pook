import axios from "axios";

import {
  listBooks,
  createBook,
  getBook,
  updateBook,
  deleteBook,
} from '../book';

jest.mock('axios');

describe('list books', () => {
  it('gets list of books', async () => {
    // setup
    const mockBook = {
      "id": 1,
      "title": "test",
      "createdAt":  Date.now(),
      "updatedAt": Date.now(),
      "userID": 1,
    }

    axios.get.mockResolvedValue({data: { books: [ mockBook ] }});

    // run
    const books = await listBooks(1);

    // check
    expect(books.length).toBe(1);
  });

  it('fails to get books', async () => {
    axios.get.mockImplementation(() => {
      return new Error();
    });

    await expect(listBooks(1)).rejects.toThrow();
  });
});

describe('create book', () => {
  it.skip('creates a book', () => {});
  it.skip('fails to create a book', () => {});
});

describe('get book', () => {
  it.skip('gets a book', () => {});
  it.skip('fails to get a book', () => {});
});

describe('update book', () => {
  it.skip('updates a book', () => {});
  it.skip('fails to update a book', () => {});
});

describe('delete book', () => {
  it.skip('deletes a book', () => {});
  it.skip('fails to delete a book', () => {});
});
