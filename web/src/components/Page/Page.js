import PropTypes from 'prop-types';

import styles from './Page.module.css';

const Page = (props) => {
  const { children } = props;

  return <div className={styles.Page}>{children}</div>;
};

Page.propTypes = {
  children: PropTypes.node,
};

export default Page;
