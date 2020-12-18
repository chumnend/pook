import React, { useEffect } from 'react';
import { Redirect } from 'react-router-dom';
import PropTypes from 'prop-types';

const Logout = (props) => {
  useEffect(() => {
    props.logout();
  }, [props]);

  return <Redirect to="/" />;
};

Logout.propTypes = {
  logout: PropTypes.func,
};

export default Logout;
