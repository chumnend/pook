import Container from '@material-ui/core/Container';
import { makeStyles } from '@material-ui/core/styles';

import Header from '../Header';

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
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Header isAuth />
      <Container>
        <h1>CONTENT</h1>
      </Container>
    </div>
  );
};

export default HomePage;
