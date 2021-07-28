import axios from 'axios';

import {
  listBooks,
  createBook,
  getBook,
  updateBook,
  deleteBook,
} from '../book';

jest.mock('axios');

const mockBook = {
  id: 1,
  title: 'test',
  createdAt: Date.now(),
  updatedAt: Date.now(),
  userID: 1,
};

describe('list books', () => {
  it('gets list of books', async () => {
    // setup
    axios.get.mockResolvedValue({
      data: {
        books: [mockBook],
      },
    });

    // run
    const books = await listBooks(1);

    // check
    expect(books.length).toBe(1);
  });

  it('fails to get books', async () => {
    // setup
    axios.get.mockImplementation(() => {
      return new Error();
    });

    // run
    await expect(listBooks(1)).rejects.toThrow();
  });
});

describe('create book', () => {
  it('creates a book', async () => {
    // setup
    axios.post.mockResolvedValue({
      data: {
        book: mockBook,
      },
    });

    // run
    const book = await createBook('test', 1);

    // check
    expect(book.id).toBe(mockBook.id);
    expect(book.title).toBe(mockBook.title);
  });

  it('fails to create a book', async () => {
    // setup
    axios.post.mockImplementation(() => {
      return new Error();
    });

    // run
    await expect(createBook('test', 1)).rejects.toThrow();
  });
});

describe('get book', () => {
  it('gets a book', async () => {
    // setup
    axios.get.mockResolvedValue({
      data: {
        book: mockBook,
      },
    });

    // run
    const book = await getBook(1);

    // check
    expect(book.id).toBe(mockBook.id);
    expect(book.title).toBe(mockBook.title);
  });

  it('fails to get a book', async () => {
    // setup
    axios.get.mockImplementation(() => {
      return new Error();
    });

    // run
    await expect(getBook(1)).rejects.toThrow();
  });
});

describe('update book', () => {
  it('updates a book', async () => {
    // setup
    axios.put.mockResolvedValue({
      data: {
        book: mockBook,
      },
    });

    // run
    const book = await updateBook('update test');

    // check
    expect(book.id).toBe(mockBook.id);
    expect(book.title).toBe(mockBook.title);
  });

  it('fails to update a book', async () => {
    // setup
    axios.put.mockImplementation(() => {
      return new Error();
    });

    // run
    await expect(updateBook(1)).rejects.toThrow();
  });
});

describe('delete book', () => {
  it('deletes a book', async () => {
    // setup
    axios.delete.mockResolvedValue({
      data: {
        result: 'success',
      },
    });

    // run
    const result = await deleteBook(1);

    // check
    expect(result).toBe(true);
  });

  it('fails to delete a book', async () => {
    // setup
    axios.delete.mockImplementation(() => {
      return new Error();
    });

    // run
    await expect(deleteBook(1)).rejects.toThrow();
  });
});
