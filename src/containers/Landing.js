import React from 'react';
import Typography from '@material-ui/core/Typography';
import Hero from '../components/Hero';

const Landing = () => {
  return (
    <div>
      <Hero>
        <Typography
          component="h1"
          variant="h2"
          align="center"
          color="textPrimary"
          gutterBottom
        >
          Bookings
        </Typography>
        <Typography variant="h5" align="center" color="textSecondary" paragraph>
          Find Hotel Bookings and book your dream vacation.
        </Typography>
      </Hero>
    </div>
  );
};

export default Landing;
