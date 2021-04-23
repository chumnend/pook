import Navbar from '../../common/components/Navbar';
import useAuth from '../../common/hooks/useAuth';

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
