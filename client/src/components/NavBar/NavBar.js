import React from 'react';
import { Link } from 'react-router-dom';
import { Nav, NavWrap, NavBrand, NavLinks, StyledLink } from './styles';

function NavBar() {
  return (
    <>
      <Nav>
        <NavWrap>
          <NavBrand>
            <Link to="/">Hotelio</Link>
          </NavBrand>
          <NavLinks>
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
              <StyledLink to="/">Book Now</StyledLink>
            </li>
          </NavLinks>
        </NavWrap>
      </Nav>
    </>
  );
}

export default NavBar;
