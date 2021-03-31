import { Route, Redirect } from 'react-router-dom';
import PropTypes from 'prop-types';

const ProtectedRoute = (props) => {
  const { condition, redirect, ...otherProps } = props;

  if (!condition) {
    return <Redirect to={redirect} />;
  }

  return <Route {...otherProps} />;
};

ProtectedRoute.propTypes = {
  condition: PropTypes.bool,
  redirect: PropTypes.string,
};

export default ProtectedRoute;
