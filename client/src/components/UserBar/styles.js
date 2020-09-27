import styled from 'styled-components';

export const StyledNav = styled.nav`
  width: 100%;
  height: 100%;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #eee;
  color: #000;
`;

export const StyledLinks = styled.ul`
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
