import Navbar from '../components/Navbar';
import Page from '../components/Page';
import { useAuth } from '../context/auth';

const NotFound = () => {
  const auth = useAuth();

  return (
    <Page>
      <Navbar isLoggedIn={auth.isLoggedIn}></Navbar>
      <h2>Page not found</h2>
    </Page>
  );
};

export default NotFound;
