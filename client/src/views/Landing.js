import React from 'react';
import styled from 'styled-components';

const StyledLanding = styled.div`
  width: 100%;
  height: 100%;
  padding: 0 2rem;
  background: orange;
  display: flex;
  justify-content: center;
  align-items: center;
`;

function Landing() {
  return (
    <StyledLanding>
      <h1>Landing Page</h1>
    </StyledLanding>
  );
}

export default Landing;
