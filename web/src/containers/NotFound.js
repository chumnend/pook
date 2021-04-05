import Navbar from '../components/Navbar';
import { useAuth } from '../context/auth';

const NotFound = () => {
  const auth = useAuth();

  return (
    <>
      <Navbar isLoggedIn={auth.isLoggedIn}></Navbar>
      <h2>Layout not found</h2>
    </>
  );
};

export default NotFound;
