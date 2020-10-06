import React from 'react';
import { Link } from 'react-router-dom';
import * as S from './styles';

function NavBar() {
  return (
    <>
      <S.Nav>
        <S.NavWrap>
          <S.NavBrand>
            <Link to="/">Hotelio</Link>
          </S.NavBrand>
          <S.NavLinks>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/">About</Link>
            </li>
            <li>
              <Link to="/">Contact</Link>
            </li>
            <li>
              <S.StyledLink to="/login">Book Now</S.StyledLink>
            </li>
          </S.NavLinks>
        </S.NavWrap>
      </S.Nav>
    </>
  );
}

export default NavBar;
