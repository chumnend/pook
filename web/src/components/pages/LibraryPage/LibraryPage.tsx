import { useEffect, useState } from 'react';

import Header from '../../common/Header';
import useAuth from '../../../helpers/hooks/useAuth';
import { getAllBooksByUserId } from '../../../helpers/services/book';
import type { Book } from '../../../helpers/types';
import styles from './LibraryPage.module.css';

const LibraryPage = () => {
  const [books, setBooks] = useState<Book[]>([]);
  const { user } = useAuth();

  useEffect(() => {
    const fetchBooks = async () => {
      if (user) {
        try {
          const data = await getAllBooksByUserId(user.id);
          setBooks(data.books ?? []);
        } catch(error) {
          console.error(error);
          return;
        }
      }
    }

    fetchBooks();
  }, [user])

  const renderBooks = (books: { title: string; imageUrl: string }[]) => {
    if (books.length === 0) {
      return (
        <div>
          <h2>Nothing is here yet...</h2>
        </div>
      )
    }

    return books.map((book, index) => (
      <div key={index} className={styles.bookContainer}>
        <img src={book.imageUrl} alt={`${book.title} cover`} />
        <div>
          <p>{book.title}</p>
        </div>
      </div>
    ));
  };

  return (
    <div>
      <Header />
      <div className={styles.wrapper}>
        <div className={styles.topPageHeader}>
          <p>User234313's Library</p>
          <p>Page x of y</p>
        </div>
        <div className={styles.booksContainer}>
          {renderBooks(books)}
        </div>
        <div className={styles.bottomPageHeader}>
          <div>
            <p>Previous Page</p>
          </div>
          <div>
            <p>1 2 3 4 5</p>
          </div>
          <div>
            <p>Next Page</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default LibraryPage;
