import React from 'react';
import PropTypes from 'prop-types';
import {
  StyledNav,
  StyledWrapper,
  StyledUl,
  StyledLi,
  StyledLink,
} from './styles';

function Profile(props) {
  return (
    <StyledNav>
      <StyledWrapper>
        <StyledUl>
          <StyledLi>
            <StyledLink to="/">Payment Options</StyledLink>
          </StyledLi>
          <StyledLi>
            <StyledLink to="/">Terms and Conditions</StyledLink>
          </StyledLi>
        </StyledUl>

        {props.isLoggedIn && (
          <StyledUl>
            <StyledLi>
              <StyledLink to="/">Your Bookings</StyledLink>
            </StyledLi>
            <StyledLi>
              <StyledLink to="/">Notifications</StyledLink>
            </StyledLi>
            <StyledLi>
              <button onClick={props.logout}>Logout</button>
            </StyledLi>
          </StyledUl>
        )}

        {!props.isLoggedIn && (
          <StyledUl>
            <StyledLi>
              <StyledLink to="/register">Register</StyledLink>
            </StyledLi>
            <StyledLi>
              <StyledLink to="/login">Login</StyledLink>
            </StyledLi>
          </StyledUl>
        )}
      </StyledWrapper>
    </StyledNav>
  );
}

Profile.propTypes = {
  isLoggedIn: PropTypes.bool.isRequired,
  logout: PropTypes.func.isRequired,
};

export default Profile;
