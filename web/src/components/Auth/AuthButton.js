import PropTypes from 'prop-types';

const AuthButton = (props) => {
  const { children, ...otherProps } = props;

  return <button {...otherProps}>{children}</button>;
};

AuthButton.propTypes = {
  children: PropTypes.node,
};

export default AuthButton;
