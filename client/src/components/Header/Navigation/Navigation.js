import React from 'react';
import {
  StyledNav,
  StyledWrapper,
  StyledUl,
  StyledLi,
  StyledLink,
  StyledLogo,
  StyledBookNow,
} from './styles';

function Navigation(props) {
  return (
    <StyledNav>
      <StyledWrapper>
        <StyledLogo to="/">Hotelio</StyledLogo>
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

export default Navigation;
