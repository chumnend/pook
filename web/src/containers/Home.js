import Navbar from '../components/Navbar';
import Page from '../components/Page';
import { useAuth } from '../context/auth';

const Home = () => {
  const auth = useAuth();

  return (
    <Page>
      <Navbar isLoggedIn={auth.isLoggedIn}></Navbar>
      <h2>Home</h2>
    </Page>
  );
};

export default Home;
