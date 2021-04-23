import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

import * as ROUTES from '../../constants/routes';
import styles from './Navbar.module.css';

const Navbar = (props) => {
  const { isLoggedIn } = props;

  return (
    <header className={styles.header}>
      <div className={styles.wrapper}>
        <Link to={ROUTES.LANDING}>POOK</Link>

        {isLoggedIn && (
          <nav className={styles.nav}>
            <Link to={ROUTES.LOGOUT}>Logout</Link>
          </nav>
        )}

        {!isLoggedIn && (
          <nav className={styles.nav}>
            {!isLoggedIn && <Link to={ROUTES.REGISTER}>Register</Link>}
            {!isLoggedIn && <Link to={ROUTES.LOGIN}>Login</Link>}
          </nav>
        )}
      </div>
    </header>
  );
};

Navbar.propTypes = {
  isLoggedIn: PropTypes.bool,
};

export default Navbar;
