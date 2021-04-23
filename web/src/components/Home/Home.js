import Navbar from '../../common/components/Navbar';
import useAuth from '../../common/hooks/useAuth';

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
