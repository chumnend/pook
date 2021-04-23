import { useState } from 'react';
import { useHistory } from 'react-router-dom';

import { HeroLayout } from '../../common/components/Hero';
import Navbar from '../../common/components/Navbar';
import * as ROUTES from '../../common/constants/routes';
import useAuth from '../../common/hooks/useAuth';

const Landing = () => {
  const [email, setEmail] = useState('');

  const auth = useAuth();
  const history = useHistory();

  const redirectWithEmail = () => {
    history.push(ROUTES.REGISTER, { email });
  };

  return (
    <>
      <Navbar isLoggedIn={auth.isLoggedIn}></Navbar>
      <HeroLayout>
        <h1>Pook gravida tincidunt sem in, semper tempus erat.</h1>
        <p>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi magna
          elit, rhoncus ac ultricies sed, congue quis tortor. Suspendisse a
          neque ut ex mattis fermentum. Praesent eleifend tortor massa, non
          cursus mi pretium eget. Quisque sem tellus, gravida tincidunt sem in,
          semper tempus erat.
        </p>
        <div>
          <input
            type="email"
            placeholder="Your Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <button onClick={redirectWithEmail}>Sign Up - It&apos;s free</button>
        </div>
      </HeroLayout>
    </>
  );
};

export default Landing;
