import PropTypes from 'prop-types';
import { Redirect, Route } from 'react-router-dom';

const ProtectedRoute = ({ path, condition, redirect, ...otherProps }) => {
  if (!condition) {
    return <Redirect to={redirect} />;
  }

  return <Route exact {...otherProps} />;
};

ProtectedRoute.propTypes = {
  path: PropTypes.string.isRequired,
  condition: PropTypes.bool.isRequired,
  redirect: PropTypes.string.isRequired,
};

export default ProtectedRoute;
