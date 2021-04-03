import Navbar from '../components/Navbar';
import Page from '../components/Page';
import { useAuth } from '../context/auth';

const Landing = () => {
  const auth = useAuth();

  return (
    <Page>
      <Navbar isLoggedIn={auth.isLoggedIn}></Navbar>
      <h2>Landing</h2>
    </Page>
  );
};

export default Landing;
