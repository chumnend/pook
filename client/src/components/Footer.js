import React from 'react';
import styled from 'styled-components';

const StyledFooter = styled.footer`
  width: 100%;
  height: 100%;
  padding: 0 2rem;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #203040;
  color: #ffffff;
`;

function Footer() {
  return (
    <StyledFooter>
      <p>All Rights Reserved. Nicholas Chumney.</p>
    </StyledFooter>
  );
}

export default Footer;
