import PropTypes from 'prop-types';

const AuthCard = (props) => {
  const { children } = props;

  return <div>{children}</div>;
};

AuthCard.propTypes = {
  children: PropTypes.node,
};

export default AuthCard;
