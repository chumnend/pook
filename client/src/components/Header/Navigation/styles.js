import styled from 'styled-components';
import { Link } from 'react-router-dom';
import { color } from '../../../theme';

export const StyledNav = styled.nav`
  width: 100%;
  height: 100%;
  background: ${color.black};
  color: ${color.white};
`;

export const StyledWrapper = styled.div`
  width: 85%;
  height: 100%;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

export const StyledUl = styled.ul`
  list-style: none;
  display: flex;
`;

export const StyledLi = styled.li`
  display: inline-block;
  padding: 1rem;
`;

export const StyledLink = styled(Link)`
  text-decoration: none;
  background: inherit;
  color: inherit;
  &:hover {
    color: #ff8000;
  }
`;

export const StyledLogo = styled(Link)`
  color: inherit;
  font-size: 2rem;
  font-weight: bold;
  text-decoration: none;
  display: inline-block;
  padding: 1rem;
`;

export const StyledBookNow = styled(Link)`
  text-decoration: none;
  background: ${color.white};
  color: ${color.black};
  padding: 0.5rem 1rem;
  &:hover {
    color: #ff8000;
  }
`;
