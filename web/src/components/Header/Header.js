import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import { makeStyles } from '@material-ui/core/styles';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import MenuIcon from '@material-ui/icons/Menu';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

import { LOGIN_ROUTE, LOGOUT_ROUTE } from '../Router';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
}));

const Header = ({ isAuth }) => {
  const classes = useStyles();

  return (
    <AppBar>
      <Toolbar>
        <IconButton
          edge="start"
          className={classes.menuButton}
          color="inherit"
          aria-label="menu"
        >
          <MenuIcon />
        </IconButton>
        <Typography variant="h6" className={classes.title}>
          Pook
        </Typography>
        {isAuth ? (
          <>
            <Button color="inherit" component={Link} to={LOGOUT_ROUTE}>
              Login
            </Button>
          </>
        ) : (
          <>
            <Button color="inherit" component={Link} to={LOGIN_ROUTE}>
              Login
            </Button>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
};

Header.propTypes = {
  isAuth: PropTypes.bool,
};

export default Header;
