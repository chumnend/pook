import { Link } from 'react-router-dom';

import Header from '../../common/Header';
import styles from './LandingPage.module.css';

const LandingPage = () => {
  return (
    <div>
      <Header />
      <div>
        <div className={styles.hero}>
          <img src="https://placehold.co/800x400" alt="placeholder hero" />
        </div>
        <div className={styles.mainContent}>
          <h2>Welcome to Pook!</h2>
          <p>
            Create and share your storybooks with the world. Join our community of storytellers today!
          </p>
          <div className={styles.ctaButtons}>
            <Link to="/register" className={styles.ctaButton}>Register</Link>
            <Link to="/login" className={styles.ctaButton}>Login</Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LandingPage;
