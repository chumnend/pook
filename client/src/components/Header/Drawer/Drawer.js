import React from 'react';
import PropTypes from 'prop-types';
import {
  StyledDrawer,
  StyledClose,
  StyledCloseIcon,
  StyledUl,
  StyledLi,
  StyledLink,
} from './styles';

function Drawer(props) {
  const handleLogout = () => {
    props.logout();
    props.closeDrawer();
  };

  return (
    <StyledDrawer show={props.show}>
      <StyledClose>
        <StyledCloseIcon onClick={props.closeDrawer}>X</StyledCloseIcon>
      </StyledClose>

      <StyledUl>
        <StyledLi>
          <StyledLink to="/" onClick={props.closeDrawer}>
            Book Now
          </StyledLink>
        </StyledLi>
        <StyledLi>
          <StyledLink to="/" onClick={props.closeDrawer}>
            Home
          </StyledLink>
        </StyledLi>
        <StyledLi>
          <StyledLink to="/" onClick={props.closeDrawer}>
            About
          </StyledLink>
        </StyledLi>
        <StyledLi>
          <StyledLink to="/" onClick={props.closeDrawer}>
            Contact Us
          </StyledLink>
        </StyledLi>
        <StyledLi />
      </StyledUl>

      {props.isLoggedIn && (
        <StyledUl>
          <StyledLi>
            <StyledLink to="/" onClick={props.closeDrawer}>
              Your Bookings
            </StyledLink>
          </StyledLi>
          <StyledLi>
            <StyledLink to="/" onClick={props.closeDrawer}>
              Notifications
            </StyledLink>
          </StyledLi>
          <StyledLi>
            <button onClick={handleLogout}>Logout</button>
          </StyledLi>
        </StyledUl>
      )}

      {!props.isLoggedIn && (
        <StyledUl>
          <StyledLi>
            <StyledLink to="/register" onClick={props.closeDrawer}>
              Register
            </StyledLink>
          </StyledLi>
          <StyledLi>
            <StyledLink to="/login" onClick={props.closeDrawer}>
              Login
            </StyledLink>
          </StyledLi>
        </StyledUl>
      )}
    </StyledDrawer>
  );
}

Drawer.propTypes = {
  show: PropTypes.bool.isRequired,
  closeDrawer: PropTypes.func.isRequired,
  isLoggedIn: PropTypes.bool.isRequired,
  logout: PropTypes.func.isRequired,
};

export default Drawer;
