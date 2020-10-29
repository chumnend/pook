import styled from 'styled-components';
import { Link } from 'react-router-dom';
import { device, color } from '../../../theme';

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
  display: none;

  @media all and (min-width: ${device.lg}) {
    display: flex;
  }
`;

export const StyledLi = styled.li`
  display: inline-block;
  padding: 1rem;
`;

export const StyledLogo = styled(Link)`
  color: ${color.blue};
  font-size: 2rem;
  font-weight: bold;
  text-decoration: none;
  display: inline-block;
  padding: 1rem;
`;

export const StyledMenu = styled.div`
  width: 40px;
  height: 100%;
  background: ${color.black};
  cursor: pointer;

  & div {
    width: 90%;
    height: 3px;
    margin: 6px 0;
    background: ${color.white};
  }

  @media all and (min-width: ${device.lg}) {
    display: none;
  }
`;

export const StyledLink = styled(Link)`
  text-decoration: none;
  background: inherit;
  color: inherit;
  &:hover {
    color: ${color.blue};
  }
`;

export const StyledBookNow = styled(Link)`
  text-decoration: none;
  background: ${color.white};
  color: ${color.black};
  padding: 0.5rem 1rem;
  &:hover {
    color: ${color.blue};
  }
`;
