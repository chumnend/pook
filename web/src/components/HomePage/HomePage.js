import Fab from '@material-ui/core/Fab';
import { makeStyles } from '@material-ui/core/styles';
import AddIcon from '@material-ui/icons/Add';

import Bookshelf from '../Bookshelf';
import Header from '../Header';
import Searchbar from '../Searchbar';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100vw',
    height: '100vh',
    background: theme.palette.secondary.main,
    color: theme.palette.text.light,
  },
}));

const HomePage = () => {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Header isAuth />
      <Searchbar />
      <Bookshelf />
      <Fab color="primary" aria-label="add">
        <AddIcon />
      </Fab>
    </div>
  );
};

export default HomePage;
