import React from 'react';
import PropTypes from 'prop-types';
import {
  StyledNav,
  StyledWrapper,
  StyledUl,
  StyledLi,
  StyledLink,
  StyledLogo,
  StyledMenu,
  StyledBookNow,
} from './styles';

function Navigation(props) {
  return (
    <StyledNav>
      <StyledWrapper>
        <StyledLogo to="/">Hotelio</StyledLogo>
        <StyledMenu onClick={props.openDrawer}>
          <div />
          <div />
          <div />
        </StyledMenu>
        <StyledUl>
          <StyledLi>
            <StyledLink to="/">Home</StyledLink>
          </StyledLi>
          <StyledLi>
            <StyledLink to="/">About</StyledLink>
          </StyledLi>
          <StyledLi>
            <StyledLink to="/">Contact Us</StyledLink>
          </StyledLi>
          <StyledLi>
            <StyledBookNow to="/">Book Now</StyledBookNow>
          </StyledLi>
        </StyledUl>
      </StyledWrapper>
    </StyledNav>
  );
}

Navigation.propTypes = {
  openDrawer: PropTypes.func.isRequired,
};

export default Navigation;
