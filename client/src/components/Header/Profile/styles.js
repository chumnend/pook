import styled from 'styled-components';
import { Link } from 'react-router-dom';
import { device, color } from '../../../theme';

export const StyledNav = styled.nav`
  display: none;
  width: 100%;
  height: 100%;
  background: ${color.white};
  color: ${color.black};

  @media all and (min-width: ${device.lg}) {
    display: block;
  }
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
