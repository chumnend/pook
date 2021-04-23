import PropTypes from 'prop-types';

import styles from './AuthLayout.module.css';

const AuthLayout = (props) => {
  const { children } = props;

  return <section className={styles.auth}>{children}</section>;
};

AuthLayout.propTypes = {
  children: PropTypes.node,
};

export default AuthLayout;
