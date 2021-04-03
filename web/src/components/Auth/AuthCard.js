import PropTypes from 'prop-types';
import styles from './AuthCard.module.css';

const AuthCard = (props) => {
  const { children } = props;

  return <div className={styles.card}>{children}</div>;
};

AuthCard.propTypes = {
  children: PropTypes.node,
};

export default AuthCard;
