import Navbar from '../components/Navbar';
import { useAuth } from '../context/auth';

const Home = () => {
  const auth = useAuth();

  return (
    <>
      <Navbar isLoggedIn={auth.isLoggedIn}></Navbar>
      <h2>Home</h2>
    </>
  );
};

export default Home;
