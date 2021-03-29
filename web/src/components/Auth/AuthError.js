import PropTypes from 'prop-types';

const AuthError = (props) => {
  const { children } = props;

  return <div>{children}</div>;
};

AuthError.propTypes = {
  children: PropTypes.node,
};

export default AuthError;
