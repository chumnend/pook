import React from 'react';
import { Page, HeroImage, HeroText } from './styles';

function Landing() {
  return (
    <Page>
      <HeroImage>
        <HeroText>
          <p>Hotelio</p>
          <h1>Enjoy a luxury experience</h1>
          <button>Book Now</button>
        </HeroText>
      </HeroImage>

      <div>Enjoy a luxury experience</div>
      <div>Heres some marketing about us</div>
      <div>Featured booking experience</div>
      <div>Sample Rooms you can book</div>
    </Page>
  );
}

export default Landing;
