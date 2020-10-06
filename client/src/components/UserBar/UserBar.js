import React from 'react';
import { Link } from 'react-router-dom';
import * as S from './styles';

function UserBar() {
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
        <S.NavLinks>
          <li>
            <Link to="/login">Login</Link>
          </li>
          <li>
            <Link to="/register">Register</Link>
          </li>
        </S.NavLinks>
      </S.NavWrap>
    </S.Nav>
  );
}

export default UserBar;
