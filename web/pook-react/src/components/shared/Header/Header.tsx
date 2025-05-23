import { Link } from 'react-router-dom';

import useAuth from '../../../helpers/hooks/useAuth';
import styles from './Header.module.css';

function Header() {
  const { isLoggedIn, user } = useAuth();

  return (
    <nav className={styles.navbar}>
      <div>
        <Link to="/" className={styles.brand}>Pook</Link>
      </div>
      <div>
        {!isLoggedIn ? (
          <>
            <Link to="/login" className={styles.navLink}>Login</Link>
            <Link to="/register" className={styles.navLink}>Register</Link>
          </>
        ) : (
          <>
            <Link to="/book/new" className={styles.navLink}>New Book</Link>
            <Link to={`/library/${user?.id}`} className={styles.navLink}>My Library</Link>
            <Link to="/logout" className={styles.navLink}>Logout</Link>
          </>
        )}
      </div>
    </nav>
  )
}

export default Header;
