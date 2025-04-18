import Header from '../components/Header';

import styles from '../styles/LandingPage.module.css';

const LandingPage = () => {
  return (
    <div>
      <Header />
      <div className={styles.mainContent}>
        <h2>Welcome to Pook!</h2>
        <p>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
          Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </p>
      </div>
    </div>
  );
};

export default LandingPage;
