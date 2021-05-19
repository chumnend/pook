import CircularProgress from '@material-ui/core/CircularProgress';
import Container from '@material-ui/core/Container';
import { makeStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';

const useStyles = makeStyles((theme) => ({
  spinner: {
    display: 'flex',

    '& > * + *': {
      marginLeft: theme.spacing(2),
    },
  },
  fullPage: {
    width: '100vw',
    height: '100vh',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',

    '& > * + *': {
      marginLeft: theme.spacing(2),
    },
  },
}));

const Loader = ({ fullPage = false }) => {
  const classes = useStyles();

  if (fullPage) {
    return (
      <Container className={classes.fullPage} component="main" maxWidth="xs">
        <CircularProgress size={80} />
      </Container>
    );
  }

  return (
    <div className={classes.spinner}>
      <CircularProgress />
    </div>
  );
};

Loader.propTypes = {
  fullPage: PropTypes.bool,
};

export default Loader;
