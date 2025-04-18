import { Link } from 'react-router-dom';

import styles from './LandingPage.module.css';

const LandingPage = () => {
  return (
    <div>
      {/* Navbar */}
      <nav className={styles.navbar}>
        <h1>Pook</h1>
        <Link to="/login" className={styles.navLink}>Login</Link>
        <Link to="/register" className={styles.navLink}>Register</Link>
      </nav>

      {/* Main Content */}
      <div className={styles.mainContent}>
        <h2>Welcome to Pook!</h2>
        <p>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
          Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </p>
      </div>
    </div>
  );
};


export default LandingPage;
