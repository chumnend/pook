import styled from 'styled-components';
import { Link } from 'react-router-dom';

export const Nav = styled.nav`
  width: 100%;
  height: 100%;
  background: #222;
  color: #fff;
`;

export const NavWrap = styled.div`
  width: 85%;
  height: 100%;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

export const NavBrand = styled.div`
  & a {
    color: inherit;
    font-size: 2rem;
    font-weight: bold;
    text-decoration: none;
    display: inline-block;
    padding: 1rem;
  }
`;

export const NavLinks = styled.ul`
  list-style: none;
  display: flex;
  & li {
    display: inline-block;
    padding: 1rem;
  }
  & a {
    color: inherit;
    text-decoration: none;
    &:hover {
      color: #ff8000;
    }
  }
`;

export const StyledLink = styled(Link)`
  color: #000 !important;
  background: #fff;
  padding: 0.5rem 1rem;
`;
