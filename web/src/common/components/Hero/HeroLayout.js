import PropTypes from 'prop-types';

import styles from './HeroLayout.module.css';

const Hero = (props) => {
  const { children } = props;

  return (
    <div className={styles.hero}>
      <div className={styles.container}>{children}</div>
    </div>
  );
};

Hero.propTypes = {
  children: PropTypes.node,
};

export default Hero;
