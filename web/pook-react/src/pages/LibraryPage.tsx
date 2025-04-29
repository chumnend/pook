import Header from "../components/Header";
import styles from '../styles/LibraryPage.module.css';

const LibraryPage = () => {
  return (
    <div>
      <Header />
      <div className={styles.wrapper}>
          <div className={styles.topPageHeader}>
            <p>User234313's Library</p>
            <p>Page x of y</p>
          </div> 
          <div className={styles.booksContainer}>
            <div className={styles.bookContainer}>
              <img src="https://placehold.co/400x400" />
              <div>
                <p>Book Title</p>
                <p>By User234313</p>
              </div>
            </div>
            <div className={styles.bookContainer}>
              <img src="https://placehold.co/400x400" />
              <div>
                <p>Book Title</p>
                <p>By User234313</p>
              </div>
            </div>
            <div className={styles.bookContainer}>
              <img src="https://placehold.co/400x400" />
              <div>
                <p>Book Title</p>
                <p>By User234313</p>
              </div>
            </div>
            <div className={styles.bookContainer}>
              <img src="https://placehold.co/400x400" />
              <div>
                <p>Book Title</p>
                <p>By User234313</p>
              </div>
            </div>
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
