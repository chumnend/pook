import React, { useState } from 'react';
import PropTypes from 'prop-types';
import Profile from './Profile';
import Navigation from './Navigation';
import Drawer from './Drawer';
import { StyledHeader } from './styles';

function Header(props) {
  const [showDrawer, setShowDrawer] = useState(false);

  return (
    <StyledHeader>
      <Profile {...props} />
      <Navigation openDrawer={() => setShowDrawer(true)} />
      <Drawer
        {...props}
        show={showDrawer}
        closeDrawer={() => setShowDrawer(false)}
      />
    </StyledHeader>
  );
}

Header.propTypes = {
  isLoggedIn: PropTypes.bool.isRequired,
  logout: PropTypes.func.isRequired,
};

export default Header;
