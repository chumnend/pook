import { Link } from 'react-router-dom';

import styles from '../styles/Header.module.css';

function Header() {
  return (
    <nav className={styles.navbar}>
      <Link to="/" className={styles.brand}>Pook</Link>
      <Link to="/login" className={styles.navLink}>Login</Link>
      <Link to="/register" className={styles.navLink}>Register</Link>
    </nav>
  )
}

export default Header;
