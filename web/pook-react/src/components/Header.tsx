import { Link } from 'react-router-dom';

import styles from '../styles/Header.module.css';

function Header() {
  return (
    <nav className={styles.navbar}>
      <h1>Pook</h1>
      <Link to="/login" className={styles.navLink}>Login</Link>
      <Link to="/register" className={styles.navLink}>Register</Link>
    </nav>
  )
}

export default Header;
