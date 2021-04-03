import PropTypes from 'prop-types';
import styles from './AuthButton.module.css';

const AuthButton = (props) => {
  const { children, ...otherProps } = props;

  return (
    <button className={styles.button} {...otherProps}>
      {children}
    </button>
  );
};

AuthButton.propTypes = {
  children: PropTypes.node,
};

export default AuthButton;
