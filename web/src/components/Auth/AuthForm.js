import PropTypes from 'prop-types';

import styles from './AuthForm.module.css';

const AuthForm = (props) => {
  const { children, ...otherProps } = props;

  return (
    <form className={styles.form} {...otherProps}>
      {children}
    </form>
  );
};

AuthForm.propTypes = {
  children: PropTypes.node,
};

export default AuthForm;
