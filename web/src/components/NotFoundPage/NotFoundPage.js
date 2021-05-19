import Box from '@material-ui/core/Box';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import { Link } from 'react-router-dom';

import { useAuth } from '../AuthProvider';
import Header from '../Header';
import { HOME_ROUTE } from '../Router';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100vw',
    height: '100vh',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    background: theme.palette.secondary.main,
  },
  box: {
    padding: '0.5rem',
    textAlign: 'center',
  },
  button: {
    padding: '1rem 1.5rem',
    marginTop: '2rem',
    background: theme.palette.primary.main,
    color: theme.palette.text.light,
    textTransform: 'uppercase',
    fontWeight: '700',

    '&:hover': {
      background: theme.palette.primary.dark,
    },
  },
}));

const NotFoundPage = () => {
  const auth = useAuth();
  const classes = useStyles();

  return (
    <Paper className={classes.root}>
      <Header isAuth={auth.isAuth} />
      <Box className={classes.box}>
        <Typography variant="h1">404</Typography>
        <Typography variant="h4">Sorry, this page was not found</Typography>
        <Button
          className={classes.button}
          variant="contained"
          component={Link}
          to={HOME_ROUTE}
        >
          Back to home
        </Button>
      </Box>
    </Paper>
  );
};

export default NotFoundPage;
