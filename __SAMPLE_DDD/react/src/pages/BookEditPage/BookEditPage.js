import { makeStyles } from '@material-ui/core/styles';

import Header from '../../components/Header';
import { useAuth } from '../../providers/AuthProvider';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100vw',
    height: '100vh',
    background: theme.palette.secondary.main,
    color: theme.palette.text.light,
  },
}));

const BookEditPage = () => {
  const auth = useAuth();
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Header isAuth={auth.isAuth} />
      <h1>BookEdit Page</h1>
    </div>
  );
};

export default BookEditPage;
