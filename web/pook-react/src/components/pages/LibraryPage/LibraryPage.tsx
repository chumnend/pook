import Header from '../../shared/Header';
import styles from './LibraryPage.module.css';

import mockBooks from '../../../../testing/books';

const LibraryPage = () => {

  const renderBooks = (books: { title: string; author: string; imageUrl: string }[]) => {
    return books.map((book, index) => (
      <div key={index} className={styles.bookContainer}>
        <img src={book.imageUrl} alt={`${book.title} cover`} />
        <div>
          <p>{book.title}</p>
          <p>By {book.author}</p>
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
          {renderBooks(mockBooks)}
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
