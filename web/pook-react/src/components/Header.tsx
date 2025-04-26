import { Link } from 'react-router-dom';

import useAuth from '../hooks/useAuth';
import styles from '../styles/Header.module.css';

function Header() {
  const { isLoggedIn } = useAuth();

  return (
    <nav className={styles.navbar}>
      <div>
        <Link to="/" className={styles.brand}>Pook</Link>
      </div>
      <div>
        {isLoggedIn ? (
          <>
            <Link to="/login" className={styles.navLink}>Login</Link>
            <Link to="/register" className={styles.navLink}>Register</Link>
          </>
        ) : (
          <>
            <Link to="#" className={styles.navLink}>New Book</Link>
            <Link to="#" className={styles.navLink}>My Library</Link>
            <Link to="#" className={styles.navLink}>Logout</Link>
          </>
        )}
      </div>
    </nav>
  )
}

export default Header;
