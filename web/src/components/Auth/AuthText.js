import PropTypes from 'prop-types';
import styles from './AuthText.module.css';

const AuthText = (props) => {
  const { children, ...otherProps } = props;

  return (
    <p className={styles.text} {...otherProps}>
      {children}
    </p>
  );
};

AuthText.propTypes = {
  children: PropTypes.node,
};

export default AuthText;
