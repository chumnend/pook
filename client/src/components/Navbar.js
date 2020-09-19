import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';

const StyledNav = styled.nav`
  width: 100%;
  height: 100%;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #203040;
  color: #ffffff;
`;

const StyledBrand = styled.div`
  & a {
    color: inherit;
    font-size: 2rem;
    font-weight: bold;
    text-decoration: none;
  }
`;

const StyledLinks = styled.ul`
  list-style: none;
  display: flex;
  & a {
    color: inherit;
    text-decoration: none;
    padding: 1rem;
    &:hover {
      color: #ff8000;
    }
  }
`;

function Navbar() {
  return (
    <StyledNav>
      <StyledBrand>
        <Link to="/">Hotelio</Link>
      </StyledBrand>
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

export default Navbar;
