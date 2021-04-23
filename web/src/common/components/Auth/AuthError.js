import PropTypes from 'prop-types';

import styles from './AuthError.module.css';

const AuthError = (props) => {
  const { children } = props;

  return <div className={styles.error}>{children}</div>;
};

AuthError.propTypes = {
  children: PropTypes.node,
};

export default AuthError;
