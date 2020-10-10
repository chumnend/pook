import React from 'react';
import PropTypes from 'prop-types';
import Profile from './Profile';
import Navigation from './Navigation';
import { StyledHeader } from './styles';

function Header(props) {
  return (
    <StyledHeader>
      <Profile {...props} />
      <Navigation />
    </StyledHeader>
  );
}

Header.propTypes = {
  isLoggedIn: PropTypes.bool.isRequired,
  logout: PropTypes.func.isRequired,
};

export default Header;
