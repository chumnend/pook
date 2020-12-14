import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import PropTypes from 'prop-types';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import useStyles from './styles';

const Header = (props) => {
  const classes = useStyles(props);

  return (
    <header className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="h1" className={classes.title}>
            Bookings
          </Typography>
          {props.links.map(
            (link) =>
              link.requiresAuth === props.isLoggedIn && (
                <Button
                  key={link.label}
                  color="inherit"
                  to={link.url}
                  component={RouterLink}
                >
                  {link.label}
                </Button>
              ),
          )}
        </Toolbar>
      </AppBar>
    </header>
  );
};

Header.propTypes = {
  links: PropTypes.array,
  isLoggedIn: PropTypes.bool,
};

export default Header;
