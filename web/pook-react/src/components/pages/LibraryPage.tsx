import Header from '../../components/Header';
import styles from '../../helpers/styles/LibraryPage.module.css';

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

  const books = [
    { title: "Book Title 1", author: "User234313", imageUrl: "https://placehold.co/400x400" },
    { title: "Book Title 2", author: "User234313", imageUrl: "https://placehold.co/400x400" },
    { title: "Book Title 3", author: "User234313", imageUrl: "https://placehold.co/400x400" },
    { title: "Book Title 4", author: "User234313", imageUrl: "https://placehold.co/400x400" },
    { title: "Book Title 5", author: "User234313", imageUrl: "https://placehold.co/400x400" },
    { title: "Book Title 6", author: "User234313", imageUrl: "https://placehold.co/400x400" },
    { title: "Book Title 7", author: "User234313", imageUrl: "https://placehold.co/400x400" },
    { title: "Book Title 8", author: "User234313", imageUrl: "https://placehold.co/400x400" },
  ];

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
