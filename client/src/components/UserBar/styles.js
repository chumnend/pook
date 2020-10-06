import styled from 'styled-components';

export const Nav = styled.nav`
  width: 100%;
  height: 100%;
  background: #fff;
  color: #000;
`;

export const NavWrap = styled.div`
  width: 85%;
  height: 100%;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
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
