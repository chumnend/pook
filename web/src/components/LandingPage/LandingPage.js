import Box from '@material-ui/core/Box';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import { Link } from 'react-router-dom';

import { REGISTER_ROUTE } from '../Router';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100vw',
    height: '100vh',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    background: theme.palette.primary.main,
    color: theme.palette.text.light,
  },
  box: {
    padding: '0.5rem',
    textAlign: 'center',
  },
  button: {
    padding: '1rem 1.5rem',
    marginTop: '2rem',
    background: theme.palette.secondary.main,
    color: theme.palette.text.light,
    textTransform: 'uppercase',
    fontWeight: '700',

    '&:hover': {
      background: theme.palette.secondary.dark,
    },
  },
}));

const LandingPage = () => {
  const classes = useStyles();

  return (
    <Paper className={classes.root}>
      <Box className={classes.box}>
        <Typography variant="h1">Welcome to Pook!</Typography>
        <Typography variant="h4">
          A super simple planning app using React and Go.
        </Typography>
        <Button
          className={classes.button}
          variant="contained"
          component={Link}
          to={REGISTER_ROUTE}
        >
          Try for Free
        </Button>
      </Box>
    </Paper>
  );
};

export default LandingPage;
