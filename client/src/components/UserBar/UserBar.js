import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { Nav, NavWrap, NavLinks } from './styles';

function UserBar(props) {
  const handleLogout = () => {
    props.logout();
  };

  return (
    <Nav>
      <NavWrap>
        <NavLinks>
          <li>
            <Link to="/">Payment Options</Link>
          </li>
          <li>
            <Link to="/">Terms/Conditions</Link>
          </li>
        </NavLinks>
        {!props.isLoggedIn ? (
          <NavLinks>
            <li>
              <Link to="/login">Login</Link>
            </li>
            <li>
              <Link to="/register">Register</Link>
            </li>
          </NavLinks>
        ) : (
          <NavLinks>
            <button onClick={handleLogout}>Logout</button>
          </NavLinks>
        )}
      </NavWrap>
    </Nav>
  );
}

UserBar.propTypes = {
  isLoggedIn: PropTypes.bool.isRequired,
  logout: PropTypes.func.isRequired,
};

export default UserBar;
