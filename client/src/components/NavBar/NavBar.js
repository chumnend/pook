import React from 'react';
import { Link } from 'react-router-dom';
import { StyledNav, StyledBrand, StyledLinks } from './styles';

function NavBar() {
  return (
    <>
      <StyledNav>
        <StyledBrand>
          <Link to="/">Hotelio</Link>
        </StyledBrand>
        <StyledLinks>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/">About</Link>
          </li>
          <li>
            <Link to="/">Contact</Link>
          </li>
          <li>
            <Link to="/">Book Now</Link>
          </li>
        </StyledLinks>
      </StyledNav>
    </>
  );
}

export default NavBar;
