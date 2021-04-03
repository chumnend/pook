import styles from './AuthInput.module.css';

const AuthInput = (props) => {
  return <input className={styles.input} {...props} />;
};

export default AuthInput;
