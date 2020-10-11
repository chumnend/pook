import styled from 'styled-components';
import { Link } from 'react-router-dom';
import { device, color } from '../../../theme';

export const StyledDrawer = styled.div`
  width: ${(props) => (props.show ? '60%' : '0')};
  height: 100vh;
  background: ${color.black};
  color: ${color.white};
  position: fixed;
  z-index: 10;
  top: 0;
  right: 0;
  transition: width 0.3s ease-out;
  display: flex;
  flex-flow: column;

  @media all and (min-width: ${device.lg}) {
    display: none;
  }
`;

export const StyledClose = styled.div`
  width: 100%;
  padding: 1rem;
`;

export const StyledCloseIcon = styled.span`
  color: inherit;
  cursor: pointer;
  font-size: 1.5rem;
  float: right;
`;

export const StyledUl = styled.ul`
  list-style: none;
  width: 100%;
`;

export const StyledLi = styled.li`
  display: block;
  padding: 1rem;
`;

export const StyledLink = styled(Link)`
  text-decoration: none;
  background: inherit;
  color: inherit;
  font-size: 1.2rem;
  &:hover {
    color: #ff8000;
  }
`;
