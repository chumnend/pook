import React from 'react';
import PropTypes from 'prop-types';
import Container from '@material-ui/core/Container';
import useStyles from './styles';

const Hero = (props) => {
  const classes = useStyles(props);

  return (
    <div className={classes.heroContent}>
      <Container maxWidth="sm">{props.children}</Container>
    </div>
  );
};

Hero.propTypes = {
  children: PropTypes.node,
};

export default Hero;
