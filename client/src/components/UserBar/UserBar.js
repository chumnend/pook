import React from 'react';
import { Link } from 'react-router-dom';
import { StyledNav, StyledLinks } from './styles';

function UserBar() {
  return (
    <StyledNav>
      <StyledLinks>
        <li>
          <Link to="/">Payment Options</Link>
        </li>
        <li>
          <Link to="/">Terms/Conditions</Link>
        </li>
      </StyledLinks>
      <StyledLinks>
        <li>
          <Link to="/login">Login</Link>
        </li>
        <li>
          <Link to="/register">Register</Link>
        </li>
      </StyledLinks>
    </StyledNav>
  );
}

export default UserBar;
