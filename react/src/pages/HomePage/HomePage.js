import Container from '@material-ui/core/Container';
import { makeStyles } from '@material-ui/core/styles';
import { useEffect, useRef, useState } from 'react';

import Header from '../../components/Header';
import { useAuth } from '../../providers/AuthProvider';
import apiHelpers from '../../services/api';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100vw',
    height: '100vh',
    background: theme.palette.secondary.main,
    color: theme.palette.text.light,
  },
  button: {
    margin: theme.spacing(1),
  },
}));

const HomePage = () => {
  const [myBooks, setMyBooks] = useState([]);
  const classes = useStyles();
  const auth = useAuth();
  const authRef = useRef(auth);

  useEffect(() => {
    const fetchMyBooks = async () => {
      try {
        const { id } = authRef.current.user;
        const books = await apiHelpers.listBooks(id);
        setMyBooks(books);
      } catch (error) {
        // FixMe: Need to handle errors if unable to fetch books
        console.error('unable to load books');
      }
    };

    fetchMyBooks();
  }, []);

  const myBooksList = myBooks.map((book, idx) => (
    <div key={idx}>
      <h2>{book.title}</h2>
    </div>
  ));

  return (
    <div className={classes.root}>
      <Header isAuth />
      <Container>
        <h1>My Books</h1>
        <div>{myBooksList}</div>
      </Container>
    </div>
  );
};

export default HomePage;
