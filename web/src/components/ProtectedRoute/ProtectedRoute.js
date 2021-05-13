import PropTypes from 'prop-types';
import { Redirect, Route } from 'react-router-dom';

const ProtectedRoute = (props) => {
  const { condition, redirect, ...otherProps } = props;

  if (!condition) {
    return <Redirect to={redirect} />;
  }

  return <Route {...otherProps} />;
};

ProtectedRoute.propTypes = {
  condition: PropTypes.bool.isRequired,
  redirect: PropTypes.string.isRequired,
};

export default ProtectedRoute;
