import PropTypes from 'prop-types';

import styles from './AuthTitle.module.css';

const AuthTitle = (props) => {
  const { children, ...otherProps } = props;

  return (
    <h2 className={styles.title} {...otherProps}>
      {children}
    </h2>
  );
};

AuthTitle.propTypes = {
  children: PropTypes.node,
};

export default AuthTitle;
