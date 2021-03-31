import PropTypes from 'prop-types';

const AuthForm = (props) => {
  const { children, ...otherProps } = props;

  return <form {...otherProps}>{children}</form>;
};

AuthForm.propTypes = {
  children: PropTypes.node,
};

export default AuthForm;
