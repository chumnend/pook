import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import * as S from './styles';

function UserBar(props) {
  const handleLogout = () => {
    props.logout();
  };

  return (
    <S.Nav>
      <S.NavWrap>
        <S.NavLinks>
          <li>
            <Link to="/">Payment Options</Link>
          </li>
          <li>
            <Link to="/">Terms/Conditions</Link>
          </li>
        </S.NavLinks>
        {!props.isLoggedIn ? (
          <S.NavLinks>
            <li>
              <Link to="/login">Login</Link>
            </li>
            <li>
              <Link to="/register">Register</Link>
            </li>
          </S.NavLinks>
        ) : (
          <S.NavLinks>
            <button onClick={handleLogout}>Logout</button>
          </S.NavLinks>
        )}
      </S.NavWrap>
    </S.Nav>
  );
}

UserBar.propTypes = {
  isLoggedIn: PropTypes.bool.isRequired,
  logout: PropTypes.func.isRequired,
};

export default UserBar;
